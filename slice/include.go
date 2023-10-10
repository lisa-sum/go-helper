package slice

// Include 判断 item 是否在 array 中
// slice 切片, 可传递 []string | []int | []bool | []float32 | []float64
// item 元素, 可传递string | int | float32 | float64
// 返回值 bool
func Include[T []string | []int | []bool | []float32 | []float64, V string | int | float32 | float64](slice T, item V) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
