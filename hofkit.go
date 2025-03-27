package hofkit

// Map applies function f to each element of the slice xs and returns a new slice with the results;
func Map[F any, T any](f func(F) T, xs []F) []T {
	mapped := make([]T, len(xs))
	for i, it := range xs {
		mapped[i] = f(it)
	}
	return mapped
}

// Filter applies predicate function p to each element of slice xs and returns a new slice with the elements that satisfy p;
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

// Reduce applies reducer function f to elements of slice xs, accumulating the result with initial value init, and returns the final result;
func Reduce[F any, T any](f func(F, T) T, init T, xs []F) T {
	acc := init
	for _, it := range xs {
		acc = f(it, acc)
	}
	return acc
}

// All applies predicate function p to each element of slice xs and returns true if all elements satisfy p;
func All[T any](p func(T) bool, xs []T) bool {
	for _, it := range xs {
		if !p(it) {
			return false
		}
	}
	return true
}

// Any applies predicate function p to each element of slice xs and returns true if any element satisfies p;
func Any[T any](p func(T) bool, xs []T) bool {
	for _, it := range xs {
		if p(it) {
			return true
		}
	}
	return false
}

// Find applies predicate function p to each element of slice xs and returns the first element that satisfies p and true; if no match is found, it returns the zero value and false;
func Find[T any](p func(T) bool, xs []T) (T, bool) {
	var zero T
	for _, it := range xs {
		if p(it) {
			return it, true
		}
	}
	return zero, false
}

// FindIndex applies predicate function p to each element of slice xs and returns the index of the first match and true; if no match is found, it returns -1 and false;
func FindIndex[T any](p func(T) bool, xs []T) (int, bool) {
	for idx, it := range xs {
		if p(it) {
			return idx, true
		}
	}
	return -1, false
}

// TakeWhile applies predicate function p to elements of slice xs and returns a new slice of leading elements that satisfy p;
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

// ForEach applies function f to each element of slice xs for side effects;
func ForEach[T any](f func(T), xs []T) {
	for _, it := range xs {
		f(it)
	}
}

// PartitionBy applies predicate function p to elements of slice xs and returns two slices: one with elements that satisfy p, and one with the rest;
func PartitionBy[T any](p func(T) bool, xs []T) ([]T, []T) {
	var a, b []T
	for _, it := range xs {
		if p(it) {
			a = append(a, it)
		} else {
			b = append(b, it)
		}
	}
	return a, b
}

// Partial1 takes a function (A, B) R and a value A, and returns func(B) R;
func Partial1[A any, B any, R any](f func(A, B) R, a A) func(B) R {
	return func(b B) R {
		return f(a, b)
	}
}

// Partial2 takes a function func(A, B, C) R and values A, B, and returns func(C) R;
func Partial2[A any, B any, C any, R any](f func(A, B, C) R, a A, b B) func(C) R {
	return func(c C) R {
		return f(a, b, c)
	}
}

// Partial3 takes a function func(A, B, C, D) R and values A, B, C, and returns func(D) R;
func Partial3[A any, B any, C any, D any, R any](f func(A, B, C, D) R, a A, b B, c C) func(D) R {
	return func(d D) R {
		return f(a, b, c, d)
	}
}
