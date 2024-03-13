package ufind

import (
	"reflect"
	"testing"
)

func createTestContainerLen7() Container[string] {
	return Container[string]{
		"a": &Node[string]{"b", 0},
		"b": &Node[string]{"c", 1},
		"c": &Node[string]{"d", 2},
		"d": &Node[string]{"d", 3},
		"e": &Node[string]{"d", 1},
		"f": &Node[string]{"e", 0},
		"g": &Node[string]{"e", 0},
	}
}

func TestNew(t *testing.T) {
	type args[T comparable] struct {
		data []T
	}
	type testCase[T comparable] struct {
		name    string
		args    args[T]
		want    Container[T]
		wantErr bool
	}
	tests := []testCase[string]{
		{
			name: "Created, 5 elements",
			args: args[string]{
				data: []string{"a", "b", "c", "d", "e"},
			},
			want: Container[string]{
				"a": &Node[string]{"a", 0},
				"b": &Node[string]{"b", 0},
				"c": &Node[string]{"c", 0},
				"d": &Node[string]{"d", 0},
				"e": &Node[string]{"e", 0},
			},
			wantErr: false,
		},
		{
			name: "Error: nil input slice",
			args: args[string]{
				data: nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainer_Find(t *testing.T) {
	type args[T comparable] struct {
		x T
	}
	type testCase[T comparable] struct {
		name    string
		c       Container[T]
		args    args[T]
		wantRes T
		wantC   Container[T]
		wantOk  bool
	}
	tests := []testCase[string]{
		{
			name: "Found, path compressed",
			c:    createTestContainerLen7(),
			args: args[string]{
				x: "a",
			},
			wantRes: "d",
			wantC: Container[string]{
				"a": &Node[string]{"d", 0},
				"b": &Node[string]{"d", 1},
				"c": &Node[string]{"d", 2},
				"d": &Node[string]{"d", 3},
				"e": &Node[string]{"d", 1},
				"f": &Node[string]{"e", 0},
				"g": &Node[string]{"e", 0},
			},
			wantOk: true,
		},
		{
			name: "Not found",
			c:    createTestContainerLen7(),
			args: args[string]{
				x: "x",
			},
			wantRes: "",
			wantC:   createTestContainerLen7(),
			wantOk:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotOk := tt.c.Find(tt.args.x)

			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Find() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(tt.c, tt.wantC) {
				t.Errorf("Find() got container = %v, want %v", tt.c, tt.wantC)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Find() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestContainer_Union(t *testing.T) {
	type args[T comparable] struct {
		x T
		y T
	}
	type testCase[T comparable] struct {
		name  string
		c     Container[T]
		args  args[T]
		want  bool
		wantC Container[T]
	}
	tests := []testCase[string]{
		{
			// TODO
			name:  "Merged 2 groups",
			c:     nil,
			args:  args[string]{},
			want:  false,
			wantC: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.c.Union(tt.args.x, tt.args.y)

			if got != tt.want {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.c, tt.wantC) {
				t.Errorf("Union() got container = %v, want %v", tt.c, tt.wantC)
			}
		})
	}
}
