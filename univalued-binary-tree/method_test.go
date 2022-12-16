package univaluedbinarytree

import "testing"

//go test -timeout 30s -run ^Test_isUnivalTree$ ./ -v
func Test_isUnivalTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		// -1 replace null
		{
			name: "example 1",
			args: args{CreateTree([]int{1, 1, 1, 1, 1, -1, 1})},
			want: true,
		},
		{
			name: "example 2",
			args: args{CreateTree([]int{2, 2, 2, 5, 2})},
			want: false,
		},
		{
			name: "example 3",
			args: args{CreateTree([]int{})},
			want: true,
		},
		{
			name: "example 4",
			args: args{CreateTree([]int{2})},
			want: true,
		},
		{
			name: "example 5",
			args: args{CreateTree([]int{1, 2, 3, 2, 2, 2, 2, -1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2})},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnivalTree(tt.args.root); got != tt.want {
				t.Errorf("isUnivalTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
