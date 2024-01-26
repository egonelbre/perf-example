package main

import (
	"fmt"
	"time"

	"github.com/loov/hrtime"
)

func main() {
	for i := 0; i < 10; i++ {
		_ = hrtime.Now()
		_ = time.Now()
	}

	var now [10000000]time.Time
	start := hrtime.Now()
	for i := range now {
		now[i] = time.Now()
	}
	finish := hrtime.Now()

	// let's look at system granularity
	var durations []time.Duration
	for i := range now[:len(now)-1] {
		durations = append(durations, now[i+1].Sub(now[i]))
	}

	histogram := hrtime.NewDurationHistogram(durations,
		&hrtime.HistogramOptions{
			BinCount:        10,
			NiceRange:       true,
			ClampMaximum:    float64(10 * time.Microsecond),
			ClampPercentile: 0,
		})

	fmt.Println(histogram)

	// let's look at measurement overhead
	fmt.Println(float64(finish-start)/float64(len(now)), "ns")
}
