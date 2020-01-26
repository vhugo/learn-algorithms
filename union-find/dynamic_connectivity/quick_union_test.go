package main

import "testing"

func TestQuickUnionNew(t *testing.T) {
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
			qu, err := NewQuickUnion(tc.n)

			if err != tc.err {
				t.Fatalf("unexpected error: got %q, wanted %q", err, tc.err)
			}

			if tc.err != nil {
				return
			}

			if len(qu.ids) != tc.n {
				t.Fatalf("got %v, wanted %v", len(qu.ids), tc.n)
			}

			if len(qu.ids) != len(tc.ids) {
				t.Fatalf("got %v, wanted %v", len(qu.ids), len(tc.ids))
			}

			for i := range qu.ids {
				if qu.ids[i] != tc.ids[i] {
					t.Fatalf("index %v got %v, wanted %v", i, qu.ids[i], tc.ids[i])
				}
			}
		})
	}
}

func TestQuickUnionUnion(t *testing.T) {
	for _, tc := range []struct {
		name     string
		p, q     int
		qu       *QuickUnion
		expected []int
	}{
		{
			name:     "union 4 3",
			p:        4,
			q:        3,
			qu:       &QuickUnion{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
			expected: []int{0, 1, 2, 3, 3, 5, 6, 7, 8, 9},
		},
		{
			name:     "union 3 8",
			p:        3,
			q:        8,
			qu:       &QuickUnion{[]int{0, 1, 2, 3, 3, 5, 6, 7, 8, 9}},
			expected: []int{0, 1, 2, 8, 3, 5, 6, 7, 8, 9},
		},
		{
			name:     "union 6 5",
			p:        6,
			q:        5,
			qu:       &QuickUnion{[]int{0, 1, 2, 8, 3, 5, 6, 7, 8, 9}},
			expected: []int{0, 1, 2, 8, 3, 5, 5, 7, 8, 9},
		},
		{
			name:     "union 9 4",
			p:        9,
			q:        4,
			qu:       &QuickUnion{[]int{0, 1, 2, 8, 3, 5, 5, 7, 8, 9}},
			expected: []int{0, 1, 2, 8, 3, 5, 5, 7, 8, 8},
		},
		{
			name:     "union 2 1",
			p:        2,
			q:        1,
			qu:       &QuickUnion{[]int{0, 1, 2, 8, 3, 5, 5, 7, 8, 8}},
			expected: []int{0, 1, 1, 8, 3, 5, 5, 7, 8, 8},
		},
		{
			name:     "union 5 0",
			p:        5,
			q:        0,
			qu:       &QuickUnion{[]int{0, 1, 1, 8, 3, 5, 5, 7, 8, 8}},
			expected: []int{0, 1, 1, 8, 3, 0, 5, 7, 8, 8},
		},
		{
			name:     "union 7 2",
			p:        7,
			q:        2,
			qu:       &QuickUnion{[]int{0, 1, 1, 8, 3, 0, 5, 7, 8, 8}},
			expected: []int{0, 1, 1, 8, 3, 0, 5, 1, 8, 8},
		},
		{
			name:     "union 6 1",
			p:        6,
			q:        1,
			qu:       &QuickUnion{[]int{0, 1, 1, 8, 3, 0, 5, 1, 8, 8}},
			expected: []int{1, 1, 1, 8, 3, 0, 5, 1, 8, 8},
		},
		{
			name:     "union 7 3",
			p:        7,
			q:        3,
			qu:       &QuickUnion{[]int{1, 1, 1, 8, 3, 0, 5, 1, 8, 8}},
			expected: []int{1, 8, 1, 8, 3, 0, 5, 1, 8, 8},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {

			tc.qu.Union(tc.p, tc.q)

			for i := range tc.qu.ids {
				if tc.qu.ids[i] != tc.expected[i] {
					t.Fatalf("index %v got %v, wanted %v", i, tc.qu.ids[i], tc.expected[i])
				}
			}
		})
	}
}

func TestQuickUnionConnected(t *testing.T) {
	for _, tc := range []struct {
		name     string
		p, q     int
		qu       *QuickUnion
		expected bool
	}{
		{
			name:     "connected",
			p:        8,
			q:        9,
			qu:       &QuickUnion{[]int{0, 1, 1, 8, 3, 5, 5, 7, 8, 8}},
			expected: true,
		},
		{
			name:     "not connected",
			p:        5,
			q:        4,
			qu:       &QuickUnion{[]int{0, 1, 1, 8, 3, 5, 5, 7, 8, 8}},
			expected: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {

			isConnected := tc.qu.Connected(tc.p, tc.q)
			if tc.expected != isConnected {
				t.Fatalf("got %v, wanted %v", isConnected, tc.expected)
			}

		})
	}
}

var connectedQU bool

func BenchmarkQuickUnion(b *testing.B) {
	var qu *QuickUnion

	for _, bm := range []struct {
		name string
		n    int
	}{
		{"n 3", 3},
		{"n 10", 10},
		{"n 100", 100},
		{"n 1000", 1000},
		{"n 10000", 10000},
	} {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				qu, _ = NewQuickUnion(bm.n)
				qu.Union(bm.n-2, bm.n-1)
				connectedQU = qu.Connected(bm.n-1, bm.n-2)
			}
		})
	}
}
