package hofkit

import (
	"testing"
)

func BenchmarkMap(b *testing.B) {
	input := make([]int, 10000)
	for i := range input {
		input[i] = i
	}
	b.Run("Hofkit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Map(func(x int) int {
				return x * 2
			}, input)
		}
	})
	b.Run("Native", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			output := make([]int, len(input))
			for j, v := range input {
				output[j] = v * 2
			}
		}
	})
}

func BenchmarkFilter(b *testing.B) {
	input := make([]int, 10000)
	for i := range input {
		input[i] = i
	}
	b.Run("Hofkit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Filter(func(x int) bool {
				return x%2 == 0
			}, input)
		}
	})
	b.Run("Native", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			output := []int{}
			for _, v := range input {
				if v%2 == 0 {
					output = append(output, v)
				}
			}
		}
	})
}

func BenchmarkReduce(b *testing.B) {
	input := make([]int, 10000)
	for i := range input {
		input[i] = i
	}
	b.Run("Hofkit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Reduce(func(acc, x int) int {
				return acc + x
			}, 0, input)
		}
	})
	b.Run("Native", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			acc := 0
			for _, v := range input {
				acc += v
			}
		}
	})
}

func BenchmarkAll(b *testing.B) {
	input := make([]int, 10000)
	for i := range input {
		input[i] = i
	}
	b.Run("Hofkit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			All(func(x int) bool {
				return x < 10000
			}, input)
		}
	})
	b.Run("Native", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, v := range input {
				if v >= 10000 {
					break
				}
			}
		}
	})
}

func BenchmarkAny(b *testing.B) {
	input := make([]int, 10000)
	for i := range input {
		input[i] = i
	}
	b.Run("Hofkit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Any(func(x int) bool {
				return x%2 == 0
			}, input)
		}
	})
	b.Run("Native", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, v := range input {
				if v%2 == 0 {
					break
				}
			}
		}
	})
}

func BenchmarkFind(b *testing.B) {
	input := make([]int, 10000)
	for i := range input {
		input[i] = i
	}
	b.Run("Hofkit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Find(func(x int) bool {
				return x == 5000
			}, input)
		}
	})
	b.Run("Native", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, v := range input {
				if v == 5000 {
					break
				}
			}
		}
	})
}

func BenchmarkFindIndex(b *testing.B) {
	input := make([]int, 10000)
	for i := range input {
		input[i] = i
	}
	b.Run("Hofkit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FindIndex(func(x int) bool {
				return x == 5000
			}, input)
		}
	})
	b.Run("Native", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, v := range input {
				if v == 5000 {
					break
				}
			}
		}
	})
}

func BenchmarkTakeWhile(b *testing.B) {
	input := make([]int, 10000)
	for i := range input {
		input[i] = i
	}
	b.Run("Hofkit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			TakeWhile(func(x int) bool {
				return x < 5000
			}, input)
		}
	})
	b.Run("Native", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var output []int
			for _, v := range input {
				if v < 5000 {
					output = append(output, v)
				} else {
					break
				}
			}
		}
	})
}

func BenchmarkForEach(b *testing.B) {
	input := make([]int, 10000)
	for i := range input {
		input[i] = i
	}
	b.Run("Hofkit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ForEach(func(x int) {
				_ = x * 2
			}, input)
		}
	})
	b.Run("Native", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, v := range input {
				_ = v * 2
			}
		}
	})
}

func BenchmarkPartitionBy(b *testing.B) {
	input := make([]int, 10000)
	for i := range input {
		input[i] = i
	}
	b.Run("Hofkit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			PartitionBy(func(x int) bool {
				return x%2 == 0
			}, input)
		}
	})
	b.Run("Native", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var a, b []int
			for _, v := range input {
				if v%2 == 0 {
					a = append(a, v)
				} else {
					b = append(b, v)
				}
			}
		}
	})
}
