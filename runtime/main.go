package main

import (
	"log"
	"toy-project-be/common/context"
	"toy-project-be/services"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx, err := context.NewCtx(
		&services.HttpService{},
	)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	err = ctx.Run()
	log.Fatalf(err.Error())

}
