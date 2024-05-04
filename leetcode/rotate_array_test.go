package leetcode

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantArr []int
	}{
		{
			name: "test 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			wantArr: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "test 2",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			wantArr: []int{3, 99, -1, -100},
		},
		{
			name: "test 3",
			args: args{
				nums: []int{1, 2},
				k:    0,
			},
			wantArr: []int{1, 2},
		},
		{
			name: "test 4",
			args: args{
				nums: []int{1, 2},
				k:    3,
			},
			wantArr: []int{2, 1},
		},
		{
			name: "test 5",
			args: args{
				nums: []int{1, 2},
				k:    3,
			},
			wantArr: []int{2, 1},
		},
		{
			name: "test 6",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6},
				k:    4,
			},
			wantArr: []int{3, 4, 5, 6, 1, 2},
		},
		{
			name: "test 7",
			args: args{
				nums: []int{1, 2, 3},
				k:    2,
			},
			wantArr: []int{2, 3, 1},
		},
		{
			name: "test 8",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27},
				k:    38,
			},
			wantArr: []int{17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)

			if !reflect.DeepEqual(tt.args.nums, tt.wantArr) {
				t.Errorf("rotate() = %v, want %v", tt.args.nums, tt.wantArr)
			}
		})
	}
}
