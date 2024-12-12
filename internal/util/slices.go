package util

// GroupBy groups the values of s according to the key selector getKey.
func GroupBy[S ~[]V, K comparable, V any](s S, getKey func(v V) K) map[K]S {
	groups := make(map[K]S)
	for _, v := range s {
		k := getKey(v)
		groups[k] = append(groups[k], v)
	}
	return groups
}
