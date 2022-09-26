package arrays_and_hashing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "example_1",
			args: args{
				strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name: "example_2",
			args: args{
				strs: []string{""},
			},
			want: [][]string{{""}},
		},
		{
			name: "example_3",
			args: args{
				strs: []string{"a"},
			},
			want: [][]string{{"a"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !checkArraysNoOrder(t, got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

// it works to assert that arrays with the same length have the same elements, even if in different order
func checkArraysNoOrder(t *testing.T, a, b [][]string) bool {
	for _, aVal := range a {
		for _, bVal := range b {
			if len(aVal) == len(bVal) {
				if !assert.ElementsMatch(t, aVal, bVal) {
					return false
				}
			}
		}
	}
	return true
}
