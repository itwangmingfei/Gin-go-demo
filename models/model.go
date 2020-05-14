package models

import (
	"fmt"
	"gin/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var db *gorm.DB

func InitDb(conf config.Mysql, debug bool) {
	if db!=nil {
		return
	}
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Passwd, conf.Host, conf.Port, conf.Db)
	var err error
	db, err = gorm.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("can't connect to db : %s\n", err.Error())
	}

	err = db.DB().Ping()
	if err != nil {
		log.Fatalf("can't ping db :%s\n", err.Error())
	}

	log.Println("ping mysql success!")

	db.DB().SetMaxOpenConns(conf.MaxIdleConns)
	db.DB().SetMaxIdleConns(conf.MaxOpenConns)
	if debug {
		db.LogMode(true)
	}
}

func GetDb() *gorm.DB {
	return db
}

func Save(value interface{}) error {
	return db.Save(&value).Error
}

func Create(values interface{}) (err error) {
	err = db.Create(values).Error

	return
}

func Update(values interface{}) error {
	return db.Update(&values).Error
}

func Updates(model interface{}, updates interface{}) (affectedRows int64, err error) {
	result := db.Where(model).Updates(updates)
	affectedRows = result.RowsAffected
	err = result.Error
	log.Infof("Updates : affectedRows %d, err: %v", affectedRows, err)
	return
}