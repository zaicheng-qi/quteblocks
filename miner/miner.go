package miner

import (
	"github.com/gin-gonic/gin"
	"github.com/vmlinz/quteblocks/blockchain"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	transactions := make([]blockchain.Transaction, 10)

	router.POST("transaction", func(c *gin.Context) {
		var tx blockchain.Transaction
		if c.Bind(&tx) == nil {
			(&tx).Inspect()
			transactions = append(transactions, tx)
			c.JSON(200, gin.H{"status": "ok"})
		}
	})

	return router
}

func Run() {
	router := setupRouter()
	router.Run()
}
