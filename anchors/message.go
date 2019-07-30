package anchors

import (
	"fmt"
	"math/big"

	"github.com/walker1992/btc_resonance/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Trade struct {
	BtcTxId  common.Hash    `json:"btcHash" gencodec:"required"`
	Receiver common.Address `json:"receiver" gencodec:"required"`
	Amount   *big.Int       `json:"amount" gencodec:"required"`

	hash common.Hash
	Sig  []byte
	//signer common.Address

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`
}

type ResonanceMsg struct {
	BtcTxId  common.Hash    `json:"btcTxId" gencodec:"required"`
	Receiver common.Address `json:"receiver" gencodec:"required"`
	Amount   *big.Int       `json:"amount" gencodec:"required"`
	Sig      []byte         `json:"sig" gencodec:"required"`
}

func NewTrade(btcHash common.Hash, receiver common.Address, amount int64) *Trade {
	return &Trade{
		BtcTxId:  btcHash,
		Receiver: receiver,
		Amount:   big.NewInt(amount),
	}
}

func NewTradeFromMsg(msg *ResonanceMsg) *Trade {
	t := &Trade{
		BtcTxId:  msg.BtcTxId,
		Receiver: msg.Receiver,
		Amount:   msg.Amount,
		Sig:      msg.Sig,
		V:        new(big.Int).SetBytes([]byte{msg.Sig[64] + 27}),
		R:        new(big.Int).SetBytes(msg.Sig[:32]),
		S:        new(big.Int).SetBytes(msg.Sig[32:64]),
	}

	var data []byte
	data = append(data, t.BtcTxId.Bytes()...)
	data = append(data, t.Receiver.Bytes()...)
	data = append(data, common.LeftPadBytes(t.Amount.Bytes(), 32)...)

	t.hash = crypto.Keccak256Hash(data)

	return t
}

func (t *Trade) SignMessage(private string) {
	var data []byte
	data = append(data, t.BtcTxId.Bytes()...)
	data = append(data, t.Receiver.Bytes()...)
	data = append(data, common.LeftPadBytes(t.Amount.Bytes(), 32)...)

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

	t.hash = hash
	t.Sig = sig
	t.R = new(big.Int).SetBytes(sig[:32])
	t.S = new(big.Int).SetBytes(sig[32:64])
	t.V = new(big.Int).SetBytes([]byte{sig[64] + 27})
}

func (t *Trade) getSigner() (common.Address, error) {
	recoveredPub, err := crypto.Ecrecover(t.hash.Bytes(), t.Sig)
	if err != nil {
		log.Warn("Ecrecover from %s is wrong!", t.hash.String())
		return common.Address{}, err
	}

	pubKey, err := crypto.UnmarshalPubkey(recoveredPub)
	if err != nil {
		log.Warn("UnmarshalPubkey is wrong,%s", err)
		return common.Address{}, err
	}
	signer := crypto.PubkeyToAddress(*pubKey)
	return signer, nil

}

func (t *Trade) Message() *ResonanceMsg {
	return &ResonanceMsg{
		BtcTxId:  t.BtcTxId,
		Sig:      t.Sig,
		Receiver: t.Receiver,
		Amount:   t.Amount,
	}
}

type TradeWithSig struct {
	BtcHash  common.Hash    `json:"btcHash" gencodec:"required"`
	Receiver common.Address `json:"receiver" gencodec:"required"`
	Amount   *big.Int       `json:"amount" gencodec:"required"`

	hash common.Hash
	sig  []byte
	// Signature values
	V []*big.Int `json:"v" gencodec:"required"`
	R []*big.Int `json:"r" gencodec:"required"`
	S []*big.Int `json:"s" gencodec:"required"`
}

func NewTradeWithSig(trade *Trade) *TradeWithSig {
	t := &TradeWithSig{
		BtcHash:  trade.BtcTxId,
		Receiver: trade.Receiver,
		Amount:   trade.Amount,
		sig:      trade.Sig,
	}
	t.hash = t.Hash()

	t.V = append(t.V, new(big.Int).SetBytes([]byte{trade.Sig[64] + 27}))
	t.R = append(t.R, new(big.Int).SetBytes(trade.Sig[:32]))
	t.S = append(t.S, new(big.Int).SetBytes(trade.Sig[32:64]))

	return t
}

func (t *TradeWithSig) Hash() common.Hash {
	var data []byte
	data = append(data, t.BtcHash.Bytes()...)
	data = append(data, t.Receiver.Bytes()...)
	data = append(data, common.LeftPadBytes(t.Amount.Bytes(), 32)...)

	return crypto.Keccak256Hash(data)
}
