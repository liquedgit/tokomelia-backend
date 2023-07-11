package Database

import (
	"github.com/liquedgit/tokoMeLia/graph/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetInstance() *gorm.DB {
	if db == nil {
		dsn := "host=localhost user=postgres password=root dbname=tokomelia port=5432 sslmode=disable TimeZone=Asia/Shanghai"
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
