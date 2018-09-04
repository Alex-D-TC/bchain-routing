package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/alex-d-tc/bchain-routing/concurrent"

	ethBind "github.com/alex-d-tc/bchain-routing/eth/build-go"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CoinContract struct {
	Coin     *ethBind.SwissCoin
	CoinAddr common.Address
}

func MakeCoinContract(coin *ethBind.SwissCoin, coinAddr common.Address) CoinContract {
	return CoinContract{
		Coin:     coin,
		CoinAddr: coinAddr,
	}
}

type RelayContract struct {
	Relay     *ethBind.RelayHandler
	RelayAddr common.Address
}

func MakeRelayContract(relay *ethBind.RelayHandler, relayAddr common.Address) RelayContract {
	return RelayContract{
		Relay:     relay,
		RelayAddr: relayAddr,
	}
}

type ThreadsafeClient struct {
	sync.RWMutex

	client *ethclient.Client
	queue  *concurrent.TransactionQueue
}

func MakeThreadsafeClient(client *ethclient.Client) *ThreadsafeClient {
	result := ThreadsafeClient{
		client: client,
		queue:  concurrent.MakeTransactionQueue(),
	}

	return &result
}

func (client *ThreadsafeClient) SubmitTransaction(tran func(*ethclient.Client) (error, bool)) error {
	return client.queue.Submit(func() (error, bool) {
		client.Lock()

		err, cont := tran(client.client)

		client.Unlock()
		return err, cont
	})
}

func (client *ThreadsafeClient) Dispose() {
	client.queue.Dispose()
}

func GetThreadsafeClient(rawUrl string) (*ThreadsafeClient, error) {
	client, err := GetClient(rawUrl)
	if err != nil {
		return nil, err
	}

	return MakeThreadsafeClient(client), nil
}

func GetClient(rawUrl string) (*ethclient.Client, error) {
	return ethclient.Dial(rawUrl)
}

func GetSwissCoin(coinAddr common.Address, client *ThreadsafeClient) (CoinContract, error) {

	client.RLock()
	coin, err := ethBind.NewSwissCoin(coinAddr, client.client)
	client.RUnlock()

	return CoinContract{Coin: coin, CoinAddr: coinAddr}, err
}

func GetRelayHandler(relayAddr common.Address, client *ThreadsafeClient) (RelayContract, error) {

	client.RLock()
	relay, err := ethBind.NewRelayHandler(relayAddr, client.client)
	client.RUnlock()

	return RelayContract{Relay: relay, RelayAddr: relayAddr}, err
}

func PrepareTransactionAuth(client *ethclient.Client, key *ecdsa.PrivateKey) (*bind.TransactOpts, error) {

	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(key.PublicKey))
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(key)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasPrice = gasPrice
	//auth.GasLimit = 500000

	return auth, nil
}

func EventWatcher(ctx context.Context, logger *log.Logger, client *ThreadsafeClient, filterProcessor func(*bind.FilterOpts)) {

	done := false
	lastBlock := uint64(0)

	for {

		select {
		case <-ctx.Done():
			done = true
			break
		default:
			break
		}

		if done {
			break
		}

		time.Sleep(5 * time.Second)

		client.RLock()

		header, err := client.client.BlockByNumber(ctx, nil)
		if err != nil {
			fmt.Println(err)
			// Retry on error
			continue
		}

		client.RUnlock()

		block := uint64(header.Number().Int64())

		if lastBlock == block {
			//fmt.Println("Block already processed. Sleeping")
			continue
		}

		//logger.Println(fmt.Sprintf("Checking events from block %d to block %d", lastBlock+1, block))

		opts := &bind.FilterOpts{
			Context: ctx,
			Start:   lastBlock + 1,
			End:     &block,
		}

		filterProcessor(opts)
		lastBlock = block
	}
}
