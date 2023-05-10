package continerwithmostwater

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
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
				height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			want: 49,
		},
		{
			name: "case 2",
			args: args{
				height: []int{1, 1},
			},
			want: 1,
		},
		{
			name: "case 3",
			args: args{
				height: []int{1, 1, 1, 1, 1, 3},
			},
			want: 5,
		},
		{
			name: "case 4",
			args: args{
				height: []int{3, 1, 1, 1, 1, 1},
			},
			want: 5,
		},
		{
			name: "case 5",
			args: args{
				height: []int{1, 1, 3, 1, 1, 1},
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
