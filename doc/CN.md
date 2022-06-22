# 导入

```go
import "github.com/wyb199877/sliceutil"
```

# 操作类函数

## Foreach

遍历切片中的元素

```go
func Foreach[T any](
  slice []T, 
  action func(item T, index int, slice []T)
)
```

```go
sl := []int{1, 2, 3}
sliceutil.Foreach(sl, func(it int, idx int, s []int) {
  // ... do something
})
```

## Map

切片内元素转换

```go
func Map[T, R any](
  slice []T,
  action func(item T, index int, slice []T) R,
) []R
```

```go
sl := []int{1, 2, 3}
strs := sliceutil.Map(sl, func(it int, idx int, s []int) string {
  return strconv.Itoa(it)
})
// strs = ["1" "2" "3"]
```

## Filter

切片元素过滤器

```go
func Filter[T any](
  slice []T,
  action func(value T, index int, slice []T) bool,
) []T 
```

```go
sl := []int{1, 2, 3}
odds := sliceutil.Filter(sl, func(it int, idx int, s []int) bool {
  return it%2 != 0
})
// odds = [1 3]
```

## Reduce

切片元素累加器（未实现初始值）

```go
func Reduce[T any](
  slice []T,
  action func(result T, item T, index int, slice []T) T,
) T
```

```go
sl := []int{1, 2, 3}
sum := sliceutil.Reduce(sl, func(ret int, it int, idx int, s []int) int {
  return ret + it
})
// sum = 6
```

# 查找类函数

## Find

查找切片内第一个满足条件的元素

```go
func Find[T any](
  slice []T,
  action func(item T, index int, slice []T) bool,
) (t T, hasFound bool)
```

```go
sl := []int{1, 2, 3}
target, hasFound := sliceutil.Find(sl, func(it int, idx int, s []int) bool {
  return it%2 != 0
})
// target = 1, hasFound = true

target, hasFound = sliceutil.Find(sl, func(it int, idx int, s []int) bool {
  return it > 3
})
// target = 0, hasFound = false
```

## FindIndex

查找切片内第一个满足条件元素的索引值

```go
func FindIndex[T any](
  slice []T,
  action func(item T, index int, slice []T) bool,
) int
```

```go
sl := []int{1, 2, 3}
index := sliceutil.FindIndex(sl, func(it int, idx int, s []int) bool {
  return it%2 != 0
})
// index = 0

index = sliceutil.FindIndex(sl, func(it int, idx int, s []int) bool {
  return it > 3
})
// index = -1
```

## IndexOf

查找切片内第一个等于目标值元素的索引值

```go
func IndexOf[T comparable](slice []T, target T) int
```

```go
sl := []int{1, 2, 3}
index := sliceutil.IndexOf(sl, 1)
// index = 0

index = sliceutil.IndexOf(sl, 4)
// index = -1
```

## LastIndexOf

查找切片内最后一个等于目标值元素的索引值

```go
func LastIndexOf[T comparable](slice []T, target T) int
```

```go
sl := []int{1, 2, 3, 1}
index := sliceutil.LastIndexOf(sl, 1)
// index = 3

index = sliceutil.LastIndexOf(sl, 4)
// index = -1
```

## Includes

判断切片内是否存在与目标值相等的元素

```go
func Includes[T comparable](slice []T, target T) bool
```

```go
sl := []int{1, 2, 3}
isExist := sliceutil.Includes(sl, 1)
// isExist = true

isExist = sliceutil.Includes(sl, 4)
// isExist = false
```

## Some

判断切片内是否存在满足条件的元素

```go
func Some[T any](
  slice []T,
  action func(item T, index int, slice []T) bool,
) bool
```

```go
sl := []int{1, 3, 5}
isExist := sliceutil.Some(sl, func(it int, idx int, s []int) bool {
  return it%2 != 0
})
// isExist = true

isExist = sliceutil.Some(sl, func(it int, idx int, s []int) bool {
  return it%2 == 0
})
// isExist = false
```

## Every

判断切片内的元素是否均满足条件

```go
func Every[T any](
  slice []T,
  action func(item T, index int, slice []T) bool,
) bool
```

```go
sl := []int{1, 3, 5}
isEvery := sliceutil.Every(sl, func(it int, idx int, s []int) bool {
  return it%2 != 0
})
// isEvery = true

isEvery = sliceutil.Every(sl, func(it int, idx int, s []int) bool {
  return it%2 == 0
})
// isEvery = false
```

# 修改类函数

**这一类的函数均直接操作切片地址指向的内存空间**

## Insert

向切片内指定索引处插入元素

```go
func Insert[T any](sp *[]T, index int, items ...T)
```

```go
sl := []int{1, 3, 4}
sliceutil.Insert(&sl, 1, 2)
// sl = [1 2 3 4]
```

## Remove

从切片中移除指定索引的元素（请确保该位置有元素存在）

```go
func Remove[T any](sp *[]T, index int)
```

```go
sl := []int{1, 3, 4}
sliceutil.Remove(&sl, 1)
// sl = [1 4]
```

## Push

向切片尾部插入元素

```go
func Push[T any](sp *[]T, items ...T)
```

```go
sl := []int{1, 2, 3}
sliceutil.Push(&sl, 4)
// sl = [1 2 3 4]
```

## Pop

弹出并返回切片的最后一个元素（请勿对空切片操作）

```go
func Pop[T any](sp *[]T) T
```

```go
sl := []int{1, 2, 3}
last := sliceutil.Pop(&sl)
// last = 3, sl = [1 2]
```

## Unshift

向切片首部插入元素

```go
func Unshift[T any](sp *[]T, items ...T)
```

```go
sl := []int{2, 3, 4}
sliceutil.Unshift(&sl, 1)
// sl = [1 2 3 4]
```

## Shift

弹出并返回切片的第一个元素（请勿对空切片操作）

```go
func Shift[T any](sp *[]T) T
```

```go
sl := []int{1, 2, 3}
first := sliceutil.Shift(&sl)
// first = 1, sl = [2 3]
```

# 反转类函数

## Reverse

反转切片内的元素顺序，**仅拷贝，不操作内存空间**

```go
func Reverse[T any](slice []T) []T
```

```go
sl := []int{1, 2, 3}
sl = sliceutil.Reverse(sl)
// sl = [3 2 1]
```

## ReverseSelf

反转切片内的元素顺序，**直接操作内存空间**

```go
func ReverseSelf[T any](slice []T)
```

```go
sl := []int{1, 2, 3}
sliceutil.ReverseSelf(sl)
// sl = [3 2 1]
```

# 转换类函数

## ToMap

将切片转换为 map

```go
func ToMap[T any, K comparable, V any](
  slice []T,
  key func(item T) K,
  value func(item T) V,
) map[K]V 
```

```go
sl := []string{"a", "b", "c"}
mp := sliceutil.ToMap(sl, func(it string) string {
  return strings.ToUpper(it)
}, func(it string) string {
  return it
})
// mp = map[string]string{"A":"a", "B":"b", "C":"c"}
```
