package main

import (
	"golangLoggerCompare/util"
	"log"
)

func main() {
	slogJsonLogger := util.GetJsonLogger()
	util.LoggerSlogWithCustomField(slogJsonLogger)
	log.Println("--------")
	util.SetChildLogger()
	log.Println("--------")
	util.LoggerWithCustomLevelsAndSource()
	log.Println("--------")
	util.HidingSensitiveData()
	log.Println("--------")
	util.UseSlogFrontendZapBackend()
	log.Println("--------")
	util.SetSlogDefaultLogger()
	log.Println("--------")
}
