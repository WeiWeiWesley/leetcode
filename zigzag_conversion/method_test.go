package zigzagconversion

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{"PAYPALISHIRING", 3},
			want: "PAHNAPLSIIGYIR",
		},
		{
			name: "example 2",
			args: args{"PAYPALISHIRING", 4},
			want: "PINALSIGYAHRPI",
		},
		{
			name: "example 3",
			args: args{"AB", 1},
			want: "AB",
		},
		{
			name: "example 4",
			args: args{"ABC", 1},
			want: "ABC",
		},
		{
			name: "example 5",
			args: args{"A", 1},
			want: "A",
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
