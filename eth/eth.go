package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetClient(rawUrl string) (*ethclient.Client, error) {
	return ethclient.Dial(rawUrl)
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

	return auth, nil
}

func EventWatcher(ctx context.Context, client *ethclient.Client, filterProcessor func(*bind.FilterOpts)) {

	done := false
	var lastEnd *big.Int

	for {

		time.Sleep(5 * time.Second)

		header, err := client.BlockByNumber(ctx, nil)
		if err != nil {
			panic(err)
		}

		start := header.Number().Sub(header.Number(), big.NewInt(1))
		end := uint64(header.Number().Int64())

		if lastEnd != nil && lastEnd.Cmp(header.Number()) == 0 {
			fmt.Println("Block already processed. Sleeping")
			continue
		}

		lastEnd = header.Number()

		fmt.Println("Checking events in block: ", end)

		opts := &bind.FilterOpts{
			Context: ctx,
			Start:   uint64(start.Int64()),
			End:     &end,
		}

		filterProcessor(opts)

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
	}

}
