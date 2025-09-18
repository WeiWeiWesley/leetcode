package findindexoffirstoccurrence

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				haystack: "sadbutsad",
				needle:   "sad",
			},
			want: 0,
		},
		{
			name: "case 2",
			args: args{
				haystack: "leetcode",
				needle:   "leeto",
			},
			want: -1,
		},
		{
			name: "case 3",
			args: args{
				haystack: "sadbutsad",
				needle:   "tsa",
			},
			want: 5,
		},
		{
			name: "case 4",
			args: args{
				haystack: "mississippi",
				needle:   "a",
			},
			want: -1,
		},
		{
			name: "case 5 - empty needle",
			args: args{
				haystack: "hello",
				needle:   "",
			},
			want: 0,
		},
		{
			name: "case 6 - needle equals haystack",
			args: args{
				haystack: "hello",
				needle:   "hello",
			},
			want: 0,
		},
		{
			name: "case 7 - single character match",
			args: args{
				haystack: "hello",
				needle:   "l",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
