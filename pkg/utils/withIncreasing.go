package utils

func WithIncreasing[T any, R any](f func(T) R) (*int, func(T) R) {
	num := 1

	return &num, func(arg T) R {
		num++
		return f(arg)
	}
}
