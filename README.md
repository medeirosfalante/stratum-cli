

![Stratum.hk](stratum_go.png "Stratum.hk")

stratum-cli is a commandline for stratum wallet api


## How to install:
``` bash
go get github.com/pquerna/ffjson
go get github.com/rafaeltokyo/stratum-cli
```


## Add API User and Key:
``` bash
cp .env_example .env
# then add user and key to .env file
```

## How to run:
``` bash
cd $GOPATH/src/github.com/rafaeltokyo/stratum-cli
 ./stratum-cli [command] [subcommand]

# Examples:
# List deposits:
./stratum-cli listOperations -query='{"operation_type":"deposit"}'

# List wallets in wallet group id 10:
./stratum-cli listWallets -query='{"wallet_group_id":10}'

# List all BTC addresses:
./stratum-cli listWalletAddress -query='{"currency":"BTC"}'
```