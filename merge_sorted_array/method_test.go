package mergesortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "case 2",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{1, 5, 6},
				n:     3,
			},
			want: []int{1, 1, 2, 3, 5, 6},
		},
		{
			name: "case 3",
			args: args{
				nums1: []int{1, 2, 3, 10, 0, 0, 0},
				m:     4,
				nums2: []int{4, 5, 6},
				n:     3,
			},
			want: []int{1, 2, 3, 4, 5, 6, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)

			assert.Equal(t, tt.want, tt.args.nums1)
		})
	}
}
