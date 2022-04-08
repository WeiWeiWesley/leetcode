package longestpredix

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "example 1",
			args: args{
				strs: []string{"flower", "flow", "flight"},
			},
			want: "fl",
		},
		{
			name: "example 2",
			args: args{
				strs: []string{"dog","racecar","car"},
			},
			want: "",
		},
		{
			name: "example 3",
			args: args{
				strs: []string{"d","ddd","car"},
			},
			want: "",
		},
		{
			name: "example 4",
			args: args{
				strs: []string{"d","ddd","dddd"},
			},
			want: "d",
		},
		{
			name: "example 5",
			args: args{
				strs: []string{"ddd","dd","d"},
			},
			want: "d",
		},
		{
			name: "example 6",
			args: args{
				strs: []string{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
