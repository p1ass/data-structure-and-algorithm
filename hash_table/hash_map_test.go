package hash_table

import (
	"testing"
)

func Test_hashTable_Insert(t *testing.T) {
	type fields struct {
		N int
	}
	type args struct {
		val []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "",
			fields: fields{
				N: 10,
			},
			args: args{
				val: []int{1, 4, 7, 11, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashTable(tt.fields.N)
			for _, val := range tt.args.val {
				h.Insert(val, val)
			}

		})
	}
}

func Test_hashMap_Get(t *testing.T) {
	type fields struct {
		N    int
		data [][]int
	}
	type args struct {
		key int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
		want1  bool
	}{
		{
			name: "",
			fields: fields{
				N:    10,
				data: [][]int{{1, 1}, {4, 4}, {7, 7}, {11, 11}, {4, 5}},
			},
			args: args{
				key: 1,
			},
			want:  1,
			want1: true,
		},
		{
			name: "",
			fields: fields{
				N:    10,
				data: [][]int{{1, 1}, {4, 4}, {7, 7}, {11, 11}, {4, 5}},
			},
			args: args{
				key: 4,
			},
			want:  5,
			want1: true,
		},
		{
			name: "",
			fields: fields{
				N:    10,
				data: [][]int{{1, 1}, {4, 4}, {7, 7}, {11, 11}, {4, 5}},
			},
			args: args{
				key: 8,
			},
			want:  0,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHashTable(tt.fields.N)
			for _, val := range tt.fields.data {
				h.Insert(val[0], val[1])
			}
			got, got1 := h.Get(tt.args.key)
			if got != tt.want {
				t.Errorf("hashMap.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("hashMap.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
