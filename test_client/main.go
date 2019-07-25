package main

import (
	"fmt"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
	"io/ioutil"
	"log"
	"path/filepath"
)

// default account address: SUsbWJE1VMREXFUv8KcBmke7cPsSXaFvnE
// main account address: SiAXFyKMF6afKEhoTYAPSYGWxbCRUddqTh
// ss account address: Sad47SLU1hsgiQb2Ew6jf39PfqhJzEjERZ SjPHLr6MUn7fjjBuZK6a3pj9NFgYGvpu25

func main() {
	client := initClient()
	err := client.WalletPassphrase("walker", 600)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(client.ListAccounts())
	fmt.Println(transferFrom(client, "default", "SdEFkyipUUzNTtRDjVBWK5ECJkRZRniGMb", 9.9))

	//for i :=0;i<10;i++{
	//	transferFrom(client,"default", "SdEFkyipUUzNTtRDjVBWK5ECJkRZRniGMb", float64(0.001))
	//
	//
	//	transferFrom(client,"default", "Sgofz1WsD8oYYkzaJFaHjT7d7f6oMP8JEE", float64(0.001))
	//}

	//fmt.Println(transferFrom("default", "SQp9JS8c94hCeZYV9UqxogmdGwVtMJbKBo", 4))
	//fmt.Println(transfer(client,"SdEFkyipUUzNTtRDjVBWK5ECJkRZRniGMb", 0.5))
	generate(client)
	//fmt.Println(getBalanceByAccount(client,"default"))
	//fmt.Println(getBalanceByAccount(client,"kingsley"))
	//fmt.Println(getBalanceByAccount(client,"walker"))
	//addresses,_ :=client.GetAccountAddress("kingsley")
	//fmt.Println(addresses)

	//txHash,_ :=chainhash.NewHashFromStr("cca8891b5ef493f8225a63ad962d67eb192e6f5f07fb0cce85a97b8f51aa55d8")
	//
	//
	//vout,err :=client.GetTxOut(txHash,0,false)
	//if err !=nil{
	//	log.Fatal(err)
	//}
	//
	//log.Printf("the out is amount: %v,%v",vout.Value,vout.ScriptPubKey.Addresses)
	//
	//
	//txHash,_ =chainhash.NewHashFromStr("9c13654d39af9cdd5fd1111dffc23eafa8c5efca1957a37616126531c79a57f0")
	//
	//
	//vout,err =client.GetTxOut(txHash,1,false)
	//if err !=nil{
	//	log.Fatal(err)
	//}
	//
	//log.Printf("the out is amount: %v,%v",vout.Value,vout.ScriptPubKey.Addresses)

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
