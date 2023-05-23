package tsmc

import "testing"

func Test_strokesRequired(t *testing.T) {
	type args struct {
		picture []string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{
				picture: []string{
					"aaaba",
					"ababa",
					"aaaca"},
			},
			want: 5,
		},
		{
			name: "case 2",
			args: args{
				picture: []string{
					"arb",
					"rrr",
					"crd",
				},
			},
			want: 5,
		},
		{
			name: "case 3",
			args: args{
				picture: []string{
					"aaa",
				},
			},
			want: 1,
		},
		{
			name: "case 4",
			args: args{
				picture: []string{
					"aaa",
					"aba",
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strokesRequired(tt.args.picture); got != tt.want {
				t.Errorf("strokesRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}
