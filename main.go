package main

import (
	"fmt"
	"id-generator/storage"
	"id-generator/routes"
	"id-generator/handlers"
)

func main() {
	routes.InitEcho()
	e := routes.New()

	storage.InitDB()

	e.Use(handlers.LogRequest)
	e.Logger.Fatal(e.Start(":8080"))
	fmt.Print(e)
}
