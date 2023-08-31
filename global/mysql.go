package global

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySqlDB *mysqlDB

type mysqlDB struct {
	*gorm.DB
}

func InitMySql() {
	MySqlDB = NewMysqlDB()
}

func NewMysqlDB() *mysqlDB {
	return &mysqlDB{connToMysql()}
}

func connToMysql() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local",
		"user",
		"password",
		"host",
		"dbname",
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("failed to link database: %v", err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db
}
