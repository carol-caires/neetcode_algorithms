package arrays_and_hashing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example_1",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "example_2",
			args: args{
				nums: []int{-1, 1, 0, -3, 3},
			},
			want: []int{0, 0, 9, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, productExceptSelf(tt.args.nums), "productExceptSelf(%v)", tt.args.nums)
			assert.Equalf(t, tt.want, productExceptSelf2(tt.args.nums), "productExceptSelf2(%v)", tt.args.nums)
		})
	}
}
