package hofkit

func Map[F any, T any](f func(F) T, xs []F) []T {
	mapped := make([]T, len(xs))
	for i, it := range xs {
		mapped[i] = f(it)
	}
	return mapped
}

func Filter[T any](p func(T) bool, xs []T) []T {
	filtered := make([]T, len(xs))
	count := 0

	for _, it := range xs {
		if p(it) {
			filtered[count] = it
			count += 1
		}
	}
	return filtered[:count]
}

func Reduce[F any, T any](f func(F, T) T, init T, xs []F) T {
	acc := init
	for _, it := range xs {
		acc = f(it, acc)
	}
	return acc
}

func All[T any](p func(T) bool, xs []T) bool {
	for _, it := range xs {
		if !p(it) {
			return false
		}
	}
	return true
}

func Any[T any](p func(T) bool, xs []T) bool {
	for _, it := range xs {
		if p(it) {
			return true
		}
	}
	return false
}

func Find[T any](p func(T) bool, xs []T) (T, bool) {
	var zero T
	for _, it := range xs {
		if p(it) {
			return it, true
		}
	}
	return zero, false
}

func FindIndex[T any](p func(T) bool, xs []T) (int, bool) {
	for idx, it := range xs {
		if p(it) {
			return idx, true
		}
	}
	return -1, false
}

func TakeWhile[T any](p func(T) bool, xs []T) []T {
	var result []T
	for _, it := range xs {
		if !p(it) {
			break
		}
		result = append(result, it)
	}
	return result
}

func ForEach[T any](f func(T), xs []T) {
	for _, it := range xs {
		f(it)
	}
}