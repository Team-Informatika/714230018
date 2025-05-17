package main

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Errorf("Add(2, 3) = %d; want 5", got)
	}
}

func TestSubtract(t *testing.T) {
	if got := Subtract(5, 3); got != 2 {
		t.Errorf("Subtract(5, 3) = %d; want 2", got)
	}
}

func TestMultiply(t *testing.T) {
	if got := Multiply(4, 3); got != 12 {
		t.Errorf("Multiply(4, 3) = %d; want 12", got)
	}
}

func TestDivide(t *testing.T) {
	quotient, remainder, err := Divide(10, 3)
	if err != nil {
		t.Errorf("Divide(10, 3) unexpected error: %v", err)
	}
	if quotient != 3 || remainder != 1 {
		t.Errorf("Divide(10, 3) = (%d, %d); want (3, 1)", quotient, remainder)
	}

	_, _, err = Divide(10, 0)
	if err == nil {
		t.Error("Divide(10, 0) expected error, got nil")
	}
}

func TestSwap(t *testing.T) {
	a, b := Swap("hello", "world")
	if a != "world" || b != "hello" {
		t.Errorf("Swap(\"hello\", \"world\") = (%s, %s); want (\"world\", \"hello\")", a, b)
	}
}

func TestMin(t *testing.T) {
	if got := Min(3, 5); got != 3 {
		t.Errorf("Min(3, 5) = %d; want 3", got)
	}
	if got := Min(10, 2); got != 2 {
		t.Errorf("Min(10, 2) = %d; want 2", got)
	}
}

func TestMaxFunction(t *testing.T) {
	if got := Max(3, 5); got != 5 {
		t.Errorf("Max(3, 5) = %d; want 5", got)
	}
	if got := Max(10, 2); got != 10 {
		t.Errorf("Max(10, 2) = %d; want 10", got)
	}
}

func TestIsEven(t *testing.T) {
	if !IsEven(4) {
		t.Error("IsEven(4) = false; want true")
	}
	if IsEven(3) {
		t.Error("IsEven(3) = true; want false")
	}
}
func TestIsOdd(t *testing.T) {
	if !IsOdd(3) {
		t.Error("IsOdd(3) = false; want true")
	}
	if IsOdd(4) {
		t.Error("IsOdd(4) = true; want false")
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n       int
		want    int
		wantErr bool
	}{
		{0, 0, false},
		{1, 1, false},
		{7, 13, false},
		{-1, -1, true},
	}
	for _, tt := range tests {
		got, err := Fibonacci(tt.n)
		if (err != nil) != tt.wantErr {
			t.Errorf("Fibonacci(%d) error = %v, wantErr %v", tt.n, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("Fibonacci(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}

func TestReverseString(t *testing.T) {
	if got := ReverseString("hello"); got != "olleh" {
		t.Errorf("ReverseString(\"hello\") = %s; want \"olleh\"", got)
	}
}

func TestCountWords(t *testing.T) {
	if got := CountWords("hello world from go"); got != 4 {
		t.Errorf("CountWords(\"hello world from go\") = %d; want 4", got)
	}
}

func TestContains(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	if !Contains(slice, 3) {
		t.Error("Contains(slice, 3) = false; want true")
	}
	if Contains(slice, 5) {
		t.Error("Contains(slice, 5) = true; want false")
	}
}

func TestSumSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	if got := SumSlice(slice); got != 10 {
		t.Errorf("SumSlice(slice) = %d; want 10", got)
	}
}

func TestFilterEven(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	want := []int{2, 4, 6}
	got := FilterEven(slice)
	if len(got) != len(want) {
		t.Errorf("FilterEven(slice) length = %d; want %d", len(got), len(want))
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("FilterEven(slice)[%d] = %d; want %d", i, got[i], want[i])
		}
	}
}

func TestMapStringLengths(t *testing.T) {
	slice := []string{"go", "lang"}
	want := map[string]int{"go": 2, "lang": 4}
	got := MapStringLengths(slice)
	if len(got) != len(want) {
		t.Errorf("MapStringLengths(slice) length = %d; want %d", len(got), len(want))
	}
	for k, v := range want {
		if got[k] != v {
			t.Errorf("MapStringLengths(slice)[%s] = %d; want %d", k, got[k], v)
		}
	}
}

func TestPersonMethods(t *testing.T) {
	p := NewPerson("Alice", 30)
	if p.Name != "Alice" || p.Age != 30 {
		t.Errorf("NewPerson(\"Alice\", 30) = %+v; want Name=Alice Age=30", p)
	}
	p.Birthday()
	if p.Age != 31 {
		t.Errorf("After Birthday, Age = %d; want 31", p.Age)
	}
	wantStr := "Alice is 31 years old"
	if p.String() != wantStr {
		t.Errorf("Person.String() = %s; want %s", p.String(), wantStr)
	}
}

func TestRectangle(t *testing.T) {
	r := Rectangle{Width: 3, Height: 4}
	if got := r.Area(); got != 12 {
		t.Errorf("Rectangle.Area() = %f; want 12", got)
	}
	if got := r.Perimeter(); got != 14 {
		t.Errorf("Rectangle.Perimeter() = %f; want 14", got)
	}
}

func TestCircle(t *testing.T) {
	c := Circle{Radius: 5}
	if got := c.Area(); math.Abs(got-78.5398) > 0.0001 {
		t.Errorf("Circle.Area() = %f; want approx 78.5398", got)
	}
	if got := c.Perimeter(); math.Abs(got-31.4159) > 0.0001 {
		t.Errorf("Circle.Perimeter() = %f; want approx 31.4159", got)
	}
}

func TestSafeDivide(t *testing.T) {
	res, err := SafeDivide(10, 2)
	if err != nil {
		t.Errorf("SafeDivide(10, 2) unexpected error: %v", err)
	}
	if res != 5 {
		t.Errorf("SafeDivide(10, 2) = %f; want 5", res)
	}

	_, err = SafeDivide(10, 0)
	if err == nil {
		t.Error("SafeDivide(10, 0) expected error, got nil")
	}
}

func TestPower(t *testing.T) {
	if got := Power(2, 3); got != 8 {
		t.Errorf("Power(2, 3) = %d; want 8", got)
	}
	if got := Power(5, 0); got != 1 {
		t.Errorf("Power(5, 0) = %d; want 1", got)
	}
}

func TestSumVariadic(t *testing.T) {
	if got := Sum(1, 2, 3, 4); got != 10 {
		t.Errorf("Sum(1, 2, 3, 4) = %d; want 10", got)
	}
	if got := Sum(); got != 0 {
		t.Errorf("Sum() = %d; want 0", got)
	}
}

func TestAdder(t *testing.T) {
	adder := Adder(5)
	if got := adder(10); got != 15 {
		t.Errorf("Adder(5)(10) = %d; want 15", got)
	}
}

func TestFilter(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	predicate := func(n int) bool { return n > 3 }
	want := []int{4, 5}
	got := Filter(slice, predicate)
	if len(got) != len(want) {
		t.Errorf("Filter(slice, predicate) length = %d; want %d", len(got), len(want))
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("Filter(slice, predicate)[%d] = %d; want %d", i, got[i], want[i])
		}
	}
}
