package set

type Set map[string]struct{}

func (set Set) Add(values ...string) {
	for _, value := range values {
		if _, ok := set[value]; !ok {
			set[value] = struct{}{}
		}
	}
}
