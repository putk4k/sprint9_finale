package main

import (
	"fmt"
	"log"
	"math/rand"
	"slices"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	sl := make([]int, size)
	for i := 0; i < size; i++ {
		sl[i] = int(r.Uint32())
	}
	return sl
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {

	if len(data) == 0 {
		log.Println("Пустой слайс")
		return 0
	}

	max := slices.Max(data)
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	var wg sync.WaitGroup
	var mu sync.Mutex

	if len(data) == 0 {
		log.Println("Пустой слайс")
		return 0
	}

	chunks := CHUNKS
	if len(data) < chunks {
		chunks = len(data)
	}

	maxSlice := make([]int, 0, chunks)
	lenChunck := len(data) / chunks

	for i := 0; i < chunks; i++ {
		wg.Add(1)

		min := i * lenChunck
		max := min + lenChunck
		if i == chunks-1 {
			max = len(data)
		}

		slice := data[min:max]

		go func(slice []int) {
			defer wg.Done()

			max := slices.Max(slice)

			mu.Lock()
			maxSlice = append(maxSlice, max)
			mu.Unlock()

		}(slice)
	}
	wg.Wait()

	return slices.Max(maxSlice)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
