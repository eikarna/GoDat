# GoDat
Go items.dat Encoder/Decoder for Growtopia. Extremely fast btw.

# Features
- [X] ProtonHash Calculator.
- [X] Items.dat Encoder & Decoder.
- [X] Simple UI & User-Friendly.
- [X] Very Fast Processing.

# Benchmark
### Encoder (8 Cores, 4x2.4GHz, 4x1.8GHz)
```go
goos: android
goarch: arm64
pkg: github.com/eikarna/GoDat
BenchmarkEncoder-8            15          69861531 ns/op         9959195 B/op    756105 allocs/op
--- BENCH: BenchmarkEncoder-8
    main_test.go:64: Success Count: 1
    main_test.go:65: Error Count: 0
    main_test.go:67: Average Encode Time: 84.600834ms
    main_test.go:64: Success Count: 13
    main_test.go:65: Error Count: 0
    main_test.go:67: Average Encode Time: 73.745452ms
    main_test.go:64: Success Count: 15
    main_test.go:65: Error Count: 0
    main_test.go:67: Average Encode Time: 69.803093ms
PASS
ok      github.com/eikarna/GoDat        10.155s
```
### Decoder (8 Cores, 4x2.4GHz, 4x1.8GHz)
```go
goos: android
goarch: arm64
pkg: github.com/eikarna/GoDat
BenchmarkDecoder-8            19          62299668 ns/op        19138962 B/op    583782 allocs/op
--- BENCH: BenchmarkDecoder-8
    main_test.go:110: Success Count: 1
    main_test.go:111: Error Count: 0
    main_test.go:113: Average Decode Time: 67.587083ms
    main_test.go:110: Success Count: 16
    main_test.go:111: Error Count: 0
    main_test.go:113: Average Decode Time: 61.871959ms
    main_test.go:110: Success Count: 19
    main_test.go:111: Error Count: 0
    main_test.go:113: Average Decode Time: 62.232516ms
PASS
ok      github.com/eikarna/GoDat        2.270s
```
You can test it by yourself, remember that the program use memory-map logic where all decoded/parsed Binary/JSON file is mapped to memory for ensure the process is fast and efficient.

## Special Thanks
- [GuckTubeYT](https://github.com/GuckTubeYT)
- [KipasGTS](https://github.com/KipasGTS)
