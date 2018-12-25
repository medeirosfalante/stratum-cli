package terminal

import (
	"flag"
	"fmt"
	"log"

	"github.com/rafaeltokyo/stratum-sdk-go/stratumsdk"

	"os"
)

// CLI responsible for processing command line arguments
type CLI struct {
	Sclient *stratumsdk.ApiClient
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("createwallet -groupID=GROUPID  - create wallet of GROUPID")
	fmt.Println(`"lisWallets - list all wallets commands: -query=ObjectQuery execute query  | -h - help use walletList`)

}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

// Run parses command line arguments and processes commands
func (cli *CLI) Run() {
	cli.validateArgs()

	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	listWalletsCmd := flag.NewFlagSet("listWallets", flag.ExitOnError)
	createWalletGroupID := createWalletCmd.Int("groupID", 0, "pass group id for create HD wallet grouped")
	createWalleWalletEid := createWalletCmd.Int("walletEid", 0, "pass walletid your custom id")
	listWalletsObjectQuery := listWalletsCmd.String("query", "", "pass ObjectQuery for query in list")
	listWalletsHelp := listWalletsCmd.Bool("h", false, "help to use walletList")

	switch os.Args[1] {
	case "createwallet":
		err := createWalletCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "listWallets":
		err := listWalletsCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if createWalletCmd.Parsed() {
		if *createWalletGroupID == 0 {
			createWalletCmd.Usage()
			os.Exit(1)
		}
		cli.createWallet(createWalletGroupID, createWalleWalletEid)
	}
	if listWalletsCmd.Parsed() {
		if *listWalletsHelp {
			cli.HelpWalletListCommandPrint()
			os.Exit(0)
		}
		cli.listWallet(listWalletsObjectQuery)
	}

}
