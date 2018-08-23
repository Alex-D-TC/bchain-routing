package swiss

import (
	"context"
	"fmt"

	"github.com/alex-d-tc/bchain-routing/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	node *SwissNode
}

func MakeClient(node *SwissNode) *Client {
	client := &Client{
		node: node,
	}

	return client
}

func (client *Client) watchForPaymentRequests(ctx context.Context) {

	userAddr := crypto.PubkeyToAddress(client.node.PrivateKey.PublicKey)

	go eth.EventWatcher(ctx, client.node.client, func(opts *bind.FilterOpts) {

		safeEthclient := client.node.client

		iterator, err := client.node.relay.FilterRelayPaymentRequested(opts, []common.Address{userAddr})
		if err != nil {
			fmt.Println(err)
			return
		}

		for iterator.Next() {
			relayEvent := iterator.Event
			request, err := client.node.relay.GetRelay(nil, crypto.PubkeyToAddress(client.node.PrivateKey.PublicKey), relayEvent.Relay)
			if err != nil {
				fmt.Println(err)
				continue
			}

			safeEthclient.SubmitTransaction(func(ethclient *ethclient.Client) error {
				auth, err := eth.PrepareTransactionAuth(ethclient, client.node.PrivateKey)
				if err != nil {
					fmt.Println(err)
					return err
				}

				tran, err := client.node.coin.Allow(auth, client.node.relayAddr, request.SentBytes)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(tran.Hash().Hex())
				}

				return err
			})

		}
	})
}

func (client *Client) WatchForTransferConfirmation(ctx context.Context) {

	userAddr := crypto.PubkeyToAddress(client.node.PrivateKey.PublicKey)

	go eth.EventWatcher(ctx, client.node.client, func(opts *bind.FilterOpts) {

		safeEthclient := client.node.client

		iterator, err := client.node.coin.FilterTransfer(opts, []common.Address{userAddr}, []common.Address{client.node.relayAddr})
		if err != nil {
			fmt.Println(err)
			return
		}

		for iterator.Next() {
			evnt := iterator.Event
			safeEthclient.SubmitTransaction(func(ethclient *ethclient.Client) error {
				auth, err := eth.PrepareTransactionAuth(ethclient, client.node.PrivateKey)
				if err != nil {
					fmt.Println(err)
					return err
				}

				tran, err := client.node.relay.HonorRelay(auth, userAddr, evnt.Value)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(tran.Hash().Hex())
				}

				return err
			})
		}

	})
}
