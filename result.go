package matchable

type _Result[T, E any] struct{}
type Result[T, E any] Matcher[_Result[T, E]]

type Ok[T, E any] ValuedMatchable[_Result[T, E], T]
type Err[T, E any] ValuedMatchable[_Result[T, E], E]

func MatchResult[T, E, R any](
	result Result[T, E],
	ok func(T) R,
	err func(E) R,
) R {
	switch v := result.(type) {
	case Ok[T, E]:
		return ok(v.Val)
	case Err[T, E]:
		return err(v.Val)
	}
	panic("Option type not supported")
}
