package blockchain

import (
	"reflect"
	"testing"
)

func Test_createGenesisBlock(t *testing.T) {
	type args struct {
		timestamp string
	}
	tests := []struct {
		name string
		args args
		want *Block
	}{
		// TODO: Add test cases.
		{
			name: "Should create genesis block",
			args: args{timestamp: "2009-11-10 23:00:00 +0000 UTC m=+0.000000001"},
			want: NewBlock(0, nil, "2009-11-10 23:00:00 +0000 UTC m=+0.000000001", ""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createGenesisBlock(tt.args.timestamp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createGenesisBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBlockChain(t *testing.T) {
	type args struct {
		timestamp string
	}
	tests := []struct {
		name string
		args args
		want *BlockChain
	}{
		// TODO: Add test cases.
		{
			name: "Should create a new blockchain",
			args: args{
				timestamp: "2009-11-10 23:00:00 +0000 UTC m=+0.000000001",
			},
			want: &BlockChain{
				Blocks: []Block{*createGenesisBlock("2009-11-10 23:00:00 +0000 UTC m=+0.000000001")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBlockChain(tt.args.timestamp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBlockChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlockChain_AppendNewBlock(t *testing.T) {
	type fields struct {
		Blocks []Block
	}
	type args struct {
		data      []byte
		timestamp string
	}
	genesisBlock := *createGenesisBlock("2009-11-10 23:00:00 +0000 UTC m=+0.000000001")
	timestamp := "2009-11-10 23:10:00 +0000 UTC m=+0.000000001"
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BlockChain
	}{
		// TODO: Add test cases.
		{
			name:   "Should append a new block",
			fields: fields{Blocks: []Block{genesisBlock}},
			args: args{
				data:      nil,
				timestamp: timestamp,
			},
			want: &BlockChain{
				Blocks: []Block{genesisBlock, *NewBlock(1, nil, timestamp, genesisBlock.hash)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			blockChain := &BlockChain{
				Blocks: tt.fields.Blocks,
			}
			if got := blockChain.AppendNewBlock(tt.args.data, tt.args.timestamp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BlockChain.AppendNewBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlockChain_PreMining(t *testing.T) {
	type fields struct {
		Blocks []Block
	}
	type args struct {
		number int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BlockChain
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			blockChain := &BlockChain{
				Blocks: tt.fields.Blocks,
			}
			if got := blockChain.PreMining(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BlockChain.PreMining() = %v, want %v", got, tt.want)
			}
		})
	}
}
