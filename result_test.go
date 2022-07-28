package matchable

import (
	"testing"
)

func TestResultIntStringMatcher(t *testing.T) {
	ok := Ok[int, string]{Val: 1}
	err := Err[int, string]{Val: "error"}

	a := MatchResult[int, string](ok, func(i int) int { return i }, func(e string) int { panic("unreachable") })
	b := MatchResult[int, string](err, func(i int) int { panic("unreachable") }, func(e string) int { return 0 })

	if a != 1 {
		t.Errorf("%v != Ok[int, strong] {1}", ok)
	}

	if b != 5 {
		t.Errorf("%v != Err[int, string] {\"error\"}", err)
	}
}

type _OtherResult[T, E any] Matchable[_Result[T, E]]

func TestNonResult(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Error: expected panic")
		}
	}()

	other := _OtherResult[int, string]{}
	MatchResult[int, string](
		other,
		func(ok int) int { return ok },
		func(err string) int { return len(err) },
	)
}
