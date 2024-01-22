package main

import (
	"context"
	"kreditplus/src/app"
	"kreditplus/src/migration"
	"log"
)

func main() {
	// Init app context
	ctx := context.TODO()
	if err := app.Init(ctx); err != nil {
		panic(err)
	}

	// Init Migration
	if err := migration.Init(app.GormDB()); err != nil {
		panic(err)
	}

	log.Println("Finish Migration")
}
