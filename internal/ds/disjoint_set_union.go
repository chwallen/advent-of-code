package ds

type DisjointSetUnion[T comparable] struct {
	parent map[T]T
	size   map[T]int
	count  int
}

func NewDisjointSetUnion[T comparable](count int) *DisjointSetUnion[T] {
	return &DisjointSetUnion[T]{
		parent: make(map[T]T),
		size:   make(map[T]int),
		count:  count,
	}
}

func (dsu *DisjointSetUnion[T]) Find(c T) T {
	p, ok := dsu.parent[c]
	if !ok {
		dsu.parent[c] = c
		dsu.size[c] = 1
		return c
	}
	if p == c {
		return c
	}

	p = dsu.Find(p)
	dsu.parent[c] = p
	return p
}

func (dsu *DisjointSetUnion[T]) Union(a, b T) bool {
	rootA := dsu.Find(a)
	rootB := dsu.Find(b)

	if rootA == rootB {
		return false
	}

	if dsu.size[rootA] < dsu.size[rootB] {
		rootA, rootB = rootB, rootA
	}

	dsu.parent[rootB] = rootA
	dsu.size[rootA] += dsu.size[rootB]
	dsu.count--

	return true
}

func (dsu *DisjointSetUnion[T]) IsRoot(n T) bool {
	r, ok := dsu.parent[n]
	return ok && r == n
}

func (dsu *DisjointSetUnion[T]) GetSize(n T) int {
	return dsu.size[n]
}

func (dsu *DisjointSetUnion[T]) Count() int {
	return dsu.count
}
