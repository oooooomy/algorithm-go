package sort

// 冒泡排序
// 时间复杂度：O(n²)
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}

// 选择排序
// 时间复杂度：O(n²)
func selectionSort(arr []int) {
	n := len(arr)
	var min int
	for i := 0; i < n-1; i++ {
		min = i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

// 插入排序
// 时间复杂度：O(n²)（最坏情况），O(n)（最好情况）
func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		t := arr[i]
		j := i - 1
		for j >= 0 && t < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = t
	}
}

// 快速排序
// 时间复杂度：O(n log n)（平均），O(n²)（最坏）
func quickSort(arr []int) {}
