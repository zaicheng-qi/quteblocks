package blockchain

import (
	"fmt"
)

type Transaction struct {
	Sender   string `json:"from" binding:"required"`
	Receiver string `json:"to" binding:"required"`
	Amount   uint32 `json:"amount" binding:"required"`
}

func (t *Transaction) Inspect() {
	fmt.Printf("%v", t)
}
