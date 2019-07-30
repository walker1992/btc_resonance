package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/walker1992/btc_resonance/anchors"
	"io/ioutil"
	"os"

	"log"
	"time"
)

type Config struct {
	BtcPoolAccount string
	BtcPoolAddress string

	EthPoolAddress string

	AnchorsServerUrl string
	AnchorsClientUrl string

	EthClientUrl string
	BtcClientUrl string
	RpcUser      string
	RpcPassword  string

	AnchorsAccounts []string
	AnchorsKeys     []string
}

func ParseConfig(path string) Config {
	configFile, err := os.Open(path)
	if err != nil {
		panic("Config file is missing: " + err.Error())
	}

	configStr, err := ioutil.ReadAll(configFile)
	if err != nil {
		panic("Parse config file err: " + err.Error())
	}
	var config Config
	if err := json.Unmarshal(configStr, &config); err != nil {
		panic("Parse config file err: " + err.Error())
	}
	return config
}

func main() {
	configPath := flag.String("conf", "./config.json", "config file path, default './config.json'")
	flag.Parse()
	config := ParseConfig(*configPath)
	//fmt.Printf("the config :%v",config)

	accounts := make(map[common.Address]bool)
	for _, account := range config.AnchorsAccounts {
		accounts[common.HexToAddress(account)] = true
	}

	resonance := &anchors.Resonance{
		Trades:           make(map[common.Hash]*anchors.TradeWithSig),
		Anchors:          accounts,
		ReceiveTradeChan: make(chan *anchors.Trade),
		SyncTradeChan:    make(chan *anchors.Trade),
	}

	tradeServer := anchors.TradeServer{
		ReceiveTradeChan: resonance.ReceiveTradeChan,
	}

	err := anchors.StartServe(config.AnchorsServerUrl, &tradeServer)
	if err != nil {
		log.Fatal(err)
	}
	go resonance.TradeLoop(context.Background())

	anchorsClient, err := anchors.NewClient(config.AnchorsClientUrl)
	if err != nil {
		log.Fatal(err)
	}

	r := anchors.NewTrade(common.HexToHash("6122ccc16301d4452b84b6c04b14a097fb5d890df638862a8cc15af93bcbdbf8"), common.HexToAddress("0x387eeb1ec3e79e593f0fd3c05d1f1284045c3d37"), 100)

	for i := 0; i < 10; i++ {
		r.SignMessage(config.AnchorsKeys[0])
		err := anchorsClient.ResonanceTrade(r.Message())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("NewResonanceTrade: %v\n", r.Sig)
		time.Sleep(5 * time.Second)

		r.SignMessage(config.AnchorsKeys[1])
		err = anchorsClient.ResonanceTrade(r.Message())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("NewResonanceTrade: %v\n", r.Sig)
		time.Sleep(5 * time.Second)

		r.SignMessage(config.AnchorsKeys[2])
		err = anchorsClient.ResonanceTrade(r.Message())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("NewResonanceTrade: %v\n", r.Sig)
		time.Sleep(5 * time.Second)
	}

}
