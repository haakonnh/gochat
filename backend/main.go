// main.go
package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
)

func main() {
	app := pocketbase.New()

	// customize it with event hooks https://pocketbase.io/docs/event-hooks/ ...

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
