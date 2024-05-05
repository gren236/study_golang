package leetcode

import (
	"reflect"
	"testing"
)

func Test_getRow(t *testing.T) {
	type args struct {
		rowIndex int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 1",
			args: args{
				rowIndex: 3,
			},
			want: []int{1, 3, 3, 1},
		},
		{
			name: "test 2",
			args: args{
				rowIndex: 0,
			},
			want: []int{1},
		},
		{
			name: "test 3",
			args: args{
				rowIndex: 1,
			},
			want: []int{1, 1},
		},
		{
			name: "test 4",
			args: args{
				rowIndex: 4,
			},
			want: []int{1, 4, 6, 4, 1},
		},
		{
			name: "",
			args: args{
				rowIndex: 21,
			},
			want: []int{1, 21, 210, 1330, 5985, 20349, 54264, 116280, 203490, 293930, 352716, 352716, 293930, 203490, 116280, 54264, 20349, 5985, 1330, 210, 21, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRow(tt.args.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
