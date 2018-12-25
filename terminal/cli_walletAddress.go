package terminal

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pquerna/ffjson/ffjson"
	"github.com/rafaeltokyo/stratum-sdk-go/stratumsdk"
)

func (cli *CLI) createWalletAddress(walletID *int, Currency *string) {

	walletCreatePayload := &stratumsdk.WalletAddressesPayload{
		WalletId:           *walletID,
		WalletAddressLabel: "stratum-cli-address",
		Currency:           *Currency,
	}

	cAWallet, apiErr, err := cli.Sclient.WalletsAddresses().Assign(walletCreatePayload)

	if err != nil {
		fmt.Printf("sdk error:  %s ", err.Error())
		os.Exit(1)
	}
	if apiErr != nil {
		fmt.Printf("apiError: %s ", apiErr.Data)
		os.Exit(1)
	}

	printWalletAddress(cAWallet)

}

func printWalletAddress(wa *stratumsdk.WalletAddressData) {
	b, err := ioutil.ReadFile("./template/walletAddresItem.txt")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf(string(b)+"\n",
		wa.WalletAddressEid,
		wa.WalletEid,
		wa.WalletAddressLabel,
		wa.WalletAddress,
		wa.CurrencyName,
		wa.Currency,
		wa.WalletLabel,
		wa.WalletBalance,
	)
}

func (cli *CLI) listWalletAddress(objectQuery *string) {
	walletAddressListPayload := &stratumsdk.WalletsAddressesListPayload{}
	if *objectQuery != "" {
		err := ffjson.Unmarshal([]byte(*objectQuery), &walletAddressListPayload)
		if err != nil {
			panic(err)
		}
	}
	walletsAddress, apiErr, err := cli.Sclient.WalletsAddresses().List(walletAddressListPayload)
	if err != nil {
		fmt.Printf("sdk error:  %s ", err.Error())
	}
	if apiErr != nil {
		fmt.Printf("apiError: %s ", apiErr.Data)
	}
	totalFind := len(*walletsAddress)
	if totalFind == 0 {
		fmt.Printf("0 register find \n")
		os.Exit(0)
	}
	fmt.Printf("---------------------- WalletAddress List -----------------------\n")
	fmt.Printf("--------------------- find %d walletsAddress --------------------\n", totalFind)
	for _, item := range *walletsAddress {
		printWalletAddress(&item)
	}
	println("--------------------------------------------------------")

}
