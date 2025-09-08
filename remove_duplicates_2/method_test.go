package removeduplicates2

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1 - 基本测试",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
			},
			want: 5,
		},
		{
			name: "case2 - 无重复超过2次",
			args: args{
				nums: []int{1, 1, 2, 2},
			},
			want: 4,
		},
		{
			name: "case3 - 单个元素",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "case4 - 空数组",
			args: args{
				nums: []int{},
			},
			want: 0,
		},
		{
			name: "case5 - 所有元素相同且超过2个",
			args: args{
				nums: []int{1, 1, 1, 1, 1},
			},
			want: 2,
		},
		{
			name: "case6 - 复杂情况",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			},
			want: 7,
		},
		{
			name: "case7 - 只有两个相同元素",
			args: args{
				nums: []int{1, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
