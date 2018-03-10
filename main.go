package main

import (
	"time"

	"github.com/vmlinz/quteblocks/blockchain"
	"github.com/vmlinz/quteblocks/miner"
)

func main() {
	blockChain := blockchain.NewBlockChain(time.Now().String())
	blockChain.Inspect()

	miner := miner.NewMiner(blockChain)
	miner.Run()
}
