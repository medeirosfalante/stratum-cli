package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rafaeltokyo/stratum-cli/terminal"
	"github.com/rafaeltokyo/stratum-sdk-go/stratumsdk"
)

func main() {
	err := godotenv.Load("./st.env")
	if err != nil {
		log.Fatal("Error loading ~/st.env file")
	}
	client := stratumsdk.Initial(os.Getenv("STRATUM_APIUSER"), os.Getenv("STRATUM_APIKEY"), false)
	cli := terminal.CLI{Sclient: client}
	cli.Run()

}
