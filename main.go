package main

import (
	"time"

	"./blockchain"
)

func main() {
	blockChain := blockchain.NewBlockChain(time.Now().String())
	blockChain.PreMining(10)
	blockChain.Inspect()
}
