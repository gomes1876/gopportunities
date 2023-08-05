package config

import (
	"os"

	"github.com/gomes1876/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dbLocation string = "./db/main.db";

func InitializeSQlite() (*gorm.DB, error){
	logger := GetLogger("📚 sqlite")
	//create and connect

	_, err := os.Stat(dbLocation)

	if os.IsNotExist(err){
		logger.InffoF("🤔 dtabase not exists, creating... ")

		err = os.MkdirAll("./db", os.ModePerm);

		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbLocation)

		if err != nil {
			return nil, err
		}

		file.Close()
	}

	db, err:= gorm.Open(sqlite.Open(dbLocation), &gorm.Config{})
	if(err != nil){
		logger.ErrorF("📂 erro ao criar e conectar com o banco")
		return nil, err;
	} 

	err = db.AutoMigrate(&schemas.Opening{})

	if(err != nil){
		logger.ErrorF("💩 sqlite automigration error")
		return nil, err
	}

	return db, nil
	
}