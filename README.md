Our goal is to examine the internal logger log/slog in golang and measure its performance with other loggers. And also give info about log/slog's features.


I used 5 different logger libraries in benchmark tests and I did 2 different tests with struct and without struct.

With Struct Benchmark Result

| Test Name                   |     Time     | Bytes Per Allocation | Objects Allocated |
|-----------------------------|:------------:|---------------------:|------------------:|
| BenchmarkApex-16            | 652365 ns/op |           54391 B/op |     343 allocs/op |
| BenchmarkLogrus-16          | 507813 ns/op |           60711 B/op |     318 allocs/op |
| BenchmarkSlog-16            | 656784 ns/op |           51827 B/op |     420 allocs/op |
| BenchmarkSlogTextHandler-16 | 681030 ns/op |           51776 B/op |     420 allocs/op |
| BenchmarkSlogJsonHandler-16 | 482934 ns/op |           46966 B/op |      60 allocs/op |
| BenchmarkZapSugar-16        | 329518 ns/op |           49453 B/op |     249 allocs/op |

---

Without Struct Benchmark Results

| Test Name                                |     Time     | Bytes Per Allocation | Objects Allocated |
|------------------------------------------|:------------:|---------------------:|------------------:|
| BenchmarkApexWithoutStruct-16            | 31461 ns/op  |             720 B/op |      15 allocs/op |
| BenchmarkLogrusWithoutStruct-16          | 513523 ns/op |           60685 B/op |     318 allocs/op |
| BenchmarkSlogWithoutStruct-16            | 24898 ns/op  |               0 B/op |       0 allocs/op |
| BenchmarkSlogTextHandlerWithoutStruct-16 | 28122 ns/op  |               0 B/op |       0 allocs/op |
| BenchmarkSlogJsonHandlerWithoutStruct-16 | 32574 ns/op  |               0 B/op |       0 allocs/op |
| BenchmarkZapSugarWithoutStruct-16        | 942.2 ns/op  |              11 B/op |       0 allocs/op |
| BenchmarkZapWithoutStruct-16             | 926.2 ns/op  |              11 B/op |       0 allocs/op |
| BenchmarkZerologWithoutStruct-16         | 29359 ns/op  |               0 B/op |       0 allocs/op |



# **log/slog features**

## Handler Options

log/slog have 2 different handler to  formatting the logs. Also you can write your custom handler.

### JSONHandler

```
logger := slog.New(slog.NewJSONHandler(os.Stdout))

    logger.Debug("This is a Debug message")
    logger.Info("This is an Info message")
    logger.Warn("This is a Warning message")
    logger.Error("This is an Error message")
```

Log

```
{"time":"2023-09-02T10:06:23.682743+03:00","level":"INFO","msg":"This is an Info message"}
{"time":"2023-09-02T10:06:23.682747+03:00","level":"WARN","msg":"This is a Warning message"}
{"time":"2023-09-02T10:06:23.68275+03:00","level":"ERROR","msg":"This is an Error message"}
```

### TextHandler

```
logger := slog.New(slog.NewTextHandler(os.Stdout))

    logger.Debug("This is a Debug message")
    logger.Info("This is an Info message")
    logger.Warn("This is a Warning message")
    logger.Error("This is an Error message")
```

Log

```
time=2023-09-02T10:06:19.906+03:00 level=INFO msg="This is an Info message"
time=2023-09-02T10:06:19.906+03:00 level=WARN msg="This is a Warning message"
time=2023-09-02T10:06:19.906+03:00 level=ERROR msg="This is an Error message"
```


## Add Custom Field

You can add custom fields to your log.

```
func LoggerSlogWithCustomField(jsonLogger *slog.Logger) {
	jsonLogger.Info("hello", "test", 1)
	jsonLogger.Info("This is an Info message", slog.Int("test", 1))
}
```

Log

```
{"time":"2023-09-02T10:09:25.279457+03:00","level":"INFO","msg":"hello","test":1}
{"time":"2023-09-02T10:09:25.280265+03:00","level":"INFO","msg":"This is an Info message","test":1}
```


## Child Logger

If you want to add new fields to your log you can use Logger.With

```
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
```

Log

```
{"time":"2023-09-02T10:09:25.280315+03:00","level":"INFO","msg":"image upload successful","program_info":{"pid":92036,"go_version":"go1.21rc4"},"image_id":"39ud88"}
{"time":"2023-09-02T10:09:25.28032+03:00","level":"WARN","msg":"storage is 90% full","program_info":{"pid":92036,"go_version":"go1.21rc4"},"available_space":"900.1 mb"}

{"time":"2023-09-02T10:09:25.280315+03:00","level":"INFO","msg":"image upload successful","program_info":{"pid":92036,"go_version":"go1.21rc4"},"image_id":"39ud88"}
{"time":"2023-09-02T10:09:25.28032+03:00","level":"WARN","msg":"storage is 90% full","program_info":{"pid":92036,"go_version":"go1.21rc4"},"available_space":"900.1 mb"}
```

## Hiding Sensitive Data

You can use interface for the hiding sensitive data 

```
type LogValuer interface {
LogValue() Value
}
```

Example 
```
type User struct {
ID        string `json:"id"`
FirstName string `json:"first_name"`
LastName  string `json:"last_name"`
Email     string `json:"email"`
Password  string `json:"password"`
}

func (u *User) LogValue() slog.Value {
    return slog.StringValue(u.ID)
}

func HidingSensitiveData() {
    handler := slog.NewJSONHandler(os.Stdout, nil)
    logger := slog.New(handler)

	u := &User{
		ID:        "user-12234",
		FirstName: "Jan",
		LastName:  "Doe",
		Email:     "jan@example.com",
		Password:  "pass-12334",
	}

	logger.Info("info", "user", u)
}
```

Log

```
{"time":"2023-09-14T10:04:51.272601+03:00","level":"INFO","msg":"info","user":"user-12234"}
```

## Use slog as a frontend

You can use slog as a frontend and use external third-party loggers as a backend.


```
func UseSlogFrontendZapBackend() {
zapLogger := zap.Must(zap.NewProduction())

	defer zapLogger.Sync()

	logger := slog.New(zapslog.NewHandler(zapLogger.Core(), nil))

	logger.Info(
		"Using Slog frontend with Zap backend!",
		slog.Int("process_id", os.Getpid()),
	)
}
```

Log

```
{"level":"info","ts":1694675091.2726688,"msg":"Using Slog frontend with Zap backend!","process_id":71993}
```



## Default Logger

Also you can use slog to default logger.
```
func SetSlogDefaultLogger(jsonLogger *slog.Logger) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	log.Println("Hello from old logger")
}
```

Log

```
{"time":"2023-09-02T10:09:25.280487+03:00","level":"INFO","msg":"Hello from old logger"}
```

### Pros

- a more performant handler can be written because it has its own interface.

- log/slog golang's internal library

- In benchmarks, log/slog is more performant than apex and logrus libraries

- SlogJsonHandler's number of allocations per operation value in tests with struct is much lower than other libraries

Cons

- Extremely slow compared to Zap in benchmarks.

Sources

- https://betterstack.com/community/guides/logging/logging-in-go/#creating-and-using-child-loggers

- https://www.highlight.io/blog/5-best-logging-libraries-for-go