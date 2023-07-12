package Database

import (
	"github.com/liquedgit/tokoMeLia/graph/model"
	"github.com/liquedgit/tokoMeLia/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetInstance() *gorm.DB {
	if db == nil {
		dsn := helper.GoDotEnvVariables("DB_CONNECTION")
		myDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(err)
		}
		db = myDb
	}
	return db
}

func MigrateTable() {
	db := GetInstance()
	db.AutoMigrate(&model.User{})
}
