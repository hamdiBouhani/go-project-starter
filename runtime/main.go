package main

import (
	"log"
	"toy-project-be/common/context"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx, err := context.NewCtx()
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	err = ctx.Run()
	log.Fatalf(err.Error())

}
