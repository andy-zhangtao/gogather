# strings
--
    import "github.com/andy-zhangtao/gogather/strings"

可以完成针对字符串的高级操作

## Usage

#### func  DouExstact

```go
func DouExstact(src string, sym string) ([]string, error)
```
DouExstact 标准截取. 从src中按照指定的sym进行截取sym之间的字符 例如从#ABC#abc#DEF#中截取出 ABC和DEF

#### func  RemoveMultipeSpace

```go
func RemoveMultipeSpace(oldStr string) (newStr string)
```
RemoveMultipeSpace 去除字符串中多余的空格

#### func  ReplaceAscii

```go
func ReplaceAscii(src string, seg []string) string
```
ReplaceAscii 替换src当中指定的字符串类型的控制字符 例如将 字符串类型的"\n"替换成 Ascii的 10(LF)

#### func  Reverse

```go
func Reverse(src, seg string) string
```
Reverse 将源字符串按照指定的分隔符进行反转 例如: src = www.mydomain.com.cn dist =
cn.com.mydomian.www

#### func  ReverseWithSeg

```go
func ReverseWithSeg(src, oldSeg, newSeg string) string
```
ReverseWithSeg 将源字符串按照指定的分隔符进行反转,同时使用新分隔符替代旧分隔符 例如: src = www.mydomain.com.cn,
newSeg="-" dist = cn-com-mydomian-www

#### func  SymExstact

```go
func SymExstact(src string, sym1, sym2 string) ([]string, error)
```
SymExstact 对称截取. 从src中截取sym1与sym2之间的字符 例如从{ABC}abc{CHD}中按照'{','}'进行截取，最后会截取到
ABC和CHD
