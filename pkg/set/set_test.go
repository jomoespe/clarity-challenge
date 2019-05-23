package set_test

import (
	"testing"

	"github.com/jomoespe/clarity-challenge/pkg/set"
)

func TestSomething(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	set := set.Set{}
	set.Add("one")
	set.Add("two")
	set.Add("three")
	set.Add("three")
	if len(set) != 3 {
		t.Errorf("wrong set size. Expected: %d, Got: %d", 3, len(set))
	}
}
