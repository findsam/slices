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