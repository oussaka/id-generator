package main

import (
	"fmt"
	"id-generator/handlers"
	"id-generator/routes"
	"id-generator/storage"
)

func main() {
	routes.InitEcho()
	e := routes.New()

	storage.InitDB()

	e.Use(handlers.LogRequest)
	e.Logger.Fatal(e.Start(":8080"))
	fmt.Print(e)
}
