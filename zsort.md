# zsort
--
    import "github.com/andy-zhangtao/gogather/zsort"

对无序序列进行排序

## Usage

#### func  BubbleSort

```go
func BubbleSort(data sort.Interface)
```
BubbleSort 冒泡排序. data必须实现sort包中的Interface接口

#### func  DictSort

```go
func DictSort(keys []string) []string
```
DictSort 字典排序

#### func  QuickSort

```go
func QuickSort(data sort.Interface)
```
QuickSort 快速排序, data必须要实现sort包的Interface接口

#### func  SelectSort

```go
func SelectSort(data sort.Interface)
```
SelectSort 选择排序, data必须实现sort包中的Interface接口

#### type Pair

```go
type Pair struct {
	Key   string
	Value int
}
```


#### type PairList

```go
type PairList []Pair
```


#### func  SortByValue

```go
func SortByValue(valueMap map[string]int) PairList
```
SortByValue 按照map的value进行降序排列

#### func (PairList) Len

```go
func (p PairList) Len() int
```

#### func (PairList) Less

```go
func (p PairList) Less(i, j int) bool
```

#### func (PairList) Swap

```go
func (p PairList) Swap(i, j int)
```
