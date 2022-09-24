package arrays_and_hashing

import "testing"

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example_1",
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			name: "example_2",
			args: args{
				s: "rat",
				t: "car",
			},
			want: false,
		},
		{ // strings with different lengths
			name: "example_3",
			args: args{
				s: "rat",
				t: "carssss",
			},
			want: false,
		},
		{ // strings with different lengths (s bigger than t)
			name: "example_4",
			args: args{
				s: "carssss",
				t: "rat",
			},
			want: false,
		},
		{ // Follow up: What if the inputs contain Unicode characters?
			name: "example_5",
			args: args{
				s: "#áro!",
				t: "!ro#á",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}

			if got := isAnagram2(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram2() = %v, want %v", got, tt.want)
			}
		})
	}
}
