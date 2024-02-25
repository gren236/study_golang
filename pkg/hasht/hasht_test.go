package hasht

import (
	"reflect"
	"testing"
)

type testKey string

func (t testKey) Bytes() []byte {
	return []byte(t)
}

func createSmallTable() *Table[testKey, int] {
	in := []*container[testKey, int]{
		{"test1", 1},
		{"test2", 2},
		{"test3", 3},
		{"test4", 4},
		{"test5", 5},
	}

	t := New[testKey, int]()

	for _, c := range in {
		insertContainer(*t.buckets, c)
	}

	t.objNum = 5

	return t
}

func TestTable_Insert(t *testing.T) {
	type args[T key, U any] struct {
		k T
		v U
	}
	type testCase[T key, U any] struct {
		name      string
		table     *Table[T, U]
		args      args[T, U]
		wantTable *Table[T, U]
	}
	tests := []testCase[testKey, int]{
		{
			name:      "Small table, no buckets grow",
			table:     createSmallTable(),
			wantTable: nil,
			args: args[testKey, int]{
				k: "test6",
				v: 6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			tt.table.Insert(tt.args.k, tt.args.v)

			if !reflect.DeepEqual(tt.table, tt.wantTable) {
				t.Errorf("Delete() got list = %v, want %v", tt.table, tt.wantTable)
			}
		})
	}
}
