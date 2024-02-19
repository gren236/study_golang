package bst

import (
	"reflect"
	"testing"
)

type testVal int

func (v testVal) Less(k testVal) bool {
	return v < k
}

func (v testVal) Equals(k testVal) bool {
	return v == k
}

func createSmallBST() *Node[testVal] {
	root := &Node[testVal]{val: 3}

	v1 := &Node[testVal]{val: 1, p: root}
	v2 := &Node[testVal]{val: 2, p: v1}
	v1.r = v2

	v5 := &Node[testVal]{val: 5, p: root}
	v4 := &Node[testVal]{val: 4, p: v5}
	v5.l = v4

	root.l, root.r = v1, v5

	return root
}

func createSmallBSTWithInsertion() *Node[testVal] {
	root := createSmallBST()

	v5 := root.r
	v5.r = &Node[testVal]{val: 6, p: v5}

	return root
}

func createSmallBSTWithoutLeaf() *Node[testVal] {
	root := &Node[testVal]{val: 3}

	v1 := &Node[testVal]{val: 1, p: root}

	v5 := &Node[testVal]{val: 5, p: root}
	v4 := &Node[testVal]{val: 4, p: v5}
	v5.l = v4

	root.l, root.r = v1, v5

	return root
}

func createSmallBSTWithoutSideNode() *Node[testVal] {
	root := &Node[testVal]{val: 3}

	v1 := &Node[testVal]{val: 1, p: root}
	v2 := &Node[testVal]{val: 2, p: v1}
	v1.r = v2

	v4 := &Node[testVal]{val: 4, p: root}

	root.l, root.r = v1, v4

	return root
}

func createSmallBSTWithoutRoot() *Node[testVal] {
	root := &Node[testVal]{val: 2}

	v1 := &Node[testVal]{val: 1, p: root}

	v5 := &Node[testVal]{val: 5, p: root}
	v4 := &Node[testVal]{val: 4, p: v5}
	v5.l = v4

	root.l, root.r = v1, v5

	return root
}

func TestNode_Insert(t *testing.T) {
	type args[T Comparable[T]] struct {
		k T
	}
	type testCase[T Comparable[T]] struct {
		name      string
		n         *Node[T]
		args      args[T]
		wantRes   testVal
		wantGraph *Node[T]
	}
	tests := []testCase[testVal]{
		{
			name: "Small graph: graph changed",
			n:    createSmallBST(),
			args: args[testVal]{
				k: 6,
			},
			wantRes:   testVal(6),
			wantGraph: createSmallBSTWithInsertion(),
		},
		{
			name: "Small graph: value found",
			n:    createSmallBST(),
			args: args[testVal]{
				k: 2,
			},
			wantRes:   testVal(2),
			wantGraph: createSmallBST(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes := tt.n.Insert(tt.args.k)

			if !reflect.DeepEqual(gotRes.val, tt.wantRes) {
				t.Errorf("Insert() gotRes = %v, want %v", gotRes.val, tt.wantRes)
			}
			if !reflect.DeepEqual(tt.n, tt.wantGraph) {
				t.Errorf("Insert() got graph = %v, want %v", tt.n, tt.wantGraph)
			}
		})
	}
}

func TestNode_Search(t *testing.T) {
	type args[T Comparable[T]] struct {
		k T
	}
	type testCase[T Comparable[T]] struct {
		name    string
		n       *Node[T]
		args    args[T]
		wantRes testVal
		wantOk  bool
	}
	tests := []testCase[testVal]{
		{
			name: "Small graph",
			n:    createSmallBST(),
			args: args[testVal]{
				k: 2,
			},
			wantRes: testVal(2),
			wantOk:  true,
		},
		{
			name: "Small graph: non existent value",
			n:    createSmallBST(),
			args: args[testVal]{
				k: 0,
			},
			wantRes: testVal(0),
			wantOk:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotOk := tt.n.Search(tt.args.k)

			if gotRes != nil && !reflect.DeepEqual(gotRes.val, tt.wantRes) {
				t.Errorf("Search() gotRes = %v, want %v", gotRes.val, tt.wantRes)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Search() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestNode_Delete(t *testing.T) {
	type args[T Comparable[T]] struct {
		k T
	}
	type testCase[T Comparable[T]] struct {
		name      string
		n         *Node[T]
		args      args[T]
		wantGraph *Node[T]
		wantOk    bool
	}
	tests := []testCase[testVal]{
		{
			name: "Small graph: delete leaf",
			n:    createSmallBST(),
			args: args[testVal]{
				k: 2,
			},
			wantOk:    true,
			wantGraph: createSmallBSTWithoutLeaf(),
		},
		{
			name: "Small graph: delete side node",
			n:    createSmallBST(),
			args: args[testVal]{
				k: 5,
			},
			wantOk:    true,
			wantGraph: createSmallBSTWithoutSideNode(),
		},
		{
			name: "Small graph: delete root",
			n:    createSmallBST(),
			args: args[testVal]{
				k: 3,
			},
			wantOk:    true,
			wantGraph: createSmallBSTWithoutRoot(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotOk := tt.n.Delete(tt.args.k)

			if !reflect.DeepEqual(gotOk, tt.wantOk) {
				t.Errorf("Delete() gotRes = %v, want %v", gotOk, tt.wantOk)
			}
			if !reflect.DeepEqual(gotRes, tt.wantGraph) {
				t.Errorf("Delete() got graph = %v, want %v", tt.n, tt.wantGraph)
			}
		})
	}
}

func TestNode_Min(t *testing.T) {
	type testCase[T Comparable[T]] struct {
		name    string
		n       *Node[T]
		wantRes testVal
	}
	tests := []testCase[testVal]{
		{
			name:    "Small graph",
			n:       createSmallBST(),
			wantRes: testVal(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes := tt.n.Min()

			if !reflect.DeepEqual(gotRes.val, tt.wantRes) {
				t.Errorf("Min() = %v, want %v", gotRes.val, tt.wantRes)
			}
		})
	}
}

func TestNode_Max(t *testing.T) {
	type testCase[T Comparable[T]] struct {
		name    string
		n       *Node[T]
		wantRes testVal
	}
	tests := []testCase[testVal]{
		{
			name:    "Small graph",
			n:       createSmallBST(),
			wantRes: testVal(5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes := tt.n.Max()

			if !reflect.DeepEqual(gotRes.val, tt.wantRes) {
				t.Errorf("Max() = %v, want %v", gotRes.val, tt.wantRes)
			}
		})
	}
}

func TestNode_Pred(t *testing.T) {
	type testCase[T Comparable[T]] struct {
		name    string
		n       *Node[T]
		wantRes testVal
		wantOk  bool
	}
	tests := []testCase[testVal]{
		{
			name:    "Small graph: found",
			n:       createSmallBST(),
			wantRes: testVal(2),
			wantOk:  true,
		},
		{
			name: "Small graph: found from leaf value",
			n: func() *Node[testVal] {
				res, _ := createSmallBST().Search(4)

				return res
			}(),
			wantRes: testVal(3),
			wantOk:  true,
		},
		{
			name: "Small graph: not found",
			n: func() *Node[testVal] {
				res, _ := createSmallBST().Search(1)

				return res
			}(),
			wantRes: testVal(0),
			wantOk:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotOk := tt.n.Pred()

			if gotRes != nil && !reflect.DeepEqual(gotRes.val, tt.wantRes) {
				t.Errorf("Pred() gotRes = %v, want %v", gotRes.val, tt.wantRes)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Pred() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestNode_Succ(t *testing.T) {
	type testCase[T Comparable[T]] struct {
		name    string
		n       *Node[T]
		wantRes testVal
		wantOk  bool
	}
	tests := []testCase[testVal]{
		{
			name:    "Small graph: found",
			n:       createSmallBST(),
			wantRes: testVal(4),
			wantOk:  true,
		},
		{
			name: "Small graph: found from leaf value",
			n: func() *Node[testVal] {
				res, _ := createSmallBST().Search(2)

				return res
			}(),
			wantRes: testVal(3),
			wantOk:  true,
		},
		{
			name: "Small graph: not found",
			n: func() *Node[testVal] {
				res, _ := createSmallBST().Search(5)

				return res
			}(),
			wantRes: testVal(0),
			wantOk:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotOk := tt.n.Succ()

			if gotRes != nil && !reflect.DeepEqual(gotRes.val, tt.wantRes) {
				t.Errorf("Succ() gotRes = %v, want %v", gotRes.val, tt.wantRes)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Succ() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
