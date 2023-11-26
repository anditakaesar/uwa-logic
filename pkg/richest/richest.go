package richest

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	MAX_LIST = 10000000 // at most 10_000_000
)

func Run() {
	startTime := time.Now()
	bankAccounts := GetTestRichest()
	totalTime := time.Since(startTime)
	fmt.Printf("Total time generating test: %v\n", totalTime)

	startTime = time.Now()
	fmt.Printf("Length of test: %d\nRichest: %d\n", len(bankAccounts), GetRichest(bankAccounts))
	totalTime = time.Since(startTime)
	fmt.Printf("Total time taken: %v\n", totalTime)

}

func GetTestRichest() [][]int {
	result := make([][]int, MAX_LIST)

	var wg sync.WaitGroup
	for i := 0; i < MAX_LIST; i++ {
		wg.Add(1)
		length := 1 + rand.Intn(9)
		go func(i, length int) {
			defer wg.Done()
			result[i] = generateRandomList(length)
		}(i, length)
	}

	wg.Wait()

	return result
}

func generateRandomList(length int) []int {
	list := make([]int, length)
	for i := 0; i < length; i++ {
		list[i] = 1 + rand.Intn(10)
	}
	return list
}

func GetRichest(bank [][]int) int {
	var total, currentAccTotal int
	for acc := range bank {
		currentAccTotal = 0
		for i := range bank[acc] {
			currentAccTotal += bank[acc][i]
		}
		if total < currentAccTotal {
			total = currentAccTotal
		}
	}

	return total
}
