package stringtointerger

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
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
				s: "42",
			},
			want: 42,
		},
		{
			name: "case 2",
			args: args{
				s: "   -42",
			},
			want: -42,
		},
		{
			name: "case 3",
			args: args{
				s: "4193 with words",
			},
			want: 4193,
		},
		{
			name: "case 5",
			args: args{
				s: "00000-42a1234",
			},
			want: 0,
		},
		{
			name: "case 6",
			args: args{
				s: "   +0 123",
			},
			want: 0,
		},
		{
			name: "case 7",
			args: args{
				s: "20000000000000000000",
			},
			want: 2147483647,
		},
		{
			name: "case 8",
			args: args{
				s: "-91283472332",
			},
			want: -2147483648,
		},
		{
			name: "case 9",
			args: args{
				s: "9223372036854775808",
			},
			want: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
