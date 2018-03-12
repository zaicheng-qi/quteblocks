package blockchain

import (
	"encoding/json"
	"fmt"
)

type Transaction struct {
	Sender   string  `json:"from" binding:"required"`
	Receiver string  `json:"to" binding:"required"`
	Amount   float32 `json:"amount" binding:"required"`
}

func (t *Transaction) Inspect() {
	fmt.Printf("%s", string(t.Marshal()))
}

func (t *Transaction) Marshal() []byte {
	txJSON, err := json.Marshal(*t)

	if err == nil {
		return txJSON
	}

	return []byte("{}")
}
