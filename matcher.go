package matchable

type matchable[M ~struct{}] struct{}
type Matchable[M ~struct{}] struct{ matchable[M] }
type ValuedMatchable[M ~struct{}, V any] struct {
	matchable[M]
	Val V
}

type Matcher[M ~struct{}] interface {
	matcher() matchable[M]
}

func (m matchable[M]) matcher() matchable[M] {
	return m
}
