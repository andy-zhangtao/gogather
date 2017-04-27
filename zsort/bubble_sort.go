package zsort

import "sort"

// BubbleSort 冒泡排序. data必须实现sort包中的Interface接口
func BubbleSort(data sort.Interface) {
	n := data.Len()
	for i := 0; i < n-1; i++ {
		isChange := false
		for j := 0; j < n-1-i; j++ {
			if data.Less(j, j+1) {
				data.Swap(j, j+1)
				isChange = true
			}
		}
		if !isChange {
			break
		}
	}
}
