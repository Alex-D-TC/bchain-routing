package swiss

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/alex-d-tc/bchain-routing/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	node   *SwissNode
	logger *log.Logger
}

func MakeClient(node *SwissNode) *Client {
	client := &Client{
		node:   node,
		logger: log.New(os.Stdout, "Swiss client: ", log.Ldate|log.Ltime),
	}

	return client
}

func (client *Client) watchForPaymentRequests(ctx context.Context) {

	userAddr := crypto.PubkeyToAddress(client.node.PrivateKey.PublicKey)

	go eth.EventWatcher(ctx, client.node.client, func(opts *bind.FilterOpts) {

		safeEthclient := client.node.client

		iterator, err := client.node.relay.Relay.FilterRelayPaymentRequested(opts, []common.Address{userAddr})
		if err != nil {
			client.debug(err)
			return
		}

		for iterator.Next() {
			relayEvent := iterator.Event
			request, err := client.node.relay.Relay.GetRelay(nil, crypto.PubkeyToAddress(client.node.PrivateKey.PublicKey), relayEvent.Relay)
			if err != nil {
				client.debug(err)
				continue
			}

			safeEthclient.SubmitTransaction(func(ethclient *ethclient.Client) error {

				auth, err := eth.PrepareTransactionAuth(ethclient, client.node.PrivateKey)
				if err != nil {
					client.debug(err)
					return err
				}

				tran, err := client.node.coin.Coin.Allow(auth, client.node.relay.RelayAddr, request.SentBytes)
				if err != nil {
					client.debug(err)
				} else {
					client.debug(tran.Hash().Hex())
				}

				return err
			})

		}
	})
}

func (client *Client) WatchForAllowedConfirmation(ctx context.Context) {

	userAddr := crypto.PubkeyToAddress(client.node.PrivateKey.PublicKey)

	go eth.EventWatcher(ctx, client.node.client, func(opts *bind.FilterOpts) {

		safeEthclient := client.node.client

		iterator, err := client.node.coin.Coin.FilterAllowed(opts, []common.Address{userAddr}, []common.Address{client.node.relay.RelayAddr})
		if err != nil {
			client.debug(err)
			return
		}

		for iterator.Next() {
			evnt := iterator.Event
			safeEthclient.SubmitTransaction(func(ethclient *ethclient.Client) error {
				auth, err := eth.PrepareTransactionAuth(ethclient, client.node.PrivateKey)
				if err != nil {
					client.debug(err)
					return err
				}

				tran, err := client.node.relay.Relay.HonorRelay(auth, userAddr, evnt.Value)
				if err != nil {
					client.debug(err)
				} else {
					client.debug(tran.Hash().Hex())
				}

				return err
			})
		}

	})
}

func (client *Client) SetOutput(writer io.Writer) {
	client.logger = log.New(writer, client.logger.Prefix(), client.logger.Flags())
}

func (client *Client) debug(msg ...interface{}) {
	client.logger.Println(msg...)
}
