package result

import "reflect"

func IsFalsy[T any](value T) bool {
	return reflect.DeepEqual(value, "") || reflect.DeepEqual(value, 0) || reflect.DeepEqual(value, false) || reflect.DeepEqual(value, nil)
}

func IsTruthy[T any](value T) bool {
	return !IsFalsy(value)
}

func ValueOr[T any](value T, defaultValue T) T {
	if IsFalsy(value) {
		return defaultValue
	}

	return value
}

func ValueOrElse[T any](value T, fallback func() T) T {
	if IsFalsy(value) {
		return fallback()
	}

	return value
}
