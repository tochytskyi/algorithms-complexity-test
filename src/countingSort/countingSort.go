package countingSort

import (
	"fmt"
	"math/rand"
	"time"
)

func Run() {
	testData := map[int][]int{
		100:   {10, 100, 1000, 5000, 20000},
		1000:  {10, 100, 1000, 5000, 20000},
		10000: {10, 100, 1000, 5000, 20000},
	}

	for size, ranges := range testData {
		for _, maxRange := range ranges {
			slice := makeRandomDataset(size, maxRange)
			before := makeNanoTimestamp()
			CountingSort(slice)
			timestamp := makeNanoTimestamp() - before
			fmt.Printf("Sort %d elements with items range %d: %d nanoseconds", size, maxRange, timestamp)
			fmt.Println()
		}
	}
}

func makeRandomDataset(size int, rangeMax int) []int {
	dataset := make([]int, 0, size)

	for i := 0; i < size; i++ {
		rand.Seed(time.Now().UnixNano())
		dataset = append(dataset, rand.Intn(rangeMax))
	}

	return dataset
}

func makeNanoTimestamp() int64 {
	return time.Now().UnixNano()
}
