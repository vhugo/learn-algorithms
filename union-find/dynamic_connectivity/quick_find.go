package main

import "fmt"

// QuickFind eager approach for the Union-Find dynamic connectivity problem
type QuickFind struct {
	ids []int
}

// ErrNotPositiveN used to handle when N is not a positive number
var ErrNotPositiveN = fmt.Errorf("N must be a positive number")

// NewQuickFind set id of each object to itself (N array accesses)
func NewQuickFind(n int) (*QuickFind, error) {
	if n < 1 {
		return nil, ErrNotPositiveN
	}

	ids := make([]int, n)
	for i := range ids {
		ids[i] = i
	}

	return &QuickFind{
		ids: ids,
	}, nil
}

// Union change all entries with id[p] to id[q] (at most 2N + 2 array accesses)
func (u *QuickFind) Union(p, q int) {
	pid := u.ids[p]
	qid := u.ids[q]

	for i := range u.ids {
		if u.ids[i] == pid {
			u.ids[i] = qid
		}
	}
}

// Connected check whether p and q are in the same component (2 array accesses)
func (u *QuickFind) Connected(p, q int) bool {
	return u.ids[p] == u.ids[q]
}
