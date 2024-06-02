# GoDat
Go items.dat Encoder/Decoder for Growtopia. Extremely fast btw.

# Features
- [X] ProtonHash Calculator.
- [X] Items.dat Encoder & Decoder.
- [X] Simple UI & User-Friendly.
- [X] Very Fast Processing.

# Benchmark
### Encoder (100 Process)
```go
goos: android
goarch: arm64
pkg: github.com/eikarna/GoDat
BenchmarkEncoder-8             6         235593108 ns/op        99593234 B/op   7561045 allocs/op
--- BENCH: BenchmarkEncoder-8
    main_test.go:64: Success Count: 10
    main_test.go:65: Error Count: 0
    main_test.go:67: Average Encode Time: 228.13302ms
    main_test.go:64: Success Count: 30
    main_test.go:65: Error Count: 0
    main_test.go:67: Average Encode Time: 282.917076ms
    main_test.go:64: Success Count: 40
    main_test.go:65: Error Count: 0
    main_test.go:67: Average Encode Time: 163.167127ms
    main_test.go:64: Success Count: 50
        ... [output truncated]
PASS
ok      github.com/eikarna/GoDat        11.689s
```
### Decoder (100 Process) / Not Tested
```go
nil
```
You can test it by yourself if your device is medium/high-end, also minimal RAM is 6GB for prevent the benchmark test force close, because this program use memory-map logic where all decoded/parsed Binary file/JSON file is mapped to memory for ensure the process is fast and efficient.

## Special Thanks
- [GuckTubeYT](https://github.com/GuckTubeYT)
- [KipasGTS](https://github.com/KipasGTS)
