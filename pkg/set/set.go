package set

type Set map[string]struct{}

func (set Set) Add(host string) {
	if _, ok := set[host]; !ok {
		set[host] = struct{}{}
	}
}
