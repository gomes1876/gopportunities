package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error{
	var err error

	db, err = InitializeSQlite()

	if err != nil {
		return fmt.Errorf("❗️🏁 erro ao iniciar aplicação - inicialização de SQlite")
	}

	return nil
}

func GetSQlite () *gorm.DB{
	return db
}

func GetLogger(prefix string) *Logger{
	logger := NewLogger(prefix);
	return logger;
}