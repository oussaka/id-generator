package main

import (
	"context"
	"flag"
	"fmt"
	"id-generator/cmd/cloudStorage/s3client"
	"id-generator/handlers"
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

	if len(os.Args) < 2 {
		services.Colorize(services.ColorRed, "Missing parameter, provide file name!")
		return
	}

	args := os.Args[2:]
	exporterCmd := flag.NewFlagSet("export", flag.ExitOnError)
	filename := exporterCmd.String("file", "yesterday", "start collect from that date.")
	useColor := exporterCmd.Bool("color", false, "display colorized output")

	exporterCmd.Parse(args)

	// fmt.Println("filename " + *filename)

	switch os.Args[1] {
	case "import":
		// accounts = clients.GetUserAccounts(*startDate, *endDate, 0, 1)
		// fmt.Printf("%v", accounts.Data)
		// createAccounts(accounts.Data)

		// if *useColor {
		// 	services.Colorize(services.ColorBlue, fmt.Sprintf("==> %d accounts imported", len(accounts.Data)))
		// } else {
		// 	fmt.Printf("==> %d accounts imported", len(accounts.Data))
		// }

	case "export":
		_, err := handleUpload(*filename)
		if err != nil {

			if *useColor {
				services.Colorize(services.ColorRed, err.Error())
			} else {
				fmt.Printf("Upload Object error : %s", err.Error())
			}

			os.Exit(1)
		}

		if *useColor {
			services.Colorize(services.ColorGreen, fmt.Sprintf("==> File %s uploaded with success", *filename))
		} else {
			fmt.Printf("==> File %s uploaded with success", *filename)
		}

	default:
		fmt.Println("expected 'import' subcommand")
		os.Exit(1)
	}

	os.Exit(0)
}

func handleUpload(fileName string) (string, error) {

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Can't read file:", fileName)
		panic(err)
	}

	// Upload file to S3
	uploadFileName := "suffix_" + fileName
	url, err := s3client.UploadFile(context.Background(), "id-generator", uploadFileName, file)
	if err != nil {
		return "", err
	}

	fmt.Println("Uploaded Object Path : " + url)

	return url, nil
}
