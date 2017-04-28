package zsort

import "sort"

// SelectSort 选择排序, data必须实现sort包中的Interface接口
func SelectSort(data sort.Interface) {

	for i := 0; i < data.Len()-1; i++ {
		// 假定首元素为最小元素
		min := i
		for j := min + 1; j < data.Len(); j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		// 将此次筛选出的最小元素放入最左边
		data.Swap(min, i)
	}
}
