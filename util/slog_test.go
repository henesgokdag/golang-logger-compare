package util

import "testing"

func BenchmarkSlog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loggerSlog()
	}
}

func BenchmarkSlogWithoutStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loggerSlogWithoutStruct()
	}
}

func BenchmarkSlogTextHandler(b *testing.B) {
	slogTextLogger := getTextLogger()
	for i := 0; i < b.N; i++ {
		loggerSlogJsonHandler(slogTextLogger)
	}
}

func BenchmarkSlogTextHandlerWithoutStruct(b *testing.B) {
	slogTextLogger := getTextLogger()
	for i := 0; i < b.N; i++ {
		loggerSlogTextHandlerWithoutStruct(slogTextLogger)
	}
}

func BenchmarkSlogJsonHandler(b *testing.B) {
	slogJsonLogger := GetJsonLogger()
	for i := 0; i < b.N; i++ {
		loggerSlogJsonHandler(slogJsonLogger)
	}
}

func BenchmarkSlogJsonHandlerWithoutStruct(b *testing.B) {
	slogJsonLogger := GetJsonLogger()
	for i := 0; i < b.N; i++ {
		loggerSlogJsonHandlerWithoutStruct(slogJsonLogger)
	}
}
