package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
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
	var db *gorm.DB
	// 建立数据库连接
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, nil
}
