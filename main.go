package main

import (
	"time"

	"github.com/vmlinz/quteblocks/blockchain"
	"github.com/vmlinz/quteblocks/miner"
)

func main() {
	blockChain := blockchain.NewBlockChain(time.Now().String())
	blockChain.PreMining(10)
	blockChain.Inspect()

	miner.Run()
}
