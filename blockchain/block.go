package blockchain

import "crypto/sha256"
import "encoding/json"
import "fmt"

type BlockData struct {
	Index        int           `json:"index"`
	Timestamp    string        `json:"timestamp"`
	Data         []Transaction `json:"data"`
	Proof        int           `json:"proof_of_work"`
	PreviousHash string        `json:"previous_hash"`
}

type Block struct {
	BlockData *BlockData
	hash      string
}

func (b BlockData) hashSum() string {
	blockData, err := json.Marshal(b)
	hashSum := sha256.Sum256(blockData)

	if err != nil {
		return ""
	}

	return fmt.Sprintf("%x", hashSum)
}

func NewBlock(index int, data []Transaction, timestamp string, previousHash string, proof int) *Block {
	blockData := new(BlockData)
	blockData.Index = index
	blockData.Data = data
	blockData.Timestamp = timestamp
	blockData.PreviousHash = previousHash
	blockData.Proof = proof

	block := new(Block)
	block.BlockData = blockData
	block.hash = blockData.hashSum()

	return block
}
