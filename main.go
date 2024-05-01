package main

import (
	"fmt"
	"id-generator/storage"
	"id-generator/routes"
)

func main() {
	routes.InitEcho()
	e := routes.New()

	storage.InitDB()

	e.Logger.Fatal(e.Start(":8080"))
	fmt.Print(e)
}
