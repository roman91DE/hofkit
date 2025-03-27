# hofkit

`hofkit` is a lightweight and idiomatic Go library providing functional programming utilities such as `Map`, `Filter`, `Reduce`, and more. It is designed to bring expressive, higher-order abstractions to Go while remaining performant and simple.

## ✨ Features

- Generic functions using Go 1.18+ type parameters
- Minimal allocations and performance close to native loops
- Fully benchmarked and tested

## 📦 Installation

```bash
go get github.com/roman91DE/hofkit
```

## 🧠 Function Reference

Each function is generic and works with any type.

### `Map`

```go
func Map[F any, T any](f func(F) T, xs []F) []T
```
Map applies function `f` to each element of the slice `xs` and returns a new slice with the results;

---

### `Filter`

```go
func Filter[T any](p func(T) bool, xs []T) []T
```
Filter applies predicate function `p` to each element of slice `xs` and returns a new slice with the elements that satisfy `p`;

---

### `Reduce`

```go
func Reduce[F any, T any](f func(F, T) T, init T, xs []F) T
```
Reduce applies reducer function `f` to elements of slice `xs`, accumulating the result with initial value `init`, and returns the final result;

---

### `All`

```go
func All[T any](p func(T) bool, xs []T) bool
```
All applies predicate function `p` to each element of slice `xs` and returns true if all elements satisfy `p`;

---

### `Any`

```go
func Any[T any](p func(T) bool, xs []T) bool
```
Any applies predicate function `p` to each element of slice `xs` and returns true if any element satisfies `p`;

---

### `Find`

```go
func Find[T any](p func(T) bool, xs []T) (T, bool)
```
Find applies predicate function `p` to each element of slice `xs` and returns the first element that satisfies `p` and true; if no match is found, it returns the zero value and false;

---

### `FindIndex`

```go
func FindIndex[T any](p func(T) bool, xs []T) (int, bool)
```
FindIndex applies predicate function `p` to each element of slice `xs` and returns the index of the first match and true; if no match is found, it returns -1 and false;

---

### `TakeWhile`

```go
func TakeWhile[T any](p func(T) bool, xs []T) []T
```
TakeWhile applies predicate function `p` to elements of slice `xs` and returns a new slice of leading elements that satisfy `p`;

---

### `ForEach`

```go
func ForEach[T any](f func(T), xs []T)
```
ForEach applies function `f` to each element of slice `xs` for side effects;

---

### `PartitionBy`

```go
func PartitionBy[T any](p func(T) bool, xs []T) ([]T, []T)
```
PartitionBy applies predicate function `p` to elements of slice `xs` and returns two slices: one with elements that satisfy `p`, and one with the rest;

---

### `Partial1`

```go
func Partial1[A any, B any, R any](f func(A, B) R, a A) func(B) R
```
Partial1 takes a function `(A, B) R` and a value `A`, and returns `func(B) R`;

---

### `Partial2`

```go
func Partial2[A any, B any, C any, R any](f func(A, B, C) R, a A, b B) func(C) R
```
Partial2 takes a function `func(A, B, C) R` and values `A, B`, and returns `func(C) R`;

---

### `Partial3`

```go
func Partial3[A any, B any, C any, D any, R any](f func(A, B, C, D) R, a A, b B, c C) func(D) R
```
Partial3 takes a function `func(A, B, C, D) R` and values `A, B, C`, and returns `func(D) R`;

---

## 🚀 Benchmarks

Benchmarks are located in the [`benchmark/`](./benchmark/) folder. They compare each `hofkit` function against an idiomatic native Go equivalent using:

```bash
go test -bench=. -benchmem
```

Plots and performance reports are generated with the Python script [`plot_bench.py`](./plot_bench.py).

Each benchmark is run on 10,000-element slices with non-trivial logic to stress CPU and memory.

Results using a M1 2020 MacBookPro:

[Memory Allocations](./benchmark/darwin_arm64_arm_20250327_123742/benchmark_allocs_op.png)
[Memory Consumption](./benchmark/darwin_arm64_arm_20250327_123742/benchmark_B_op.png)
[Time Consumption](./benchmark/darwin_arm64_arm_20250327_123742/benchmark_ns_op.png)


---

## 🧪 Testing

Run all tests:

```bash
go test ./...
```


---

## 🛠️ Contributing

Pull requests welcome! This library is still young — feel free to open issues or suggest improvements.

