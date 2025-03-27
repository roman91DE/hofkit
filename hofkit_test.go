package hofkit

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	result := Map(func(x int) int {
		return x * 2
	}, []int{1, 2, 3})
	expected := []int{2, 4, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestMap failed: expected %v, got %v", expected, result)
	}
}

func TestFilter(t *testing.T) {
	result := Filter(func(x int) bool {
		return x%2 == 0
	}, []int{1, 2, 3, 4})
	expected := []int{2, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestFilter failed: expected %v, got %v", expected, result)
	}
}

func TestTakeWhile(t *testing.T) {
	result := TakeWhile(func(x int) bool {
		return x < 3
	}, []int{1, 2, 3, 4})
	expected := []int{1, 2}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TestTakeWhile failed: expected %v, got %v", expected, result)
	}
}

func TestPartitionBy(t *testing.T) {
	resultA, resultB := PartitionBy(func(x int) bool {
		return x%2 == 0
	}, []int{1, 2, 3, 4})
	expectedA := []int{2, 4}
	expectedB := []int{1, 3}
	if !reflect.DeepEqual(resultA, expectedA) {
		t.Errorf("TestPartitionBy A failed: expected %v, got %v", expectedA, resultA)
	}
	if !reflect.DeepEqual(resultB, expectedB) {
		t.Errorf("TestPartitionBy B failed: expected %v, got %v", expectedB, resultB)
	}
}

func TestReduce(t *testing.T) {
	result := Reduce(func(acc int, x int) int {
		return acc + x
	}, 0, []int{1, 2, 3})
	expected := 6
	if result != expected {
		t.Errorf("TestReduce failed: expected %v, got %v", expected, result)
	}
}

func TestAll(t *testing.T) {
	result := All(func(x int) bool {
		return x > 0
	}, []int{1, 2, 3})
	expected := true
	if result != expected {
		t.Errorf("TestAll failed: expected %v, got %v", expected, result)
	}
}

func TestAny(t *testing.T) {
	result := Any(func(x int) bool {
		return x%2 == 0
	}, []int{1, 3, 5})
	expected := false
	if result != expected {
		t.Errorf("TestAny failed: expected %v, got %v", expected, result)
	}
}

func TestFind(t *testing.T) {
	result, ok := Find(func(x int) bool {
		return x == 2
	}, []int{1, 2, 3})
	expected := 2

	if (result != expected) || (ok != true) {
		t.Errorf("TestFind failed: expected %v/%v, got %v/%v", expected, ok, result, ok)
	}
}

func TestFindIndex(t *testing.T) {
	result, ok := FindIndex(func(x int) bool {
		return x == 3
	}, []int{1, 2, 3})
	expected := 2
	if (result != expected) || (ok != true) {
		t.Errorf("TestFind failed: expected %v/%v, got %v/%v", expected, ok, result, ok)
	}
}

func TestForEach(t *testing.T) {
	var sum int
	ForEach(func(x int) {
		sum += x
	}, []int{1, 2, 3})
	expected := 6
	if sum != expected {
		t.Errorf("TestForEach failed: expected %v, got %v", expected, sum)
	}
}

func TestPartial1(t *testing.T) {
	add := func(x int, y int) int {
		return x + y
	}
	partialAdd := Partial1(add, 1)
	result := partialAdd(2)
	expected := 3
	if result != expected {
		t.Errorf("TestPartial1 failed: expected %v, got %v", expected, result)
	}
}

func TestPartial2(t *testing.T) {
	multiply := func(x int, y int, z int) int {
		return x * y * z
	}
	partialMultiply := Partial2(multiply, 2, 3)
	result := partialMultiply(4)
	expected := 24
	if result != expected {
		t.Errorf("TestPartial2 failed: expected %v, got %v", expected, result)
	}
}

func TestPartial3(t *testing.T) {
	concat := func(a string, b string, c string, d string) string {
		return a + b + c + d
	}
	partialConcat := Partial3(concat, "Hello", ", ", "World" )
	result := partialConcat("!")
	expected := "Hello, World!"
	if result != expected {
		t.Errorf("TestPartial3 failed: expected %v, got %v", expected, result)
	}
}
