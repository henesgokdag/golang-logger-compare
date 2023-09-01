package util

import "testing"

func BenchmarkZerologWithoutStruct(b *testing.B) {
	setZerologLogLevel()
	for i := 0; i < b.N; i++ {
		loggerZerologWithoutStruct()
	}
}
