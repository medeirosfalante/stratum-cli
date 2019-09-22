package terminal

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pquerna/ffjson/ffjson"
	"github.com/rafaeltokyo/stratum-sdk-go/stratumsdk"
)

func printOperation(op *stratumsdk.OperationData) {
	b, err := ioutil.ReadFile("./template/operationItem.txt")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf(string(b)+"\n",
		op.OperationId,
		op.WalletId,
		op.OperationAmount,
		op.OperationTamount,
		op.OperationFee,
		op.OperationDesc,
		op.OperationEid,
		op.OperationEtxid,
		op.OperationTs,
		op.OperationUpdTs,
		op.OperationConf,
		op.OperationConfreq,
		op.DestTypeData,
		op.OperationInfo,
		op.CurrencyUsdtrate,
		op.OperationStatus,
		op.OperationType,
		op.WalletEid,
		op.WalletGroupId,
		op.WalletGroupEid,
		op.WalletLabel,
		op.WalletType,
		op.Currency,
		op.CurrencyUnit,
		op.CurrencyType,
		op.DestType,
		op.DirectionType,
	)
}

func printFees(op *stratumsdk.FeeData) {
	b, err := ioutil.ReadFile("./template/fees.txt")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf(string(b)+"\n",
		op.Currency,
		op.DestType,
		op.OperationFee,
		op.OperationType,
	)
}

func (cli *CLI) listOperations(objectQuery *string) {
	operationsListPayload := &stratumsdk.OperationPayload{}
	if *objectQuery != "" {
		err := ffjson.Unmarshal([]byte(*objectQuery), &operationsListPayload)
		if err != nil {
			panic(err)
		}
	}
	operations, apiErr, err := cli.Sclient.Operations().List(operationsListPayload)
	if err != nil {
		fmt.Printf("sdk error:  %s ", err.Error())
	}
	if apiErr != nil {
		fmt.Printf("apiError: %s ", apiErr.Data)
	}
	totalFind := len(*operations)
	if totalFind == 0 {
		fmt.Printf("0 register find \n")
		os.Exit(0)
	}
	fmt.Printf("---------------------- Operations List -----------------------\n")
	fmt.Printf("--------------------- find %d Operations --------------------\n", totalFind)
	for _, item := range *operations {
		printOperation(&item)
	}
	println("--------------------------------------------------------")

}

func (cli *CLI) feesCurrency(currency *string) {
	feePayload := &stratumsdk.FeePayload{Currency: *currency}
	fees, apiErr, err := cli.Sclient.Operations().Fees(feePayload)
	if err != nil {
		fmt.Printf("sdk error:  %s ", err.Error())
	}
	if apiErr != nil {
		fmt.Printf("apiError: %s ", apiErr.Data)
	}
	fmt.Printf("---------------------- Fees  -----------------------\n")
	for _, item := range *fees {
		printFees(&item)
	}
	println("--------------------------------------------------------")

}

//HelpOperationsListCommandPrint list command support wallet list
func (cli *CLI) HelpOperationsListCommandPrint() {
	fmt.Printf("---------------- Help to use Operations List ----------------\n")
	fmt.Println(`example command listOperations -query='{"wallet_eid":10}'`)
	fmt.Println(`fields accept {
		"dest_type":  "intra",    // types: in,out,intra
		"direction_type": "intra",
		"operation_eid":12,
		"operation_status":"done"      // new,processing,done,failed
		"operation_ts_from":111111,    
		"operation_ts_to": 11111111,     
		"operation_type": "deposit",
		"operation_upd_ts_from":111111,
		"operation_upd_ts_to": 111111,
		"wallet_eid":12
		"wallet_group_eid":12
		"wallet_id":12
	}`)
}
