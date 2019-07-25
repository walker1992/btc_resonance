package main

import (
	"fmt"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"io/ioutil"
	"log"
	"path/filepath"
)

const ResonancePoolAccount = "walker"
const ResonancePoolAddr = "SdEFkyipUUzNTtRDjVBWK5ECJkRZRniGMb"

const rpchost = "localhost:18556"
const rpcuser = "walker"
const rpcpassword = "12345"

func main() {
	// Only override the handlers for notifications you care about.
	// Also note most of these handlers will only be called if you register
	// for notifications.  See the documentation of the rpcclient
	// NotificationHandlers type for more details about each handler.
	ntfnHandlers := rpcclient.NotificationHandlers{
		OnFilteredBlockConnected: func(height int32, header *wire.BlockHeader, txns []*btcutil.Tx) {
			//log.Printf("Block connected: %v (%d)",
			//	header.BlockHash(), height)

			getTransactionByBlock(header.BlockHash())

		},
		OnFilteredBlockDisconnected: func(height int32, header *wire.BlockHeader) {
			log.Printf("Block disconnected: %v (%d) %v",
				header.BlockHash(), height, header.Timestamp)
		},

		OnTxAccepted: func(hash *chainhash.Hash, amount btcutil.Amount) {
			log.Printf("TxAccepted : %v,%v", hash, amount)
			getTransaction(hash)

		},
	}

	// Connect to local btcd RPC server using websockets.
	btcdHomeDir := btcutil.AppDataDir("btcd", false)

	log.Println(filepath.Join(btcdHomeDir, "rpc.cert"))
	certs, err := ioutil.ReadFile(filepath.Join(btcdHomeDir, "rpc.cert"))
	if err != nil {
		log.Fatal(err)
	}
	connCfg := &rpcclient.ConnConfig{
		Host:         rpchost,
		Endpoint:     "ws",
		User:         rpcuser,
		Pass:         rpcpassword,
		Certificates: certs,
	}
	client, err := rpcclient.New(connCfg, &ntfnHandlers)
	if err != nil {
		log.Fatal(err)
	}

	// Register for block connect and disconnect notifications.
	if err := client.NotifyBlocks(); err != nil {
		log.Fatal(err)
	}
	log.Println("NotifyBlocks: Registration Complete")

	if err := client.NotifyNewTransactions(false); err != nil {
		log.Fatal(err)
	}
	log.Println("NotifyNewTransactions: Registration Complete") // Get the current block count.
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block count: %d", blockCount)
	log.Println("Wait for Client shutdown ...")

	client.WaitForShutdown()
}

func getTransactionByBlock(blockHash chainhash.Hash) {
	client := NewClient()
	block, err := client.GetBlock(&blockHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(">>>>>>>>>>>>>>>>>>>.\n%d tx in this block\n", len(block.Transactions))
	for _, tx := range block.Transactions {
		//log.Println("vout counnt:", len(tx.TxOut))
		hash := tx.TxHash()
		//fmt.Printf("the txid : %v\n", hash)
		for i := 0; i < len(tx.TxOut); i++ {
			txout, err := client.GetTxOut(&hash, uint32(i), false)
			if err != nil {
				log.Fatal(err)
			}
			var receiveAddr []string
			if txout.Value != 0 { //TODO txout.ScriptPubKey.Addresses !=nil??
				receiveAddr = txout.ScriptPubKey.Addresses
			}

			if receiveAddr[0] == ResonancePoolAddr {
				btcAmount, _ := btcutil.NewAmount(txout.Value)

				fmt.Println("BTC Resonance NOW!!")
				fmt.Printf("the amount is          %v\n", btcAmount)
				fmt.Printf("the receive address is %v\n", receiveAddr[0])
			}
		}

	}
}

func getTransaction(txhash *chainhash.Hash) {
	client := NewClient()

	tx, err := client.GetRawTransaction(txhash)
	if err != nil {
		log.Println(err)
		return
	}

	hash := tx.Hash()
	//fmt.Printf("the txid : %v\n", hash)
	//log.Printf("txout count %v", len(tx.MsgTx().TxOut))
	for i := 0; i < len(tx.MsgTx().TxOut); i++ {
		txout, err := client.GetTxOut(hash, uint32(i), false)
		if err != nil || txout == nil {
			//log.Println(err)
			continue
		}

		var receiveAddr []string
		if txout.Value != 0 { //TODO txout.ScriptPubKey.Addresses !=nil??
			receiveAddr = txout.ScriptPubKey.Addresses
		}

		if receiveAddr[0] == ResonancePoolAddr {
			btcAmount, _ := btcutil.NewAmount(txout.Value)

			fmt.Println("BTC Resonance NOW!! from OnTxAccepted")
			fmt.Printf("the amount is          %v\n", btcAmount)
			fmt.Printf("the receive address is %v\n", receiveAddr[0])
		}
	}

}

func NewClient() *rpcclient.Client {
	// Connect to local btcd RPC server using websockets.
	btcdHomeDir := btcutil.AppDataDir("btcd", false)
	certs, err := ioutil.ReadFile(filepath.Join(btcdHomeDir, "rpc.cert"))
	connCfg := &rpcclient.ConnConfig{
		Host:         rpchost,
		Endpoint:     "ws",
		User:         rpcuser,
		Pass:         rpcpassword,
		Certificates: certs,
	}
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
