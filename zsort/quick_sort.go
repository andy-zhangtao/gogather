package zsort

import "sort"

// QuickSort 快速排序, data必须要实现sort包的Interface接口
func QuickSort(data sort.Interface) {
	quickSort(data, 0, data.Len())
}

func quickSort(data sort.Interface, begin, end int) {
	if begin == end {
		return
	}

	// 假定数组第一个元素为中位数
	mid, i := begin, begin+1
	// 当head==tail则表示当前分区数据排序完毕
	head, tail := begin, end-1

	for head < tail {
		if data.Less(i, mid) {
			data.Swap(head, i)
			head++
			i++
			mid = head
		} else {
			data.Swap(tail, i)
			tail--
		}
	}

	// 对左分区进行排序
	quickSort(data, begin, mid)
	// 对右分区进行排序
	quickSort(data, mid+1, end)
}
