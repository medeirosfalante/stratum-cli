package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rafaeltokyo/stratum-cli/terminal"
	"github.com/rafaeltokyo/stratum-sdk-go/stratumsdk"
)

func main() {
	godotenv.Load("./.env")
	client := stratumsdk.Initial(os.Getenv("STRATUM_APIUSER"), os.Getenv("STRATUM_APIKEY"), false)
	cli := terminal.CLI{Sclient: client}
	cli.Run()

}
