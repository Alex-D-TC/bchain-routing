package swiss

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/big"
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

	ctx        context.Context
	cancelFunc context.CancelFunc

	started bool
}

func MakeClient(node *SwissNode) *Client {
	ctx, cancelFunc := context.WithCancel(context.Background())

	client := &Client{
		node:       node,
		logger:     log.New(os.Stdout, "Swiss client: ", log.Ldate|log.Ltime),
		ctx:        ctx,
		cancelFunc: cancelFunc,
		started:    false,
	}

	return client
}

func (client *Client) Start(processor func(*Message)) {
	if !client.started {
		client.started = true

		go client.watchForAllowedConfirmation(client.ctx)
		go client.watchForPaymentRequests(client.ctx)

		client.node.Start(processor)
	}
}

func (client *Client) JoinAndStart(processor func(*Message), bootstrapIP string, bootstrapPort int) error {
	if !client.started {
		client.started = true

		go client.watchForAllowedConfirmation(client.ctx)
		go client.watchForPaymentRequests(client.ctx)

		err := client.node.JoinAndStart(processor, bootstrapIP, bootstrapPort)
		if err != nil {
			return err
		}
	}

	return nil
}

func (client *Client) Terminate() {
	if client.started {
		client.started = false

		client.node.Terminate()
		client.cancelFunc()
	}
}

func (client *Client) watchForPaymentRequests(ctx context.Context) {

	client.debug("Starting relay request watcher")

	nodeID := client.node.ID

	eth.EventWatcher(ctx, client.node.client, func(opts *bind.FilterOpts) {

		safeEthclient := client.node.client

		iterator, err := client.node.relay.Relay.FilterRelayPaymentRequested(opts, []*big.Int{nodeID.Base10()})
		if err != nil {
			client.debug(err)
			return
		}

		for iterator.Next() {

			client.debug("Honoring request")

			relayEvent := iterator.Event
			request, err := client.node.relay.Relay.GetRelay(nil, nodeID.Base10(), relayEvent.Relay)
			if err != nil {
				client.debug(err)
				continue
			}

			client.debug("Submitting transaction")

			safeEthclient.SubmitTransaction(func(ethclient *ethclient.Client) (error, bool) {

				auth, err := eth.PrepareTransactionAuth(ethclient, client.node.PrivateKey)
				if err != nil {
					client.debug(err)
					return err, false
				}

				tran, err := client.node.coin.Coin.Allow(auth, client.node.relay.RelayAddr, request.SentBytes)
				if err != nil {
					client.debug(err)
				} else {
					client.debug(tran.Hash().Hex())
				}

				return err, false
			})

		}
	})
}

func (client *Client) watchForAllowedConfirmation(ctx context.Context) {

	client.debug("Starting allowance watcher")

	userAddr := crypto.PubkeyToAddress(client.node.PrivateKey.PublicKey)
	nodeID := client.node.ID

	eth.EventWatcher(ctx, client.node.client, func(opts *bind.FilterOpts) {

		safeEthclient := client.node.client

		iterator, err := client.node.coin.Coin.FilterAllowed(opts, []common.Address{userAddr}, []common.Address{client.node.relay.RelayAddr})
		if err != nil {
			client.debug(err)
			return
		}

		for iterator.Next() {

			fmt.Println("Found allowance. Honoring")

			evnt := iterator.Event
			safeEthclient.SubmitTransaction(func(ethclient *ethclient.Client) (error, bool) {
				auth, err := eth.PrepareTransactionAuth(ethclient, client.node.PrivateKey)
				if err != nil {
					client.debug(err)
					return err, false
				}

				tran, err := client.node.relay.Relay.HonorRelay(auth, nodeID.Base10(), userAddr, evnt.Value)
				if err != nil {
					client.debug(err)
				} else {
					client.debug(tran.Hash().Hex())
				}

				return err, false
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
