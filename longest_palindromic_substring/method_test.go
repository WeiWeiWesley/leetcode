package longestpalindromicsubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				"babad",
			},
			want: []string{"bab", "aba"},
		},
		{
			name: "case 2",
			args: args{
				"cbbd",
			},
			want: []string{"bb"},
		},
		{
			name: "case 3",
			args: args{
				"ab",
			},
			want: []string{"a", "b"},
		},
		{
			name: "case 4",
			args: args{
				"c",
			},
			want: []string{"c"},
		},
		{
			name: "case 5",
			args: args{
				"",
			},
			want: []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Contains(t, tt.want, longestPalindrome(tt.args.s))
		})
	}
}
