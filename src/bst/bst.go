package bst

import (
	"log"
	//"math"
	"math/rand"
	"time"
)

func RunInsert() {
	datasetSizes := []int{100, 1000, 2000, 5000, 10000, 20000, 50000, 100000}

	for _, datasetSize := range datasetSizes {
		result := runInserts(datasetSize)
		log.Printf("Inserted %d elements for %d nanoseconds", datasetSize, result)
	}
}

func RunSearch() {
	datasetSizes := []int{100, 1000, 2000, 5000, 10000, 20000, 50000, 100000}

	for _, datasetSize := range datasetSizes {
		result := runFinds(datasetSize)
		log.Printf("Find element from %d nodes for %d nanoseconds", datasetSize, result)
	}
}

func RunDelete() {
	datasetSizes := []int{100, 1000, 2000, 5000, 10000, 20000, 50000, 100000}

	for _, datasetSize := range datasetSizes {
		result := runDeletes(datasetSize)
		log.Printf("Delete element from %d nodes for %d nanoseconds", datasetSize, result)
	}
}

func printPreOrder(n *Node) {
	if n == nil {
		return
	} else {
		log.Printf("%d ", n.key)
		printPreOrder(n.left)
		printPreOrder(n.right)
	}
}

func runInserts(size int) int64 {
	var timestampSum int64
	iterations := 20

	for i := 0; i < iterations; i++ {
		dataset := makeRandomDataset(size)
		before := makeNanoTimestamp()
		tree := &Tree{}

		for _, element := range dataset {
			tree.Insert(element)
		}

		timestampSum += makeNanoTimestamp() - before
	}

	return int64(timestampSum / int64(iterations))
}

func runFinds(size int) int64 {
	var timestampSum int64
	iterations := 20

	for i := 0; i < iterations; i++ {
		dataset := makeRandomDataset(size)
		tree := &Tree{}
		for _, element := range dataset {
			tree.Insert(element)
		}

		before := makeNanoTimestamp()
		find(tree.root, rand.Intn(50000))
		timestampSum += makeNanoTimestamp() - before
	}

	return int64(timestampSum / int64(iterations))
}

func runDeletes(size int) int64 {
	var timestampSum int64
	iterations := 20

	for i := 0; i < iterations; i++ {
		dataset := makeRandomDataset(size)
		tree := &Tree{}
		for _, element := range dataset {
			tree.Insert(element)
		}

		before := makeNanoTimestamp()
		tree.Remove(dataset[int(size/2)])
		timestampSum += makeNanoTimestamp() - before
	}

	return int64(timestampSum / int64(iterations))
}

func makeRandomDataset(size int) []int {
	dataset := make([]int, 0, size)

	for i := 0; i < size; i++ {
		rand.Seed(time.Now().UnixNano())
		dataset = append(dataset, rand.Intn(50000))
	}

	return dataset
}

func makeNanoTimestamp() int64 {
	return time.Now().UnixNano()
}
