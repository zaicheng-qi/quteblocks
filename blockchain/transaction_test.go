package blockchain

import (
	"reflect"
	"testing"
)

func TestTransaction_Marshal(t *testing.T) {
	type fields struct {
		Sender   string
		Receiver string
		Amount   float32
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
		{
			name: "Should marshal to a json byte array",
			fields: fields{
				Sender:   "0",
				Receiver: "1",
				Amount:   0.1,
			},
			want: []byte(`{"from":"0","to":"1","amount":0.1}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := &Transaction{
				Sender:   tt.fields.Sender,
				Receiver: tt.fields.Receiver,
				Amount:   tt.fields.Amount,
			}
			if got := tx.Marshal(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transaction.Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}
