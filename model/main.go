package model

import (
	"fmt"
	"gin-example/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(config *config.Config) (err error) {
    var db *gorm.DB

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
        config.SQL.Host,
        config.SQL.Username,
        config.SQL.Password,
        config.SQL.Database,
        config.SQL.Port)

    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err == nil {
        DB = db
        err = db.AutoMigrate(&Album{})

        if err != nil {
            return err
        }
    } else {
        return err
    }

    return err
}

func CloseDB() error {
    sqlDB, err := DB.DB()
    sqlDB.Close()

    if err != nil {
        return err
    }

    err = sqlDB.Close()
    return err
}
