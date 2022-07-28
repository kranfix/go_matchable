package matchable

type _Option[T any] struct{}
type Option[T any] Matcher[_Option[T]]

type Some[T any] ValuedMatchable[_Option[T], T]
type None[T any] Matchable[_Option[T]]

func MatchOption[T, R any](
	option Option[T],
	some func(T) R,
	none func() R,
) R {
	switch v := option.(type) {
	case Some[T]:
		return some(v.Val)
	case None[T]:
		return none()
	}
	panic("Option type not supported")
}
