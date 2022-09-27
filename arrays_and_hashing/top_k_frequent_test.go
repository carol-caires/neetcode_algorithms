package arrays_and_hashing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example_1",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			want: []int{1, 2},
		},
		{
			name: "example_2",
			args: args{
				nums: []int{1},
				k:    1,
			},
			want: []int{1},
		},
		{
			name: "example_3",
			args: args{
				nums: []int{4, 3, 4, 1, 1, 4, 1, 2, 4, 3, 3, 2, 3, 4},
				k:    2,
			},
			want: []int{4, 3},
		},
	}
	for _, tt := range tests {
		if tt.name != "example_2" {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, topKFrequent(tt.args.nums, tt.args.k), "topKFrequent(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
