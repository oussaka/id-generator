package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"id-generator/cmd/cloudStorage/s3client"
	"id-generator/handlers"
	"id-generator/models"
	"id-generator/repositories"
	"id-generator/routes"
	"id-generator/services"
	"id-generator/storage"
	"log"
	"path"
	"strings"

	"os"

	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
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

	fmt.Println("filename " + *filename)

	extension := path.Ext(*filename) //obtain the extension of file
	if strings.ToLower(extension) != ".csv" {
		fmt.Println("The file must be a csv.")

		os.Exit(1)
	}

	switch os.Args[1] {
	case "import":
		downloadedFile, err := handleDownload(*filename)
		if err != nil {
			if *useColor {
				services.Colorize(services.ColorRed, err.Error())
			} else {
				fmt.Printf("Download Object error : %s", err.Error())
			}

			os.Exit(1)
		}

		if *useColor {
			services.Colorize(services.ColorGreen, fmt.Sprintf("==> File %s downloaded with success", *filename))
		} else {
			fmt.Printf("==> File %s uploaded with success", *filename)
		}
		userCreatedCounter := importProcess(downloadedFile)

		if *useColor {
			services.Colorize(services.ColorBlue, fmt.Sprintf("==> %d accounts created", userCreatedCounter))
		} else {
			fmt.Printf("==> %d accounts created", userCreatedCounter)
		}

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

func handleDownload(object string) (string, error) {

	fileName := "testobject_local.csv"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Failed to create file", err)

		return "", err
	}
	defer file.Close()

	// Download file from S3
	_, err = s3client.DownloadFile(context.Background(), "id-generator", object, file, true)
	if err != nil {
		return "", err
	}

	fmt.Println("Downlod Object Path : " + fileName)

	return fileName, nil
}

func importProcess(csvfile string) int {

	var userCreatedCount int

	file, err := os.Open(csvfile)
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records")
	}

	var colIndexEmail, colIndexName int
	var header = records[0]
	for index, col := range header {
		if col == "email" {
			colIndexEmail = index
		}
		if col == "First_name" {
			colIndexName = index
		}
	}

	if colIndexEmail == 0 || colIndexName == 0 {
		panic("email or user name not found in csv file")
	}

	for index, row := range records {
		if index == 0 {
			continue
		}

		email := row[colIndexEmail]
		username := row[colIndexName]

		if email != "" {
			filter := make(map[string]interface{})
			filter["email"] = email
			_, err := repositories.FindUserBy(filter)

			if err == mongo.ErrNoDocuments {
				newUser, _ := repositories.CreateUser(*models.NewUser(username, email, ""))
				userCreatedCount++
				fmt.Printf("---> user created : %v\t\n", newUser)
			}
		}
	}

	return userCreatedCount
}
