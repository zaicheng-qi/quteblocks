package block

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
			want: "7b22696e646578223a312c2274696d657374616d70223a22323030392d31312d31302032333a30303a3030202b3030303020555443206d3d2b302e303030303030303031222c2264617461223a6e756c6c2c2270726576696f75735f68617368223a22227de3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
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
				hash: "7b22696e646578223a312c2274696d657374616d70223a22323030392d31312d31302032333a30303a3030202b3030303020555443206d3d2b302e303030303030303031222c2264617461223a6e756c6c2c2270726576696f75735f68617368223a22227de3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
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
