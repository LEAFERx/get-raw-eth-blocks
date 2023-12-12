package main

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	client, err := ethclient.Dial("https://rpc.ankr.com/eth")
	if err != nil {
		log.Fatal(err)
	}
	blockNumber := uint64(100000)
	block, err := getBlock(context.Background(), client, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	serializedBlock, err := serializeBlock(block)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Block number: ", block.Number().Uint64())
	log.Info("Block serialized: ", hexutil.Encode(serializedBlock))
}

func getBlock(ctx context.Context, client *ethclient.Client, number uint64) (*types.Block, error) {
	return client.BlockByNumber(ctx, big.NewInt(int64(number)))
}

func serializeBlock(block *types.Block) ([]byte, error) {
	return rlp.EncodeToBytes(block)
}
