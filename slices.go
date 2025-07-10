package slices

func Map[T any, R any](s []T, fn func(T) R) []R {
	result := make([]R, len(s))
	for i, v := range s {
		result[i] = fn(v)
	}
	return result
}

func Reduce[T any, R any](s []T, init R, fn func(R, T) R) R {
	acc := init
	for _, v := range s {
		acc = fn(acc, v)
	}
	return acc
}

func Filter[T any](s []T, fn func(T) bool) []T {
	result := make([]T, 0, len(s))
	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}
