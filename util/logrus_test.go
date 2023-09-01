package util

import "testing"

func BenchmarkLogrus(b *testing.B) {
	setFormatter()
	for i := 0; i < b.N; i++ {
		loggerLogrus()
	}
}

func BenchmarkLogrusWithoutStruct(b *testing.B) {
	setFormatter()
	for i := 0; i < b.N; i++ {
		loggerLogrus()
	}
}
