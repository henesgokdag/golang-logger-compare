package util

import "testing"

func BenchmarkApex(b *testing.B) {
	setHandler()
	for i := 0; i < b.N; i++ {
		loggerApex()
	}
}

func BenchmarkApexWithoutStruct(b *testing.B) {
	setHandler()
	for i := 0; i < b.N; i++ {
		loggerApexWithoutStruct()
	}
}
