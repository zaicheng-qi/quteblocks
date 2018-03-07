package blockchain

import (
	"reflect"
	"testing"
)

func Test_blockData_hashSum(t *testing.T) {
	type fields struct {
		Index        int
		Timestamp    string
		Data         []byte
		PreviousHash string
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
			},
			want: "8d726f5b6e435b2d400c80a36269b4ab42694da98a468fa7bc214f958c4c5e53",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &blockData{
				Index:        tt.fields.Index,
				Timestamp:    tt.fields.Timestamp,
				Data:         tt.fields.Data,
				PreviousHash: tt.fields.PreviousHash,
			}
			if got := b.hashSum(); got != tt.want {
				t.Errorf("blockData.hashSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBlock(t *testing.T) {
	type args struct {
		index        int
		data         []byte
		timestamp    string
		previousHash string
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
			},
			want: &Block{
				blockData: &blockData{
					Index:        1,
					Timestamp:    "2009-11-10 23:00:00 +0000 UTC m=+0.000000001",
					Data:         nil,
					PreviousHash: "",
				},
				hash: "8d726f5b6e435b2d400c80a36269b4ab42694da98a468fa7bc214f958c4c5e53",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBlock(tt.args.index, tt.args.data, tt.args.timestamp, tt.args.previousHash); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}
