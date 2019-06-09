package sorting

import (
	"reflect"
	"testing"
)

func Test_radixSort(t *testing.T) {
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
				ary: []int{170, 45, 75, 90, 2, 24, 802, 66},
			},
			want: []int{2, 24, 45, 66, 75, 90, 170, 802},
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
			if got := radixSort(tt.args.ary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("radixSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
