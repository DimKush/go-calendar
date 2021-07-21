package main

import (
	"log"

	"github.com/DimKush/go-calendar"
)

func main() {
	server := new(calendar.Server)
	if err := server.Run("8000"); err != nil {
		// TODO : port from config
		log.Fatalf("error occured while running http server on port %s. Error : %s ", "8000", err.Error())
	}
}
