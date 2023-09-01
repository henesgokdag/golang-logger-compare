package util

import "testing"

func BenchmarkZapSugar(b *testing.B) {
	zapLogger := getZapSugarLogger()
	for i := 0; i < b.N; i++ {
		loggerZapSugar(zapLogger)
	}
}

func BenchmarkZapSugarWithoutStruct(b *testing.B) {
	zapLogger := getZapSugarLogger()
	for i := 0; i < b.N; i++ {
		loggerZapSugarWithoutStruct(zapLogger)
	}
}

func BenchmarkZapWithoutStruct(b *testing.B) {
	zapLogger := getZapSugarLogger()
	for i := 0; i < b.N; i++ {
		loggerZapwWithoutStruct(zapLogger)
	}
}
