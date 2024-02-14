package heap

import (
	"reflect"
	"testing"
)

type testVal int

func (tv testVal) Less(v testVal) bool {
	return tv < v
}

func TestNewHeapFromSlice(t *testing.T) {
	type args[T Comparable[T]] struct {
		arr []T
	}
	type testCase[T Comparable[T]] struct {
		name string
		args args[T]
		want *Container[T]
	}
	tests := []testCase[testVal]{
		{
			name: "Regular sized: reversed input",
			args: args[testVal]{
				arr: []testVal{9, 8, 7, 6, 5, 4, 3, 2, 1},
			},
			want: &Container[testVal]{1, 2, 3, 6, 5, 4, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHeapFromSlice(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHeapFromSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_container_ExtractMin(t *testing.T) {
	tests := []struct {
		name  string
		c     Container[testVal]
		wantC Container[testVal]
		want  testVal
	}{
		{
			name:  "Regular size",
			c:     Container[testVal]{4, 4, 8, 9, 4, 12, 9, 11, 13},
			wantC: Container[testVal]{4, 4, 8, 9, 13, 12, 9, 11},
			want:  4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.c.ExtractMin()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractMin() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.c, tt.wantC) {
				t.Errorf("ExtractMin() Container = %v, want Container %v", tt.c, tt.wantC)
			}
		})
	}
}

func Test_container_Insert(t *testing.T) {
	type args struct {
		v testVal
	}
	tests := []struct {
		name  string
		c     Container[testVal]
		args  args
		wantC Container[testVal]
	}{
		{
			name: "Regular size: new min",
			c:    Container[testVal]{4, 4, 8, 9, 4, 12, 9, 11, 13},
			args: args{
				v: 3,
			},
			wantC: Container[testVal]{3, 4, 8, 9, 4, 12, 9, 11, 13, 4},
		},
		{
			name: "Regular size: do one swap",
			c:    Container[testVal]{4, 4, 8, 9, 6, 12, 9, 11, 13},
			args: args{
				v: 5,
			},
			wantC: Container[testVal]{4, 4, 8, 9, 5, 12, 9, 11, 13, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Insert(tt.args.v)

			if !reflect.DeepEqual(tt.c, tt.wantC) {
				t.Errorf("ExtractMin() Container = %v, want Container %v", tt.c, tt.wantC)
			}
		})
	}
}
