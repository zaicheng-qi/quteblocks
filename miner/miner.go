package miner

import (
	"github.com/gin-gonic/gin"
	"github.com/vmlinz/quteblocks/blockchain"
)

type Miner struct {
	blockChain   *blockchain.BlockChain
	router       *gin.Engine
	minerAddr    string
	transactions []blockchain.Transaction
}

const MinerAddr = "0000000000000000000000000000000000000000000000000000000000000001"

func NewMiner(blockChain *blockchain.BlockChain) *Miner {
	m := new(Miner)
	m.blockChain = blockChain
	m.router = m.setupRouter()
	m.minerAddr = MinerAddr
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
		blockChainLength := len(m.blockChain.Blocks)
		lastBlock := m.blockChain.Blocks[blockChainLength-1]
		lastProof := lastBlock.blockData
	})

	return router
}

func (m *Miner) Run() {
	m.router.Run()
}
