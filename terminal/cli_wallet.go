package terminal

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pquerna/ffjson/ffjson"

	"github.com/rafaeltokyo/stratum-sdk-go/stratumsdk"
)

func (cli *CLI) createWallet(groupID *int, WalletEid *int) {

	walletCreatePayload := &stratumsdk.WalletPayload{
		WalletEid:     *WalletEid,
		WalletGroupId: *groupID,
		WalletLabel:   "stratum-cli",
		Currency:      "BTC",
		WalletType:    "checking",
	}

	cWallet, apiErr, err := cli.Sclient.Wallets().Create(walletCreatePayload)

	if err != nil {
		fmt.Printf("sdk error:  %s ", err.Error())
		os.Exit(1)
	}
	if apiErr != nil {
		fmt.Printf("apiError: %s ", apiErr.Data)
		os.Exit(1)
	}

	printWallet(cWallet)

}

func printWallet(w *stratumsdk.WalletData) {
	b, err := ioutil.ReadFile("./template/walletItem.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf(string(b)+"\n",
		w.WalletId,
		w.WalletEid,
		w.WalletLabel,
		w.WalletBalance,
		w.WalletGroupId,
		w.WalletGroupLabel,
		w.WalletGroupEid,
		w.WalletType,
		w.Currency,
		w.CurrencyUnitDigits,
	)

}

func (cli *CLI) listWallet(objectQuery *string) {
	walletListPayload := &stratumsdk.WalletsListPayload{}
	if *objectQuery != "" {
		err := ffjson.Unmarshal([]byte(*objectQuery), &walletListPayload)
		if err != nil {
			panic(err)
		}
	}
	wallets, apiErr, err := cli.Sclient.Wallets().List(walletListPayload)
	if err != nil {
		fmt.Printf("sdk error:  %s ", err.Error())
	}
	if apiErr != nil {
		fmt.Printf("apiError: %s ", apiErr.Data)
	}
	totalFind := len(*wallets)
	if totalFind == 0 {
		fmt.Printf("0 register find \n")
		os.Exit(0)
	}
	fmt.Printf("------------------------ Wallet List ------------------------\n")
	fmt.Printf("---------------------- find %d wallets ----------------------\n", totalFind)
	for _, item := range *wallets {
		printWallet(&item)
	}
	println("--------------------------------------------------------")

}

//HelpWalletListCommandPrint list command support wallet list
func (cli *CLI) HelpWalletListCommandPrint() {
	fmt.Printf("------------------- Help to use Wallet List -------------------\n")
	fmt.Println(`example command listWallets -query='{"wallet_group_id":10}'`)
	fmt.Println(`fields accept {
   "wallet_eid":10
   "wallet_balance_min":12,
   "wallet_balance_max":12,
   "wallet_group_eid":12,	
   "wallet_group_id":12,
   "wallet_type":"checking",
   "currency":"BTC,
 }`)
}
