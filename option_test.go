package matchable

import (
	"testing"
)

func TestOptionIntMatcher(t *testing.T) {
	some := Some[int]{Val: 1}
	none := None[int]{}

	a := MatchOption[int](some, func(i int) int { return i }, func() int { panic("unreachable") })
	b := MatchOption[int](none, func(i int) int { panic("unreachable") }, func() int { return 0 })

	if a != 1 {
		t.Errorf("%v != Some[int] {1}", some)
	}

	if b != 0 {
		t.Errorf("%v != None[int] {}", none)
	}
}

type _Other[T any] Matchable[_Option[T]]

func TestNonOption(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Error: expected panic")
		}
	}()

	other := _Other[int]{}
	MatchOption[int](
		other,
		func(i int) int { return i },
		func() int { return 0 },
	)
}
