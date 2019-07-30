package anchors

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/rpc"

	"github.com/walker1992/btc_resonance/log"

	"github.com/ethereum/go-ethereum/common"
)

type Resonance struct {
	//NewResonanceChan chan *NewResonanceTask
	//SigResonanceChan chan *SigResonance

	Trades  map[common.Hash]*TradeWithSig
	Anchors map[common.Address]bool

	ReceiveTradeChan chan *Trade
	SyncTradeChan    chan *Trade
}

type TradeServer struct {
	ReceiveTradeChan chan *Trade
}

func (r *Resonance) VerifySig(trade *Trade) bool {
	signer, err := trade.getSigner()
	if err != nil {
		return false
	}
	_, ok := r.Anchors[signer]
	if !ok {
		return false
	}
	return true
}

func (r *TradeServer) ResonanceTrade(msg *ResonanceMsg, out *string) error {
	log.Info("receive a resonanceTrade: %v", msg)

	fmt.Println(msg.BtcTxId.String(), msg.Receiver.String(), msg.Amount.String(), msg.Sig)

	trade := NewTradeFromMsg(msg)

	r.ReceiveTradeChan <- trade

	return nil
}

func (r *Resonance) ResonanceTrade(msg *ResonanceMsg, out *string) error {
	log.Info("receive a resonanceTrade: %v", msg)

	fmt.Println(msg.BtcTxId.String(), msg.Receiver.String(), msg.Amount.String(), msg.Sig)

	trade := NewTradeFromMsg(msg)

	r.ReceiveTradeChan <- trade

	return nil
}

func (r *Resonance) TradeLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Info("RelayNonce closed")
			return
		case receiveTrade := <-r.ReceiveTradeChan:
			if r.VerifySig(receiveTrade) {
				r.collectTrade(receiveTrade)
			}

		case syncTrade := <-r.SyncTradeChan:
			r.collectTrade(syncTrade)
		}
	}

}
func (r *Resonance) collectTrade(trade *Trade) {
	tradeWithSig, ok := r.Trades[trade.BtcTxId]
	if !ok {
		r.Trades[trade.BtcTxId] = NewTradeWithSig(trade)
		return
	} else {
		tradeWithSig.V = append(tradeWithSig.V, trade.V)
		tradeWithSig.R = append(tradeWithSig.R, trade.R)
		tradeWithSig.S = append(tradeWithSig.S, trade.S)
	}

	fmt.Println("the signature count is ", len(tradeWithSig.V))
	if len(tradeWithSig.V) >= 3 {
		fmt.Println("send eth resonance.sol ResonanceTrade()")

		delete(r.Trades, trade.BtcTxId)
	}

}

func StartServe(url string, resonance *TradeServer) error {
	err := rpc.Register(resonance)
	if err != nil {
		return err
	}

	rpc.HandleHTTP()

	listen, err := net.Listen("tcp", url)
	if err != nil {
		return err
	}
	log.Info("start server: %s", url)

	go http.Serve(listen, nil)
	return nil
}
