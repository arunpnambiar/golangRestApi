package db

import "github.com/jinzhu/gorm"

func Connection() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	return db, nil
}
