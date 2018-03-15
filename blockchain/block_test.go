package blockchain

import (
	"reflect"
	"testing"
)

func Test_blockData_hashSum(t *testing.T) {
	type fields struct {
		Index        int
		Timestamp    string
		Data         []Transaction
		PreviousHash string
		Proof        int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Should have correct hash sum",
			fields: fields{
				Index:        1,
				Timestamp:    "2009-11-10 23:00:00 +0000 UTC m=+0.000000001",
				Data:         nil,
				PreviousHash: "",
				Proof:        1,
			},
			want: "2a7859c78a8ea33b299fb71eafd40762ba17ff01560e210f84a15f1175c302d0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BlockData{
				Index:        tt.fields.Index,
				Timestamp:    tt.fields.Timestamp,
				Data:         tt.fields.Data,
				PreviousHash: tt.fields.PreviousHash,
				Proof:        tt.fields.Proof,
			}
			if got := b.hashSum(); got != tt.want {
				t.Errorf("BlockData.hashSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBlock(t *testing.T) {
	type args struct {
		index        int
		data         []Transaction
		timestamp    string
		previousHash string
		proof        int
	}
	tests := []struct {
		name string
		args args
		want *Block
	}{
		{
			name: "Should create a new block",
			args: args{
				index:        1,
				timestamp:    "2009-11-10 23:00:00 +0000 UTC m=+0.000000001",
				data:         nil,
				previousHash: "",
				proof:        1,
			},
			want: &Block{
				BlockData: &BlockData{
					Index:        1,
					Timestamp:    "2009-11-10 23:00:00 +0000 UTC m=+0.000000001",
					Data:         nil,
					PreviousHash: "",
					Proof:        1,
				},
				Hash: "2a7859c78a8ea33b299fb71eafd40762ba17ff01560e210f84a15f1175c302d0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBlock(tt.args.index, tt.args.data, tt.args.timestamp, tt.args.previousHash, tt.args.proof); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}
