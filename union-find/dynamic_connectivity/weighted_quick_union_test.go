package main

import "testing"

func TestWeightedQuickUnionNew(t *testing.T) {
	for _, tc := range []struct {
		name string
		n    int
		err  error
		ids  []int
	}{
		{
			name: "empty",
			err:  ErrNotPositiveN,
		},
		{
			name: "n is negative number",
			n:    -1,
			err:  ErrNotPositiveN,
		},
		{
			name: "n is positive number",
			n:    10,
			ids:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			wqu, err := NewWeightedQuickUnion(tc.n)

			if err != tc.err {
				t.Fatalf("unexpected error: got %q, wanted %q", err, tc.err)
			}

			if tc.err != nil {
				return
			}

			if len(wqu.ids) != tc.n {
				t.Fatalf("got %v, wanted %v", len(wqu.ids), tc.n)
			}

			if len(wqu.ids) != len(tc.ids) {
				t.Fatalf("got %v, wanted %v", len(wqu.ids), len(tc.ids))
			}

			for i := range wqu.ids {
				if wqu.ids[i] != tc.ids[i] {
					t.Fatalf("index %v got %v, wanted %v", i, wqu.ids[i], tc.ids[i])
				}
			}
		})
	}
}

func TestWeightedQuickUnionUnion(t *testing.T) {
	for _, tc := range []struct {
		name     string
		p, q     int
		wqu      *WeightedQuickUnion
		expected []int
	}{
		{
			name:     "union 4 3",
			p:        4,
			q:        3,
			wqu:      &WeightedQuickUnion{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			expected: []int{0, 1, 2, 4, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "union 3 8",
			p:        3,
			q:        8,
			wqu:      &WeightedQuickUnion{[]int{0, 1, 2, 4, 4, 5, 6, 7, 8, 9}, []int{0, 0, 0, 0, 1, 0, 0, 0, 0, 0}},
			expected: []int{0, 1, 2, 4, 4, 5, 6, 7, 4, 9},
		},
		{
			name:     "union 6 5",
			p:        6,
			q:        5,
			wqu:      &WeightedQuickUnion{[]int{0, 1, 2, 4, 4, 5, 6, 7, 4, 9}, []int{0, 0, 0, 0, 1, 0, 1, 0, 0, 0}},
			expected: []int{0, 1, 2, 4, 4, 6, 6, 7, 4, 9},
		},
		{
			name:     "union 9 4",
			p:        9,
			q:        4,
			wqu:      &WeightedQuickUnion{[]int{0, 1, 2, 4, 4, 6, 6, 7, 4, 9}, []int{0, 0, 0, 0, 2, 0, 1, 0, 0, 0}},
			expected: []int{0, 1, 2, 4, 4, 6, 6, 7, 4, 4},
		},
		{
			name:     "union 2 1",
			p:        2,
			q:        1,
			wqu:      &WeightedQuickUnion{[]int{0, 1, 2, 4, 4, 6, 6, 7, 4, 4}, []int{0, 0, 1, 0, 2, 0, 1, 0, 0, 0}},
			expected: []int{0, 2, 2, 4, 4, 6, 6, 7, 4, 4},
		},
		{
			name:     "union 5 0",
			p:        5,
			q:        0,
			wqu:      &WeightedQuickUnion{[]int{0, 2, 2, 4, 4, 6, 6, 7, 4, 4}, []int{0, 0, 1, 0, 2, 0, 2, 0, 0, 0}},
			expected: []int{6, 2, 2, 4, 4, 6, 6, 7, 4, 4},
		},
		{
			name:     "union 7 2",
			p:        7,
			q:        2,
			wqu:      &WeightedQuickUnion{[]int{6, 2, 2, 4, 4, 6, 6, 7, 4, 4}, []int{0, 0, 2, 0, 2, 0, 2, 0, 0, 0}},
			expected: []int{6, 2, 2, 4, 4, 6, 6, 2, 4, 4},
		},
		{
			name:     "union 6 1",
			p:        6,
			q:        1,
			wqu:      &WeightedQuickUnion{[]int{6, 2, 2, 4, 4, 6, 6, 2, 4, 4}, []int{0, 0, 2, 0, 2, 0, 3, 0, 0, 0}},
			expected: []int{6, 2, 6, 4, 4, 6, 6, 2, 4, 4},
		},
		{
			name:     "union 7 3",
			p:        7,
			q:        3,
			wqu:      &WeightedQuickUnion{[]int{6, 2, 6, 4, 4, 6, 6, 2, 4, 4}, []int{0, 0, 2, 0, 2, 0, 4, 0, 0, 0}},
			expected: []int{6, 2, 6, 4, 6, 6, 6, 2, 4, 4},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {

			tc.wqu.Union(tc.p, tc.q)

			for i := range tc.wqu.ids {
				if tc.wqu.ids[i] != tc.expected[i] {
					t.Fatalf("index %v got %v, wanted %v", i, tc.wqu.ids[i], tc.expected[i])
				}
			}
		})
	}
}

func TestWeightedQuickUnionConnected(t *testing.T) {
	for _, tc := range []struct {
		name     string
		p, q     int
		wqu      *WeightedQuickUnion
		expected bool
	}{
		{
			name:     "connected",
			p:        8,
			q:        9,
			wqu:      &WeightedQuickUnion{[]int{0, 1, 1, 8, 3, 5, 5, 7, 8, 8}, []int{}},
			expected: true,
		},
		{
			name:     "not connected",
			p:        5,
			q:        4,
			wqu:      &WeightedQuickUnion{[]int{0, 1, 1, 8, 3, 5, 5, 7, 8, 8}, []int{}},
			expected: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {

			isConnected := tc.wqu.Connected(tc.p, tc.q)
			if tc.expected != isConnected {
				t.Fatalf("got %v, wanted %v", isConnected, tc.expected)
			}

		})
	}
}

var connectedWQU bool

// BenchmarkWU - Weighted Quick-Union
func BenchmarkWU(b *testing.B) {
	var wqu *WeightedQuickUnion

	for _, bm := range []struct {
		name string
		n    int
	}{
		{"_____10", 10},
		{"____100", 100},
		{"__10_00", 1000},
		{"_10_000", 10000},
		{"100_000", 100000},
	} {
		b.Run(bm.name, func(b *testing.B) {
			wqu, _ = NewWeightedQuickUnion(bm.n)

			// allowing bechmark for array accesses removing allocation time
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				wqu.Union(bm.n-2, bm.n-1)
				connectedWQU = wqu.Connected(bm.n-1, bm.n-2)
			}
		})
	}
}
