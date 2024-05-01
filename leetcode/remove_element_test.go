package leetcode

import (
	"reflect"
	"slices"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantArr []int
	}{
		{
			name: "test 1",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want:    2,
			wantArr: []int{2, 2},
		},
		{
			name: "test 2",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want:    5,
			wantArr: []int{0, 1, 4, 0, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.args.nums, tt.args.val)
			gotArr := tt.args.nums[:got]

			if got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}

			slices.Sort(gotArr)
			slices.Sort(tt.wantArr)

			if !reflect.DeepEqual(gotArr, tt.wantArr) {
				t.Errorf("removeElement() = %v, want %v", gotArr, tt.wantArr)
			}
		})
	}
}
