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
	fmt.Println("createwalletAddress -walletID=WALLETID  - create walletaAddress with walletID")
	fmt.Println(`"listWalletAddress - list all walletAddress commands: -query=ObjectQuery execute query  | -h - help use walletAddressList`)

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

	// Create wallet commands
	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)

	createWalletGroupID := createWalletCmd.Int("groupID", 0, "pass group id for create HD wallet grouped")
	createWalleWalletEid := createWalletCmd.Int("walletEid", 0, "pass walletid your custom id")
	// List Wallet commands
	listWalletsCmd := flag.NewFlagSet("listWallets", flag.ExitOnError)
	listWalletsObjectQuery := listWalletsCmd.String("query", "", "pass ObjectQuery for query in list")
	listWalletsHelp := listWalletsCmd.Bool("h", false, "help to use walletList")
	// Create WalletAddress commands
	createWalletAddressCmd := flag.NewFlagSet("createwalletAddress", flag.ExitOnError)
	createWalletAddressID := createWalletAddressCmd.Int("walletID", 0, "pass walletID for generate address")
	createWalletAddressCurrency := createWalletAddressCmd.String("currency", "BTC", "pass Currency for generate address")
	// List WalletAddress commands
	listWalletAddressCmd := flag.NewFlagSet("listWalletAddresss", flag.ExitOnError)
	listWalletAddressObjectQuery := listWalletAddressCmd.String("query", "", "pass ObjectQuery for query in list")
	listWalletAddressHelp := listWalletAddressCmd.Bool("h", false, "help to use walletList")
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
	case "listWalletAddress":
		err := listWalletAddressCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createwalletAddress":
		err := createWalletAddressCmd.Parse(os.Args[2:])
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

	if createWalletAddressCmd.Parsed() {
		if *createWalletAddressID == 0 || *createWalletAddressCurrency == "" {
			createWalletAddressCmd.Usage()
			os.Exit(1)
		}
		cli.createWalletAddress(createWalletAddressID, createWalletAddressCurrency)
	}

	if listWalletAddressCmd.Parsed() {
		if *listWalletAddressHelp {
			cli.HelpWalletAddressListCommandPrint()
			os.Exit(0)
		}
		cli.listWalletAddress(listWalletAddressObjectQuery)
	}

}
