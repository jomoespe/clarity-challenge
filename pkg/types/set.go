package types

type Set map[string]struct{}

// Add values to a Set
func (set Set) Add(values ...string) {
	for _, value := range values {
		if _, ok := set[value]; !ok {
			set[value] = struct{}{}
		}
	}
}

// Clean removes all Set elements
func (set Set) Clean() {
	for k := range set {
		delete(set, k)
	}
}