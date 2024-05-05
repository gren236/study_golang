package leetcode

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{1, 1, 2},
			},
			want:    2,
			wantArr: []int{1, 2},
		},
		{
			name: "test 2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want:    5,
			wantArr: []int{0, 1, 2, 3, 4},
		},
		{
			name: "test 3",
			args: args{
				nums: []int{1, 1},
			},
			want:    1,
			wantArr: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.nums)
			gotArr := tt.args.nums[:got]

			if got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(gotArr, tt.wantArr) {
				t.Errorf("removeDuplicates() = %v, want %v", gotArr, tt.wantArr)
			}
		})
	}
}
