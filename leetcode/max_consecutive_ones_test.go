package leetcode

import "testing"

func Test_findMaxConsecutiveOnes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				nums: []int{1, 1, 0, 1, 1, 1},
			},
			want: 3,
		},
		{
			name: "test 2",
			args: args{
				nums: []int{1, 0, 1, 1, 0, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxConsecutiveOnes(tt.args.nums); got != tt.want {
				t.Errorf("findMaxConsecutiveOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
