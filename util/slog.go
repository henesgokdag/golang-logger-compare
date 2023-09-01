package util

import (
	"context"
	"log"
	"log/slog"
	"os"
	"runtime/debug"
)

func getTextLogger() *slog.Logger {
	var programLevel = new(slog.LevelVar)
	programLevel.Set(slog.LevelInfo)
	return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel}))
}

func GetJsonLogger() *slog.Logger {
	var programLevel = new(slog.LevelVar)
	programLevel.Set(slog.LevelInfo)
	return slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel}))
}

func loggerSlog() {
	slog.Debug("This is a Debug message", "data", getExampleData())
	slog.Info("This is an Info message", "data", getExampleData())
	slog.Warn("This is a Warning message", "data", getExampleData())
	slog.Error("This is an Error message", "data", getExampleData())
}

func loggerSlogWithoutStruct() {
	slog.Debug("This is a Debug message")
	slog.Info("This is an Info message")
	slog.Warn("This is a Warning message")
	slog.Error("This is an Error message")
}

func loggerSlogTextHandler(textLogger *slog.Logger) {
	textLogger.Debug("This is a Debug message", "data", getExampleData())
	textLogger.Info("This is an Info message", "data", getExampleData())
	textLogger.Warn("This is a Warning message", "data", getExampleData())
	textLogger.Error("This is an Error message", "data", getExampleData())
}

func loggerSlogJsonHandler(jsonLogger *slog.Logger) {
	jsonLogger.Debug("This is a Debug message", "data", getExampleData())
	jsonLogger.Info("This is an Info message", "data", getExampleData())
	jsonLogger.Warn("This is a Warning message", "data", getExampleData())
	jsonLogger.Error("This is an Error message", "data", getExampleData())
}

func loggerSlogTextHandlerWithoutStruct(textLogger *slog.Logger) {
	textLogger.Debug("This is a Debug message")
	textLogger.Info("This is an Info message")
	textLogger.Warn("This is a Warning message")
	textLogger.Error("This is an Error message")
}

func loggerSlogJsonHandlerWithoutStruct(jsonLogger *slog.Logger) {
	jsonLogger.Debug("This is a Debug message")
	jsonLogger.Info("This is an Info message")
	jsonLogger.Warn("This is a Warning message")
	jsonLogger.Error("This is an Error message")
}

func LoggerSlogWithCustomField(jsonLogger *slog.Logger) {
	jsonLogger.Info("hello", "test", 1)
	jsonLogger.Info("This is an Info message", slog.Int("test", 1))
}

func SetSlogDefaultLogger(jsonLogger *slog.Logger) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	log.Println("Hello from old logger")
}

func SetChildLogger(jsonLogger *slog.Logger) {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	buildInfo, _ := debug.ReadBuildInfo()

	logger := slog.New(handler)

	child := logger.With(
		slog.Group("program_info",
			slog.Int("pid", os.Getpid()),
			slog.String("go_version", buildInfo.GoVersion),
		),
	)
	child.Info("image upload successful", slog.String("image_id", "39ud88"))
	child.Warn(
		"storage is 90% full",
		slog.String("available_space", "900.1 mb"),
	)
}

const (
	LevelTrace  = slog.Level(-8)
	LevelNotice = slog.Level(2)
	LevelFatal  = slog.Level(12)
)

func LoggerWithCustomLevelsAndSource() {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     LevelTrace,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	ctx := context.Background()
	logger.Log(ctx, LevelTrace, "Trace message")
	logger.Log(ctx, LevelNotice, "Notice message")
	logger.Log(ctx, LevelFatal, "Fatal level")
}
