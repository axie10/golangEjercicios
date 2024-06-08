package main

import (
	"fmt"
	"flag"

	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/jsonrpc"
)

var(
	client *jsonrpc.Client
	myAddress ethgo.Address
)

func initValues() error {

	client, err := jsonrpc.NewClient("https://mainnet.infura.io")
	if err != nil {
		return err
	}

	var addressStr string
	flag.StringVar(&addressStr, "address", "0x...", "The public address that you want to check")

	flag.Parse()
	if addressStr == "0x..." {
		flag.Usage()
		return fmt.Errorf("address is mandatory")
	}

	myAddress = ethgo.HexToAddress(addressStr)
	return nil

}

func allTransactions(currentBlockHash []) {

	eth := client.Eth()
	txs := make([]*ethgo.Transaction, 0)

	for currentBlockHash != lastBlockHash {
		block, err := eth.GetBlockByHash(currentBlockHash, true)
		if err != nil {
			return txs, err
		}

		currentBlockHash = block.ParentHash
		for _, tx := range block.Transactions{
			if tx.From == myAddress || (tx.To != nil && *tx.To == myAddress) {
				txs = append(txs, tx)
			}
		}
	}

	return txs, nil

}

func main() {



}

