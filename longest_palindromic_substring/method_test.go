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
		{
			name: "case 6 - 偶數長度回文",
			args: args{
				"abccba",
			},
			want: []string{"abccba"},
		},
		{
			name: "case 7 - 長回文",
			args: args{
				"raceacar",
			},
			want: []string{"aca"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Contains(t, tt.want, longestPalindrome(tt.args.s))
		})
	}
}

func Test_longestPalindromeBruteForce(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Contains(t, tt.want, longestPalindromeBruteForce(tt.args.s))
		})
	}
}

// 性能比較測試
func BenchmarkLongestPalindrome(b *testing.B) {
	testString := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"

	b.Run("ExpandAroundCenter", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			longestPalindrome(testString)
		}
	})

	b.Run("BruteForce", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			longestPalindromeBruteForce(testString)
		}
	})
}
