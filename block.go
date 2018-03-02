package blockchain

import "crypto/sha256"
import "encoding/json"
import "encoding/hex"

type blockData struct {
	Index        int    `json:"index"`
	Timestamp    string `json:"timestamp"`
	Data         []byte `json:"data"`
	PreviousHash string `json:"previous_hash"`
}

type Block struct {
	blockData *blockData
	hash      string
}

func (b blockData) hashSum() string {
	blockData, err := json.Marshal(b)
	hashSum := sha256.New().Sum(blockData)

	if err != nil {
		return ""
	}

	return hex.EncodeToString(hashSum)
}

func NewBlock(index int, data []byte, timestamp string, previousHash string) *Block {
	blockData := new(blockData)
	blockData.Index = index
	blockData.Data = data
	blockData.Timestamp = timestamp
	blockData.PreviousHash = previousHash

	block := new(Block)
	block.blockData = blockData
	block.hash = blockData.hashSum()

	return block
}
