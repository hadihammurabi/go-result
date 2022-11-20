package result

type Option[T any] interface {
	Ok() bool
	Get() T
	OkOr(T) T
	Equals(Option[T]) bool
}
