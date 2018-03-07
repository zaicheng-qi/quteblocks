package blockchain

import (
	"fmt"
	"time"
)

type BlockChain struct {
	Blocks []Block
}

func createGenesisBlock(timestamp string) *Block {
	return NewBlock(0, nil, timestamp, "")
}

func NewBlockChain(timestamp string) *BlockChain {
	blockChain := new(BlockChain)
	blockChain.Blocks = append(blockChain.Blocks, *createGenesisBlock(timestamp))

	return blockChain
}

func (blockChain *BlockChain) AppendNewBlock(data []byte, timestamp string) *BlockChain {
	blockLength := len(blockChain.Blocks)

	if blockLength < 1 {
		return nil
	}

	currentBlock := blockChain.Blocks[blockLength-1]
	newBlock := NewBlock(blockLength, data, timestamp, currentBlock.hash)

	blockChain.Blocks = append(blockChain.Blocks, *newBlock)

	return blockChain
}

func (blockChain *BlockChain) PreMining(number int) *BlockChain {
	for i := 1; i < number; i++ {
		blockChain = blockChain.AppendNewBlock(nil, time.Now().String())
	}

	return blockChain
}

func (blockChain *BlockChain) Inspect() {
	blocks := blockChain.Blocks
	for _, block := range blocks {
		fmt.Printf("{%v,\n %v}\n", block.blockData.PreviousHash, block.hash)
	}
}
