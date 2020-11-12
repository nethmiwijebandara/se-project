package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Both source directory and destination folder are here.
	files, err := Unzip("sample2.zip", "unzipped")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Unzipped the following files:\n" + strings.Join(files, "\n"))

	//Read line by line
	f, err := os.Open("./unzipped/sample2.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	fmt.Println(txtlines[0])

	//Connecting to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://neth98:neth123@test.dq1l7.mongodb.net/quickstart?retryWrites=true&w=majority"))

	if err != nil {
		log.Fatal(err)
	}

	connectMongo(*client, txtlines, err)
	fmt.Println("Successful")

}
