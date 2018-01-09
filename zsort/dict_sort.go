package zsort

import "sort"

type dict []string
func (d dict) Len() int           { return len(d) }
func (d dict) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d dict) Less(i, j int) bool { return d[i] < d[j] }

// DictSort 字典排序
func DictSort(keys []string)([]string){
	var d dict
	d = keys
	sort.Sort(d)
	return d
}