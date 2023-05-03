package reverseinterger

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
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
				x: 543,
			},
			want: 345,
		},
		{
			name: "case 2",
			args: args{
				x: 111,
			},
			want: 111,
		},
		{
			name: "case 3",
			args: args{
				x: -123,
			},
			want: -321,
		},
		{
			name: "case 4",
			args: args{
				x: 1534236469,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
