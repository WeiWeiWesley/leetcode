package palindrome

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "example 1",
			args: args{121},
			want: true,
		},
		{
			name: "example 2",
			args: args{123},
			want: false,
		},
		{
			name: "example 3",
			args: args{1234},
			want: false,
		},
		{
			name: "example 4",
			args: args{1221},
			want: true,
		},
		{
			name: "example 5",
			args: args{1},
			want: true,
		},
		{
			name: "example 6",
			args: args{-121},
			want: false,
		},
		{
			name: "example 7",
			args: args{-1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
