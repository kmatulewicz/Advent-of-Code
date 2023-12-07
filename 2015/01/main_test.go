package main

import "testing"

func BenchmarkCountFromString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countFromString("input")
	}
}

func BenchmarkCountByteByByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countByteByByte("input")
	}
}
