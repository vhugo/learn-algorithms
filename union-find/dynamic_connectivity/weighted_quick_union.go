package main

// WeightedQuickUnion represents the lazy approach for the Union-Find dynamic connectivity problem
type WeightedQuickUnion struct {
	ids  []int
	size []int
}

// NewQuickUnion set id of each object to itself (N array accesses)
func NewWeightedQuickUnion(n int) (*WeightedQuickUnion, error) {
	if n < 1 {
		return nil, ErrNotPositiveN
	}

	ids := make([]int, n)
	for i := range ids {
		ids[i] = i
	}

	return &WeightedQuickUnion{
		ids:  ids,
		size: make([]int, n),
	}, nil
}

// Union change root of p to point to root of q (depth of p and q array accesses)
func (u *WeightedQuickUnion) Union(p, q int) {
	i := u.root(p)
	j := u.root(q)

	if i == j {
		return
	}

	if u.size[i] < u.size[j] {
		u.ids[i] = j
		u.size[j] += u.size[i]
		// u.size[j]++
		return
	}

	u.ids[j] = i
	u.size[i] += u.size[j]
	// u.size[i]++
}

// Connected check if p and q have same root (depth of p and q array accesses)
func (u *WeightedQuickUnion) Connected(p, q int) bool {
	return u.root(p) == u.root(q)
}

// root chase parent pointers until reach root (depth of i array accesses)
func (u *WeightedQuickUnion) root(i int) int {
	for i != u.ids[i] {
		i = u.ids[i]
	}
	return i
}
