package leetcode

import "testing"

func Test_reverseLetters(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 1",
			args: args{
				s: "Let's take LeetCode contest",
			},
			want: "s'teL ekat edoCteeL tsetnoc",
		},
		{
			name: "test 2",
			args: args{
				s: "Mr Ding",
			},
			want: "rM gniD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseLetters(tt.args.s); got != tt.want {
				t.Errorf("reverseLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
