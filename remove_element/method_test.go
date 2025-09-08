package removeelement

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: 2,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want: 5,
		},
		{
			name: "not found",
			args: args{
				nums: []int{-1, 1, 2, 3},
				val:  8,
			},
			want: 4,
		},
		{
			name: "empty",
			args: args{
				nums: []int{},
				val:  0,
			},
			want: 0,
		},
		{
			name: "first element",
			args: args{
				nums: []int{1},
				val:  1,
			},
			want: 0,
		},
		{
			name: "last element",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				val:  5,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
