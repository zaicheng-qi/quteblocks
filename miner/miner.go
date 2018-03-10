package miner

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vmlinz/quteblocks/blockchain"
)

const MINER_ADDR = "0000000000000000000000000000000000000000000000000000000000000001"
const NETWORK_ADDR = "0000000000000000000000000000000000000000000000000000000000000000"

type Miner struct {
	blockChain   *blockchain.BlockChain
	router       *gin.Engine
	minerAddr    string
	transactions []blockchain.Transaction
}

func NewMiner(blockChain *blockchain.BlockChain) *Miner {
	m := new(Miner)
	m.blockChain = blockChain
	m.router = m.setupRouter()
	m.minerAddr = MINER_ADDR
	m.transactions = []blockchain.Transaction{}

	return m
}

func proofOfWork(lastProof int) int {
	incrementor := lastProof + 1
	for !(incrementor%9 == 0 && incrementor%lastProof == 0) {
		incrementor++
	}

	return incrementor
}

func (m *Miner) setupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("transaction", func(c *gin.Context) {
		var tx blockchain.Transaction
		if c.Bind(&tx) == nil {
			tx.Inspect()
			m.transactions = append(m.transactions, tx)
			c.JSON(200, gin.H{"status": "ok"})
		}
	})

	router.GET("mine", func(c *gin.Context) {
		// get last proof
		blockChainLength := len(m.blockChain.Blocks)
		lastBlock := m.blockChain.Blocks[blockChainLength-1]
		lastProof := lastBlock.BlockData.Proof

		// doing POW
		proof := proofOfWork(lastProof)
		// receive rewarded coins by adding a new transaction from network to miner
		m.transactions = append(m.transactions, blockchain.Transaction{Sender: NETWORK_ADDR, Receiver: m.minerAddr, Amount: 1})

		// append new mined block to blockchain
		m.blockChain.AppendNewBlock(m.transactions, time.Now().String(), proof)
		m.transactions = []blockchain.Transaction{}
	})

	return router
}

func (m *Miner) Run() {
	m.router.Run()
}
