package result

import "reflect"

type some[T any] struct {
	value T
}

func Some[T any](res T) Option[T] {
	return some[T]{res}
}

func (s some[T]) Ok() bool {
	return true
}

func (s some[T]) Get() T {
	return s.value
}

func (s some[T]) OkOr(fallback T) T {
	return s.value
}

func (s some[T]) Equals(option Option[T]) bool {
	return s.Ok() && option.Ok() && (reflect.DeepEqual(s.value, option.Get()) || reflect.ValueOf(s.value) == reflect.ValueOf(option.Get()))
}
