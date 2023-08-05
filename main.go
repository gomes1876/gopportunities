package main

import (
	"github.com/gomes1876/gopportunities/config"
	"github.com/gomes1876/gopportunities/router"
)

var (
	logger *config.Logger
)

func main(){

	logger:= config.NewLogger("main")
	error := config.Init()
	
	if error != nil {
		logger.ErrorF("Houve um erro no main %v", error)
		return
	}

	router.Initialize();

}