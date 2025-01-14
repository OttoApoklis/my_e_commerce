package utils

import (
	"github.com/google/uuid"
	"log"
)

func GetUUID() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Printf("生成 UUID 时出错: %v", err)
		return "", err
	}
	log.Printf("生成的 UUID: %s", uuid.String())
	var uid string
	uid = uuid.String()
	return uid, nil
}
