package config

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
	once     sync.Once
)

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	Charset  string `yaml:"charset"`
}

type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

type DBInstance struct {
	db *gorm.DB
}

func GetDB() *gorm.DB {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("err %+v", err)
		}
	}()
	db, err := initDatabase()
	if err != nil {
		log.Printf("err: %+v", err)
	}
	return db
}

func initDatabase() (*gorm.DB, error) {
	once.Do(func() {
		var config Config
		file, err := os.Open("config.yaml")
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			err := recover()
			if err != nil {
				log.Printf("err: %+v", err)
			}
		}()
		defer func() {
			file.Close()
		}()
		decoder := yaml.NewDecoder(file)
		err = decoder.Decode(&config)
		if err != nil {
			log.Fatal(err)
		}
		dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True",
			config.Database.User, config.Database.Password, config.Database.Host,
			config.Database.Port, config.Database.Dbname, config.Database.Charset)
		instance, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}

		// 配置连接池
		sqlDB, err := instance.DB()
		if err != nil {
			log.Fatal(err)
		}
		sqlDB.SetMaxOpenConns(1)            // 设置最大打开连接数
		sqlDB.SetMaxIdleConns(1)            // 设置最大空闲连接数
		sqlDB.SetConnMaxLifetime(time.Hour) // 设置连接可复用的最大时间
	})
	return instance, nil
}
