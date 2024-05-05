package leetcode

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantArr []int
	}{
		{
			name: "test 1",
			args: args{
				nums: []int{0, 1, 0, 3, 12},
			},
			wantArr: []int{1, 3, 12, 0, 0},
		},
		{
			name: "test 2",
			args: args{
				nums: []int{0},
			},
			wantArr: []int{0},
		},
		{
			name: "test 3",
			args: args{
				nums: []int{1, 0, 1},
			},
			wantArr: []int{1, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			got := tt.args.nums

			if !reflect.DeepEqual(got, tt.wantArr) {
				t.Errorf("moveZeroes() got = %v, want %v", got, tt.wantArr)
			}
		})
	}
}
