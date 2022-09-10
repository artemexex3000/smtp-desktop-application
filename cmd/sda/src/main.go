package main

import (
	. "github.com/artemexex3000/smtp-desktop-application/internal/sda-private-lib"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	Auth()
}
