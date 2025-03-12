package main

import (
	"fmt"
	"testing"
)

// Cộng chuỗi bằng +
func ConcatPlus(a, b string) string {
	return a + b
}

// Cộng chuỗi bằng fmt.Sprintf
func ConcatFmt(a, b string) string {
	return fmt.Sprintf("%s%s", a, b)
}

func BenchmarkConcatPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatPlus("Hello", "World")
	}
}

func BenchmarkConcatFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatFmt("Hello", "World")
	}
}
