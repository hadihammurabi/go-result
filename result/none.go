package result

type none[T any] struct {
	value T
}

func None[T any]() Option[T] {
	return none[T]{}
}

func (s none[T]) Ok() bool {
	return false
}

func (s none[T]) Get() T {
	return s.value
}

func (s none[T]) OkOr(fallback T) T {
	return fallback
}

func (s none[T]) Equals(option Option[T]) bool {
	return !option.Ok()
}
