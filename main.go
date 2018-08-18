package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"fmt"
	"time"
	"math/big"
	"C"
)

const rpcUrl = "rpc url here."

func GetBalance(hexAddress string, unit string) (error,string) {


	client,err := ethclient.Dial(rpcUrl)
	if err != nil {
		return fmt.Errorf("dial error: %v",err),"0"
	}
	address := common.HexToAddress(hexAddress)

	d := time.Now().Add(5000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(),d)
	defer cancel()
	balance, err := client.BalanceAt(ctx, address,nil)
	if err != nil {
		return fmt.Errorf("read error: %v",err),"0"
	}

	result := ""

	if unit == "NEW" {
		divV := big.NewFloat(0.0)
		divV.SetString("1000000000000000000")

		balanceOfNew := big.NewFloat(0.0)
		balanceOfNew.SetInt(balance)

		balanceOfNew.Quo(balanceOfNew, divV)

		result = balanceOfNew.Text('f',6)
	} else {
		result = balance.String()
	}

	return nil,result
}

//go build -o GoLibraryExample.dll -buildmode=c-shared main.go

//export GoGetBalance
func GoGetBalance(address *C.char, unit *C.char) (C.int,*C.char) {
	GoAddress := C.GoString(address)
	GoUnit := C.GoString(unit)
	err, balance := GetBalance(GoAddress,GoUnit)
	if err != nil {
		return 1,C.CString(err.Error())
	} else {
		return 0,C.CString(balance)
	}
}


func main() {
	address := "0x4E074bE4bc31DE624A3Df6035103747c3c5539c5"
	err,balance := GetBalance(address,"NEW")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(balance)
	}
}