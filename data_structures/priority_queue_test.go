package data_structures

import "testing"

func Test_priorityQueue_Pop(t *testing.T) {
	type fields struct {
		data []int
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				data: []int{6, 8, 2, 5, 1},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "",
			fields: fields{
				data: []int{1, 6, 8, 2, 5},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := &priorityQueue{}

			for _, val := range tt.fields.data {
				pq.Push(val)
			}

			got, err := pq.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("priorityQueue.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("priorityQueue.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
