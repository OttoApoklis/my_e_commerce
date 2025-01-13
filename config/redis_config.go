package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"passwd"`
	Dbname   string `yaml:"dbname"`
	Charset  string `yaml:"charset"`
}

type RConfig struct {
	Redis RedisConfig `yaml:"redis"`
}

type RedisInstance struct {
	db *gorm.DB
}

func GetRedis() *gorm.DB {
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

func initRedis() (*gorm.DB, error) {
	var config Config
	file, err := os.Open("/users/archerxxu/Documents/UGit/src/my_e_commerce/config.yaml")
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
	fmt.Println(config.Database.Host)
	if err != nil {
		log.Fatal(err)
	}
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True",
		config.Database.User, config.Database.Password, config.Database.Host,
		config.Database.Port, config.Database.Dbname, config.Database.Charset)
	//dsn = "root:@(localhost:3306)/users?charset=utf8mb4&parseTime=True"
	fmt.Println(dsn)
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