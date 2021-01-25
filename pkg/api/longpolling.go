package api

import (
	"github.com/davecgh/go-spew/spew"
	jsonrpc "github.com/ybbus/jsonrpc/v2"
)

const uriLongPolling = ":8444/remote/json-rpc"

func (b *boschShcAPI) Subscribe() (e error) {
	url := "https://" + b.shcIPAddress + uriLongPolling

	rpcClient := jsonrpc.NewClientWithOpts(url, &jsonrpc.RPCClientOpts{
		HTTPClient: b.client,
	})

	err := rpcClient.CallFor(&b.pollingID, "RE/subscribe", "com/bosch/sh/remote/*", nil)
	if err != nil {
		b.pollingID = ""
		return err
	}
	return nil
}

func (b *boschShcAPI) Unsubscribe() (e error) {
	url := "https://" + b.shcIPAddress + uriLongPolling

	rpcClient := jsonrpc.NewClientWithOpts(url, &jsonrpc.RPCClientOpts{
		HTTPClient: b.client,
	})

	r, err := rpcClient.Call("RE/unsubscribe", b.pollingID)
	if err != nil {
		return err
	}
	if r.Error != nil {
		return r.Error
	}
	b.pollingID = ""

	return nil
}

const timeoutLongpolling = 30

type rpcData interface{}

func (b *boschShcAPI) Poll() (e error) {
	url := "https://" + b.shcIPAddress + uriLongPolling

	rpcClient := jsonrpc.NewClientWithOpts(url, &jsonrpc.RPCClientOpts{
		HTTPClient: b.client,
	})

	r, err := rpcClient.Call("RE/longPoll", b.pollingID, timeoutLongpolling)
	if err != nil {
		return err
	}
	if r.Error != nil {
		return r.Error
	}
	var d rpcData
	r.GetObject(&d)
	spew.Dump(d)
	return nil
}
