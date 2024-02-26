package hasht

import (
	"fmt"
	"reflect"
	"testing"
)

type testKey string

func (t testKey) Bytes() []byte {
	return []byte(t)
}

func createTestTable(size int) *Table[testKey, int] {
	t := New[testKey, int]()

	if size > 11 {
		t.buckets = newBuckets[testKey, int](initialSize * growCoefficient)
	}

	for i := range size {
		c := &container[testKey, int]{
			key: testKey(fmt.Sprintf("test%d", i+1)),
			val: i + 1,
		}

		insertContainer(*t.buckets, c)
	}

	t.objNum = size

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
			name:  "Small table, no buckets' grow",
			table: createTestTable(5),
			wantTable: func() *Table[testKey, int] {
				st := createTestTable(5)
				insertContainer(*st.buckets, &container[testKey, int]{"test6", 6})
				st.objNum++

				return st
			}(),
			args: args[testKey, int]{
				k: "test6",
				v: 6,
			},
		},
		{
			name:  "Empty table, no buckets grow",
			table: New[testKey, int](),
			wantTable: func() *Table[testKey, int] {
				nt := New[testKey, int]()
				insertContainer(*nt.buckets, &container[testKey, int]{"test6", 6})
				nt.objNum++

				return nt
			}(),
			args: args[testKey, int]{
				k: "test6",
				v: 6,
			},
		},
		{
			name:      "Small table, element already in the bucket - do not change",
			table:     createTestTable(5),
			wantTable: createTestTable(5),
			args: args[testKey, int]{
				k: "test2",
				v: 2,
			},
		},
		{
			name:  "Big table, buckets array growing",
			table: createTestTable(12),
			wantTable: func() *Table[testKey, int] {
				st := createTestTable(12)
				insertContainer(*st.buckets, &container[testKey, int]{"test13", 13})
				st.objNum++

				return st
			}(),
			args: args[testKey, int]{
				k: "test13",
				v: 13,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.table.Insert(tt.args.k, tt.args.v)

			if !reflect.DeepEqual(tt.table, tt.wantTable) {
				t.Errorf("Insert() got table = %v, want %v", tt.table, tt.wantTable)
			}
		})
	}
}

func TestTable_Delete(t *testing.T) {
	type args[T key, U any] struct {
		k T
	}
	type testCase[T key, U any] struct {
		name      string
		table     *Table[T, U]
		args      args[T, U]
		wantTable *Table[T, U]
	}
	tests := []testCase[testKey, int]{
		{
			name:      "Small table, value found and deleted",
			table:     createTestTable(6),
			wantTable: createTestTable(5),
			args: args[testKey, int]{
				k: "test6",
			},
		},
		{
			name:      "Small table, value not found",
			table:     createTestTable(6),
			wantTable: createTestTable(6),
			args: args[testKey, int]{
				k: "test7",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.table.Delete(tt.args.k)

			if !reflect.DeepEqual(tt.table, tt.wantTable) {
				t.Errorf("Delete() got table = %v, want %v", tt.table, tt.wantTable)
			}
		})
	}
}

func TestTable_Get(t *testing.T) {
	type args[T key] struct {
		k T
	}
	type testCase[T key, U any] struct {
		name    string
		table   *Table[T, U]
		args    args[T]
		wantRes U
		wantOk  bool
	}
	tests := []testCase[testKey, int]{
		{
			name:  "Small table, key found",
			table: createTestTable(5),
			args: args[testKey]{
				k: "test3",
			},
			wantRes: 3,
			wantOk:  true,
		},
		{
			name:  "Small table, key not found",
			table: createTestTable(5),
			args: args[testKey]{
				k: "test6",
			},
			wantRes: 0,
			wantOk:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotOk := tt.table.Get(tt.args.k)

			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Get() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Get() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
