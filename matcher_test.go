package matchable

import (
	"testing"
)

func TestMatcherMethod(t *testing.T) {
	someInt1 := Some[int]{Val: 1}
	someInt2 := Some[int]{Val: 2}
	noneInt := None[int]{}

	matcherInt1 := someInt1.matcher()
	matcherInt2 := someInt2.matcher()
	matcherInt3 := noneInt.matcher()

	if matcherInt1 != matcherInt2 || matcherInt1 != matcherInt3 {
		t.Errorf("matcher1 != matcher2 || matcher1 != matcher3")
	}

	someString1 := Some[string]{Val: "1"}
	someString2 := Some[string]{Val: "2"}
	noneString := None[string]{}
	matcherString1 := someString1.matcher()
	matcherString2 := someString2.matcher()
	matcherString3 := noneString.matcher()

	if matcherString1 != matcherString2 || matcherString1 != matcherString3 {
		t.Errorf("matcher1 != matcher2 || matcher1 != matcher3")
	}

	// This comparison is not possible because they are of different types.
	//if matcherInt1 == matcherString1 {
	//	t.Errorf("matcher1 == matcher2")
	//}
}
