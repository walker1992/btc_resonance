package main

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/walker1992/btc_resonance/contract"
	"github.com/walker1992/btc_resonance/log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var rpcCCC = "http://127.0.0.1:8545"

//0xCdce1E6B706e7e4Ce269Ba1D4aaAED2fdc78690a
const senderKey = "b80dbf638b9128e54f3222d2b6d3213d45521d49bb6317abdf34b219a55204b7"

//0xbF00C30b93d76ab3D45625645b752b68199c8221
const receiverKey = "08cd4fde21e980c7d05afa3b0d4d27534e646be3cc3a67b303b055d1166cbae3"

//address 0x51236B7C02956B698185B755B1f3a6a0D4BF20E5
const Key1 = "5071ddb8d4853b1aeffe413bdbea81489ee9ef033c60859ef517eb2c5f4c3888"

//address 0xCbf60c00Cda6d16BBDd4A286748Ff0f640EB4d84
const Key2 = "9578b6657475873003c20d61b4574d26845d89aad7c2b3938d39878b1b331211"

//address 0x47CE74D1A1B395117F1B2e9d79072299415F7503
const Key3 = "607a7307383a6c28c003a2b6e419b4f070160f9d15fce028471d43c5f030d4ed"

const ResonanceEthPoolAddr = "0xd3c91803b4816f79dc5af0090fd0fb6d5a44d412"

func main() {
	rpcPath := flag.String("p", "http://127.0.0.1:8545", "rpc host of CCC chain")
	flag.Parse()

	clientCcc, err := NewEthClient(*rpcPath)
	if err != nil {
		log.Crit(err.Error())
	}

	value, err := clientCcc.Client.BalanceAt(context.Background(), common.HexToAddress("0xa0191488709119f87ebb65a7f2ef8dec76a9547b"), nil)
	if err != nil {
		log.Crit(err.Error())
	}
	log.Info("the value from CCC chain: %s", value.String())

	//deployContract(clientCcc)

}

type EthClient struct {
	Client *ethclient.Client
	Path   string
}

func NewEthClient(path string) (*EthClient, error) {
	client, err := ethclient.Dial(path)
	if err != nil {
		return nil, err
	}

	return &EthClient{
		Client: client,
		Path:   path,
	}, nil
}

func (c *EthClient) CalResonance(btcAmount int64) (ethAmount *big.Int, err error) {
	value, err := c.Client.BalanceAt(context.Background(), common.HexToAddress(ResonanceEthPoolAddr), nil)
	if err != nil {
		return nil, err
	}
	log.Info("the value from CCC chain: %s", value.String())

	//poolBalance :=value.Int64()
	//按照规则开始计算eth上的发币量

	return
}

func makeAuth(private string, client *ethclient.Client, value int64) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(private)
	if err != nil {
		log.Crit("Failed to parse private key: %s", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Crit("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Crit("Failed to get nonce: %s", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Crit("Failed to get gasPrice: %s", err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(value)  //in wei
	auth.GasLimit = uint64(3000000) //in uints

	gasPriceInt, _ := big.NewInt(0).SetString(gasPrice.String(), 10)
	auth.GasPrice = gasPriceInt

	return auth
}

func deployContract(client *ethclient.Client) (contractAddr common.Address, trx *types.Transaction, err error) {
	auth := makeAuth(Key1, client, 1000000)

	log.Info("Deploy contract...")

	address, tx, _, err := resonance.DeployResonance(auth, client)
	if err != nil {
		return common.Address{}, nil, err
	}

	log.Info("contract address = %s", address.Hex())
	log.Info("transaction hash = %s", tx.Hash().Hex())
	log.Info("trx owner %s", auth.From.String())
	return address, tx, nil
}

type ResonanceTrade struct {
	BtcHash  common.Hash    `json:"btcHash" gencodec:"required"`
	Receiver common.Address `json:"receiver" gencodec:"required"`
	Amount   *big.Int       `json:"amount" gencodec:"required"`

	sig []byte
	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`
}

type ResonanceMsg struct {
	BtcTxId common.Hash `json:"btcTxId" gencodec:"required"`
	Sig     []byte      `json:"sig" gencodec:"required"`
}

func NewResonanceTrade(btcHash, receiver string, amount int64) *ResonanceTrade {
	r := ResonanceTrade{
		BtcHash:  common.HexToHash(btcHash),
		Receiver: common.HexToAddress(receiver),
		Amount:   big.NewInt(amount),
	}
	return &r
}

func (m *ResonanceTrade) SignMessage(private string) {
	var data []byte
	data = append(data, m.BtcHash.Bytes()...)
	data = append(data, m.Receiver.Bytes()...)
	data = append(data, common.LeftPadBytes(m.Amount.Bytes(), 32)...)

	hash := crypto.Keccak256Hash(data)

	privateKey, err := crypto.HexToECDSA(private)
	if err != nil {
		log.Crit("Failed to parse private key: %s", err)
	}

	sig, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		panic(fmt.Sprintf("wrong size for signature: got %d, want 65", len(sig)))
	}
	if len(sig) != 65 {
		panic(fmt.Sprintf("wrong size for signature: got %d, want 65", len(sig)))
	}

	m.sig = sig
	m.R = new(big.Int).SetBytes(sig[:32])
	m.S = new(big.Int).SetBytes(sig[32:64])
	m.V = new(big.Int).SetBytes([]byte{sig[64] + 27})
}

func (m *ResonanceTrade) Message() *ResonanceMsg {
	return &ResonanceMsg{
		BtcTxId: m.BtcHash,
		Sig:     m.sig,
	}
}

func CreateResonanceTrade() {
	r := NewResonanceTrade("6122ccc16301d4452b84b6c04b14a097fb5d890df638862a8cc15af93bcbdbf8", "0x387eeb1ec3e79e593f0fd3c05d1f1284045c3d37", 100)
	fmt.Println(r.BtcHash.String())
	r.SignMessage(Key1)
	fmt.Printf("NewResonanceTrade: %v\n", r)

	fmt.Println(hexutil.Encode(common.LeftPadBytes(r.R.Bytes(), 32)))
	fmt.Println(hexutil.Encode(common.LeftPadBytes(r.S.Bytes(), 32)))
	r.SignMessage(Key2)
	fmt.Printf("NewResonanceTrade: %v\n", r)

	fmt.Println(hexutil.Encode(common.LeftPadBytes(r.R.Bytes(), 32)))
	fmt.Println(hexutil.Encode(common.LeftPadBytes(r.S.Bytes(), 32)))
	r.SignMessage(Key3)
	fmt.Printf("NewResonanceTrade: %v\n", r)

	fmt.Println(hexutil.Encode(common.LeftPadBytes(r.R.Bytes(), 32)))
	fmt.Println(hexutil.Encode(common.LeftPadBytes(r.S.Bytes(), 32)))
}
