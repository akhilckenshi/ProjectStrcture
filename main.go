package main

import (
	"exclusive-sample/internal/config"
	"exclusive-sample/internal/connection"
	"log"
	"os"
)

func main() {

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	server, connErr := connection.InitializeAPI(config)

	if connErr != nil {
		log.Fatal("cannot start server: ", connErr)
	} else {
		server.Start(config, infoLog, errorLog)
	}

}
