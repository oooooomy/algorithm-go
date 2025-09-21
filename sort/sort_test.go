package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	arr := []int{3, 5, 8, 4, 6}
	bubbleSort(arr)
	t.Log(arr)
}

func TestSelectionSort(t *testing.T) {
	arr := []int{3, 5, 8, 4, 6}
	selectionSort(arr)
	t.Log(arr)
}

func TestInsertionSort(t *testing.T) {
	arr := []int{3, 5, 8, 4, 6}
	insertionSort(arr)
	t.Log(arr)
}
