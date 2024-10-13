package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
)

func main() {
	app *pocketbase.PocketBase = pocketbase.New()

	if err := app.start(); err != nil {
		log.Fatal(err)
	}
}
