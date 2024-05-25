package main

import (
	"github.com/codecat/go-libs/log"
	"github.com/eikarna/GoDat/Components/Decoder"
	"github.com/eikarna/GoDat/Components/Encoder"
	. "github.com/eikarna/GoDat/Components/Enums"
	"github.com/goccy/go-json"
	"os"
	"sync"
	"testing"
	"time"
)

func BenchmarkEncoder(b *testing.B) {
	// Define Data Variables
	target, err := os.ReadFile("items.json") // just pass the file name
	if err != nil {
		log.Error(err.Error())
		return
	}
	data := &ItemInfo{}
	err = json.Unmarshal(target, data)
	if err != nil {
		log.Error(err.Error())
		return
	}
	// Define the number of concurrent to simulate
	concurrency := 1000
	successCount := 0
	errorCount := 0
	var totalTime time.Duration

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(concurrency)
		// start := time.Now()

		for j := 0; j < concurrency; j++ {
			go func() {
				defer wg.Done()
				start := time.Now()
				err := Encoder.Encode(data, "items.encoded.dat", start)
				totalTime += time.Since(start)
				if err != nil {
					b.Errorf("Error encoding: %v", err)
					errorCount++
					return
				} else {
					successCount++
					return
				}
			}()
		}

		wg.Wait()
	}

	b.StopTimer()

	b.Logf("Success Count: %d", successCount)
	b.Logf("Error Count: %d", errorCount)
	if successCount > 0 {
		b.Logf("Average Encode Time: %s", totalTime/time.Duration(successCount))
	} else {
		b.Log("No successful encode process")
	}
}

func BenchmarkDecoder(b *testing.B) {
	// Define the number of concurrent to simulate
	concurrency := 1000
	successCount := 0
	errorCount := 0
	var totalTime time.Duration

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(concurrency)
		// start := time.Now()

		for j := 0; j < concurrency; j++ {
			go func() {
				defer wg.Done()
				start := time.Now()
				_, err := Decoder.Decode("items.encoded.dat", start)
				totalTime += time.Since(start)
				if err != nil {
					b.Errorf("Error Decoding: %v", err)
					errorCount++
					return
				} else {
					successCount++
					return
				}
			}()
		}

		wg.Wait()
	}

	b.StopTimer()

	b.Logf("Success Count: %d", successCount)
	b.Logf("Error Count: %d", errorCount)
	if successCount > 0 {
		b.Logf("Average Decode Time: %s", totalTime/time.Duration(successCount))
	} else {
		b.Log("No successful Decode process")
	}
}
