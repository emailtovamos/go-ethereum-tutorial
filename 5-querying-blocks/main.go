package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	// header
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(header.Number.String()) // 15495248

	blockNumber := big.NewInt(15495248)

	// full block
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 15495248
	fmt.Println(block.Time())                // 1662621329
	fmt.Println(block.Difficulty().Uint64()) // 12774904243062970
	fmt.Println(block.Hash().Hex())          // 0xacd7dd26f51fb92079d1486195b8d10dd3ea83b047f96922877fdb6226828260
	fmt.Println(len(block.Transactions()))   // 217

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // 217
}
