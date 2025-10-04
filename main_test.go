package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func TestGenerateSlice(t *testing.T) {

	slice := generateRandomElements(0)
	assert.Empty(t, slice)

	slice = generateRandomElements(1)
	assert.NotEmpty(t, slice)

}

func TestGenerateSliceSize(t *testing.T) {

	for _, size := range []int{10, 100, 12345} {
		slice := generateRandomElements(size)
		assert.Equal(t, size, len(slice))
	}
}

func TestGenerateSliceRandomNums(t *testing.T) {
	slice := generateRandomElements(100)
	allSame := true
	for i := 1; i < len(slice); i++ {
		if slice[i] != slice[0] {
			allSame = false
			break
		}
	}
	assert.False(t, allSame)
}

func TestMaxNumEmpty(t *testing.T) {
	var slice []int
	max := maximum(slice)
	assert.Equal(t, 0, max)
}

func TestMaxOneNum(t *testing.T) {
	slice := []int{3}
	max := maximum(slice)
	assert.Equal(t, 3, max)
}

func TestMaxManyNum(t *testing.T) {
	slice := []int{3, 5, 1, 4, 2, 9, 4, 11}
	max := maximum(slice)
	assert.Equal(t, 11, max)
}

func TestMaxEqualNums(t *testing.T) {
	slice := []int{6, 6, 6, 6, 6}
	max := maximum(slice)
	assert.Equal(t, 6, max)
}

func TestMaxChunksEmpty(t *testing.T) {
	slice := []int{}
	max := maxChunks(slice)
	assert.Equal(t, 0, max)
}

func TestMaxChunksSmallSlice(t *testing.T) {
	slice := []int{3, 5, 1, 4, 2}
	max := maxChunks(slice)
	assert.Equal(t, 5, max)
}
