package terminal

import (
	"fmt"
	"io/ioutil"

	"github.com/rafaeltokyo/stratum-sdk-go/stratumsdk"
)

//HelpWithdrawCommandPrint list command support wallet list
func (cli *CLI) HelpWithdrawCommandPrint() {
	fmt.Printf("------------------- Help to use Withdraw -------------------\n")
	fmt.Println(`example command withdraw -dest='39oqLQehtJBssE2cjRTc5nUtmPQNywTqmM' -amount=0.0010 -eid=100 -otp='121212' -walletId=102 -desc='test'`)
}

//RequestWithdraw Request a new withdraw
func (cli *CLI) RequestWithdraw(WalletId *int, Amount *float64, Desc *string, Dest *string, Otp *string, Eid *int) {
	withdrawPayload := &stratumsdk.WithdrawsPayload{
		WalletId:        *WalletId,
		OperationAmount: *Amount,
		OperationDesc:   *Desc,
		OperationEid:    *Eid,
		OperationOtp:    *Otp,
		DestAddress:     *Dest,
	}
	fmt.Printf("withdrawPayload %#v", withdrawPayload)
	withdrawReq, apiErr, err := cli.Sclient.Withdraws().Crypto(withdrawPayload)
	fmt.Printf("withdrawReq %#v", withdrawReq)
	if err != nil {
		fmt.Printf("sdk error:  %s ", err.Error())
	}
	if apiErr != nil {
		fmt.Printf("apiError: %s ", apiErr.Data)
	}
	fmt.Printf("\n---------------------- WithDraw Request -----------------------\n")
	printWithdraw(withdrawReq)
	fmt.Println("--------------------------------------------------------------")

}

func printWithdraw(wd *stratumsdk.WithdrawsData) {
	b, err := ioutil.ReadFile("./template/withdrawReq.txt")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf(string(b)+"\n",
		wd.DestAddress,
		wd.OperationAmount,
		wd.OperationDesc,
		wd.OperationEid,
		wd.OperationOtp,
		wd.WalletId,
	)
}
