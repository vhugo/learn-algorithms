package main

// QuickUnion represents the lazy approach for the Union-Find dynamic connectivity problem
type QuickUnion struct {
	ids []int
}

// NewQuickUnion set id of each object to itself (N array accesses)
func NewQuickUnion(n int) (*QuickUnion, error) {
	if n < 1 {
		return nil, ErrNotPositiveN
	}

	ids := make([]int, n)
	for i := range ids {
		ids[i] = i
	}

	return &QuickUnion{
		ids: ids,
	}, nil
}

// Union change root of p to point to root of q (depth of p and q array accesses)
func (u *QuickUnion) Union(p, q int) {
	i := u.root(p)
	j := u.root(q)
	u.ids[i] = j
}

// Connected check if p and q have same root (depth of p and q array accesses)
func (u *QuickUnion) Connected(p, q int) bool {
	return u.root(p) == u.root(q)
}

// root chase parent pointers until reach root (depth of i array accesses)
func (u *QuickUnion) root(i int) int {
	for i != u.ids[i] {
		i = u.ids[i]
	}
	return i

	// this was my attempt to implement:
	// root := u.ids[i]
	// if root == u.ids[root] {
	// 	return root
	// }
	// return u.root(root)
}
