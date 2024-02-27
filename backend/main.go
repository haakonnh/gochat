// main.go
package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
)

func main() {
	app := pocketbase.New()

	// create handlers for routes

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
