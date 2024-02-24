package linkedl

import (
	"reflect"
	"testing"
)

type testVal struct {
	key   int
	value int
}

func createTestLinkedList() *Singly[testVal] {
	return &Singly[testVal]{
		size: 3,
		head: &Node[testVal]{
			n: &Node[testVal]{
				n: &Node[testVal]{
					n: nil,
					v: testVal{0, 42},
				},
				v: testVal{1, 43},
			},
			v: testVal{2, 44},
		},
	}
}

func TestNewSingly(t *testing.T) {
	type args[T comparable] struct {
		vals []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want *Singly[T]
	}
	tests := []testCase[testVal]{
		{
			name: "New empty linked list created",
			args: args[testVal]{
				vals: nil,
			},
			want: &Singly[testVal]{},
		},
		{
			name: "New prefilled linked list created",
			args: args[testVal]{
				vals: []testVal{{0, 42}, {1, 43}, {2, 44}},
			},
			want: createTestLinkedList(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSingly(tt.args.vals...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSingly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingly_Search(t *testing.T) {
	type args[T any] struct {
		eq func(T) bool
	}
	type testCase[T any] struct {
		name    string
		s       *Singly[T]
		args    args[T]
		wantRes T
		wantOk  bool
	}
	tests := []testCase[testVal]{
		{
			name: "Found value",
			s:    createTestLinkedList(),
			args: args[testVal]{
				eq: func(val testVal) bool {
					return val.key == 2
				},
			},
			wantRes: testVal{2, 44},
			wantOk:  true,
		},
		{
			name: "Value not found",
			s:    createTestLinkedList(),
			args: args[testVal]{
				eq: func(val testVal) bool {
					return val.key == 5
				},
			},
			wantRes: testVal{},
			wantOk:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotOk := tt.s.Search(tt.args.eq)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Search() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Search() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
