package main

import (
	"time"

	blockchain "./blockchain"
)

func main() {
	blockChain := blockchain.NewBlockChain(time.Now().String())
	blockChain.PreMining(20)
	blockChain.Inspect()
}
