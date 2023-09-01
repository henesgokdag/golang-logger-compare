package main

import (
	"golangLoggerCompare/util"
	"log"
)

func main() {
	slogJsonLogger := util.GetJsonLogger()
	util.LoggerSlogWithCustomField(slogJsonLogger)
	log.Println("--------")
	util.SetChildLogger(slogJsonLogger)
	log.Println("--------")
	util.LoggerWithCustomLevelsAndSource()
	log.Println("--------")
	util.SetSlogDefaultLogger(slogJsonLogger)
	log.Println("--------")
}
