package sorting

import (
	"reflect"
	"testing"
)

func Test_heapSort(t *testing.T) {
	type args struct {
		ary []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				ary: []int{6, 8, 2, 5, 1},
			},
			want: []int{1, 2, 5, 6, 8},
		},
		{
			name: "",
			args: args{
				ary: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := heapSort(tt.args.ary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("heapSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
