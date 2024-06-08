package main

import (
	"flag"
	"fmt"
	"id-generator/clients"
	"id-generator/handlers"
	"id-generator/models"
	"id-generator/repositories"
	"id-generator/routes"
	"id-generator/services"
	"id-generator/storage"
	"os"

	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := routes.New()
	storage.InitDB()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(handlers.LogRequest)

	args := os.Args[2:]
	importCmd := flag.NewFlagSet("import", flag.ExitOnError)
	startDate := importCmd.String("startDate", "yesterday", "start collect from that date.")
	endDate := importCmd.String("endDate", "now", "end collect at that date.")
	useColor := importCmd.Bool("color", false, "display colorized output")

	importCmd.Parse(args)

	var accounts clients.Response

	switch os.Args[1] {
	case "import":
		accounts = clients.GetUserAccounts(*startDate, *endDate, 0, 1)
		fmt.Printf("%v", accounts.Data)
		createAccounts(accounts.Data)

		if *useColor {
			services.Colorize(services.ColorBlue, fmt.Sprintf("==> %d accounts imported", len(accounts.Data)))
		} else {
			fmt.Printf("==> %d accounts imported", len(accounts.Data))
		}

	default:
		fmt.Println("expected 'import' subcommand")
		os.Exit(1)
	}

	os.Exit(0)
}

func createAccounts(users []clients.User) {

	newUser := &models.User{}

	for _, user := range users {
		newUser = models.NewUser(user.Name, user.Email, services.GenerateUid())
		_, err := repositories.CreateUser(*newUser)
		if err != nil {
			fmt.Errorf("failed to create user : %v", err)
			os.Exit(1)
		}
	}
}
