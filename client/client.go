package client

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type NewResonanceTask struct {
	Address   string
	Amount    string
	Signature string
}
type SigResonance struct {
	Address   string
	Amount    string
	Signature string
}

type Resonance struct {
	NewResonanceChan chan *NewResonanceTask
	SigResonanceChan chan *SigResonance
}

type Message struct {
	Address    string
	Amount     string
	Signatures []string
}

func (r *Resonance) NewResonanceTrade(in *Message) {

}

func (r *Resonance) SignResonance() {

}

func StartServe(url string) {
	rpc.Register(new(Resonance))
	rpc.HandleHTTP()

	listen, err := net.Listen("tcp", url)
	if err != nil {
		log.Fatal("fatal error: ", err)
	}
	fmt.Println("start connection")

	http.Serve(listen, nil)
}

type Client struct {
	client *rpc.Client
}

func NewClient(url string) *Client {
	client, err := rpc.DialHTTP("tcp", url)
	if err != nil {
		log.Fatal("dailing error: ", err)
	}

	return &Client{client: client}
}

func (c *Client) NewResonanceTrade(body *Message) error {
	return c.client.Call("Resonance.NewResonanceTrade", body, nil)
}

func (c *Client) SignResonance(body *Message) error {
	return c.client.Call("Resonance.SignResonance", body, nil)
}
