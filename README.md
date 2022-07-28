# Matchable

A union and pattern matching package for Go

## Examples

### Angle

```go
type _Angle struct{}
type Angle Matcher[_Angle]
type Deg ValuedMatchable[_Angle, float32]
type Rad ValuedMatchable[_Angle, float32]

func MatchAngle[R any](
	angle Angle,
	deg func(*Deg) R,
	rad func(*Rad) R,
) R {
	switch v := angle.(type) {
	case Deg:
		return deg(&v)
	case Rad:
		return rad(&v)
	}
	panic("MatchAngle: angle type not supported")
}
```

### Option[T]

```go
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
```

### Result[T, E]

```go
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
```
