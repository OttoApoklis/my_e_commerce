package utils

import (
	"fmt"
	"sync"
	"time"
)

const (
	workerIDBits     = uint(5)
	datacenterIDBits = uint(5)
	maxWorkerID      = -1 ^ (-1 << workerIDBits)
	maxDatacenterID  = -1 ^ (-1 << datacenterIDBits)
	sequenceBits     = uint(12)

	workerIDShift     = sequenceBits
	datacenterIDShift = sequenceBits + workerIDBits
	timestampShift    = sequenceBits + workerIDBits + datacenterIDBits
	sequenceMask      = -1 ^ (-1 << sequenceBits)

	twepoch = 1288834974657
)

type SnowflakeIDWorker struct {
	workerID      int64
	datacenterID  int64
	sequence      int64
	lastTimestamp int64
	mutex         sync.Mutex
}

func NewSnowflakeIDWorker(workerID, datacenterID int64) *SnowflakeIDWorker {
	return &SnowflakeIDWorker{
		workerID:      workerID,
		datacenterID:  datacenterID,
		sequence:      0,
		lastTimestamp: -1,
	}
}

func (sf *SnowflakeIDWorker) NextID() (int64, error) {
	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	now := time.Now().UnixNano() / 1e6
	if now < sf.lastTimestamp {
		return 0, fmt.Errorf("Clock moved backwards. Refusing to generate id for %d milliseconds", sf.lastTimestamp-now)
	}

	if now == sf.lastTimestamp {
		sf.sequence = (sf.sequence + 1) & sequenceMask
		if sf.sequence == 0 {
			for now <= sf.lastTimestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		sf.sequence = 0
	}

	sf.lastTimestamp = now

	id := ((now - twepoch) << timestampShift) |
		(sf.datacenterID << datacenterIDShift) |
		(sf.workerID << workerIDShift) |
		sf.sequence

	return id, nil
}

func GetSnowCode() int64 {
	worker := NewSnowflakeIDWorker(1, 1)
	id, err := worker.NextID()
	if err != nil {
		fmt.Println("Error generating ID:", err)
		panic("snow code failed!")
	}
	return id
}
