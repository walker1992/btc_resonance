package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
)

//default:
//"sb1qg6ydpvdx8hhdrtvfs46zx0wx049nyht8xjn7zr",
//"SVBCk1pR524saYZWeUMUiurkR7hj65QeKV",
//"SYAFazGGjch4Afni6AgsgJMACZrtAj1cix"
//
//kingsley:
//"SXavGjPTKSEo7JZoW6gsxu684NTmJmqts1",
//"Sgofz1WsD8oYYkzaJFaHjT7d7f6oMP8JEE"
//
//walker:
//"SQp9JS8c94hCeZYV9UqxogmdGwVtMJbKBo",
//"SdEFkyipUUzNTtRDjVBWK5ECJkRZRniGMb"

func main() {
	client := initClient()
	err := client.WalletPassphrase("walker", 600)
	if err != nil {
		log.Fatal(err)
	}

	client.ListAccounts()
	transferFrom(client, "default", "SdEFkyipUUzNTtRDjVBWK5ECJkRZRniGMb", 9.9)

	//fmt.Println(transfer(client, "SQp9JS8c94hCeZYV9UqxogmdGwVtMJbKBo", 0.5))
	generate(client)
	//fmt.Println(getBalanceByAccount(client,"default"))
	//fmt.Println(getBalanceByAccount(client,"kingsley"))
	//fmt.Println(getBalanceByAccount(client,"walker"))
	//addresses,err :=client.GetAccountAddress("walker")
	//fmt.Println(addresses,err)

}

func initClient() *rpcclient.Client {
	ntfnHandlers := rpcclient.NotificationHandlers{
		OnAccountBalance: func(account string, balance btcutil.Amount, confirmed bool) {
			log.Printf("New balance for account %s: %v", account,
				balance)
		},
		OnRecvTx: func(transaction *btcutil.Tx, details *btcjson.BlockDetails) {
			log.Printf("RecvTx: %v,%v", transaction, details)
		},
		OnTxAccepted: func(hash *chainhash.Hash, amount btcutil.Amount) {
			log.Printf("TxAccepted : %v,%v", hash, amount)
		},
	}
	certHomeDir := btcutil.AppDataDir("btcwallet", false)
	certs, err := ioutil.ReadFile(filepath.Join(certHomeDir, "rpc.cert"))
	if err != nil {
		log.Fatal(err)
	}

	connCfg := &rpcclient.ConnConfig{
		Host:         "localhost:18554",
		Endpoint:     "ws",
		User:         "walker",
		Pass:         "12345",
		Certificates: certs,
	}
	client, err := rpcclient.New(connCfg, &ntfnHandlers)
	if err != nil {
		log.Fatal(err)
	}
	address, _ := btcutil.DecodeAddress("SQp9JS8c94hCeZYV9UqxogmdGwVtMJbKBo", &chaincfg.SimNetParams)
	if err := client.NotifyReceived([]btcutil.Address{address}); err != nil {
		log.Fatal(err)
	}

	if err := client.NotifyReceived([]btcutil.Address{address}); err != nil {
		log.Fatal(err)
	}
	if err := client.NotifyNewTransactions(false); err != nil {
		log.Fatal(err)
	}

	return client
}

func transferFrom(client *rpcclient.Client, account, addr string, amount float64) bool {
	address, _ := btcutil.DecodeAddress(addr, &chaincfg.SimNetParams)
	btcAmount, _ := btcutil.NewAmount(amount)

	hash, err := client.SendFrom(account, address, btcAmount)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return false
	} else {
		log.Println("txid:", hash)
		return true
	}

}

func transfer(client *rpcclient.Client, addr string, amount float64) bool {
	err := client.WalletPassphrase("walker", 600)
	if err != nil {
		log.Fatal(err)
	}
	address, _ := btcutil.DecodeAddress(addr, &chaincfg.SimNetParams)
	btcAmount, _ := btcutil.NewAmount(amount)
	hash, err := client.SendToAddress(address, btcAmount)
	if err != nil {
		log.Fatal(err)
		return false
	} else {
		log.Println("txid:", hash)
		return true
	}
}

func generate(client *rpcclient.Client) bool {
	_, err := client.Generate(12)
	if err != nil {
		return false
	} else {
		return true
	}
}

func getBalanceByAccount(client *rpcclient.Client, account string) btcutil.Amount {
	err := client.WalletPassphrase("walker", 600)
	if err != nil {
		log.Fatal(err)
	}
	amount, err := client.GetBalance(account)
	if err != nil {
		return 0
	} else {
		return amount
	}
}
