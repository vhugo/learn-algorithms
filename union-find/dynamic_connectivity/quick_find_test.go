package main

import "testing"

func TestQuickFindNew(t *testing.T) {
	for _, tc := range []struct {
		m   string
		n   int
		err error
		ids []int
	}{
		{
			m:   "empty",
			err: ErrNotPositiveN,
		},
		{
			m:   "n is negative number",
			n:   -1,
			err: ErrNotPositiveN,
		},
		{
			m:   "n is positive number",
			n:   10,
			ids: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	} {
		t.Run(tc.m, func(t *testing.T) {
			qf, err := NewQuickFind(tc.n)

			if err != tc.err {
				t.Fatalf("unexpected error: got %q, wanted %q", err, tc.err)
			}

			if tc.err != nil {
				return
			}

			if len(qf.ids) != tc.n {
				t.Fatalf("got %v, wanted %v", len(qf.ids), tc.n)
			}

			if len(qf.ids) != len(tc.ids) {
				t.Fatalf("got %v, wanted %v", len(qf.ids), len(tc.ids))
			}

			for i := range qf.ids {
				if qf.ids[i] != tc.ids[i] {
					t.Fatalf("index %v got %v, wanted %v", i, qf.ids[i], tc.ids[i])
				}
			}
		})
	}
}

func TestQuickFindUnion(t *testing.T) {
	for _, tc := range []struct {
		m        string
		p, q     int
		qf       *QuickFind
		expected []int
	}{
		{
			m:        "replace single term",
			p:        1,
			q:        2,
			qf:       &QuickFind{[]int{0, 1, 2, 3}},
			expected: []int{0, 2, 2, 3},
		},
		{
			m:        "replace multiple terms",
			p:        2,
			q:        6,
			qf:       &QuickFind{[]int{0, 1, 1, 3, 4, 1, 6, 7, 8, 9}},
			expected: []int{0, 6, 6, 3, 4, 6, 6, 7, 8, 9},
		},
	} {
		t.Run(tc.m, func(t *testing.T) {

			tc.qf.Union(tc.p, tc.q)

			for i := range tc.qf.ids {
				if tc.qf.ids[i] != tc.expected[i] {
					t.Fatalf("index %v got %v, wanted %v", i, tc.qf.ids[i], tc.expected[i])
				}
			}
		})
	}
}

func TestQuickFindConnected(t *testing.T) {
	for _, tc := range []struct {
		m        string
		p, q     int
		qf       *QuickFind
		expected bool
	}{
		{
			m:        "not connected",
			p:        1,
			q:        2,
			qf:       &QuickFind{[]int{0, 1, 2, 3}},
			expected: false,
		},
		{
			m:        "connected",
			p:        1,
			q:        2,
			qf:       &QuickFind{[]int{0, 1, 1, 3}},
			expected: true,
		},
	} {
		t.Run(tc.m, func(t *testing.T) {

			isConnected := tc.qf.Connected(tc.p, tc.q)
			if tc.expected != isConnected {
				t.Fatalf("got %v, wanted %v", isConnected, tc.expected)
			}

		})
	}
}
