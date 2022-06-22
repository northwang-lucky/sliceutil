# Slice Util

Implementing array manipulation functions in JavaScript in Golang

# Usage

English | [简体中文](./doc/README_CN.md)

## Import

```go
import "github.com/wyb199877/sliceutil"
```

## Operation Functions

### Foreach

Iterate over the elements in the slice

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

### Map

Conversion of elements in slices

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

### Filter

Slice element filter

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

### Reduce

Slice element accumulator (no initial value implemented)

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

## Search Functions

### Find

Find the first element that satisfies the condition in the slice

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

### FindIndex

Find the index of the first element that meets the condition in the slice

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

### IndexOf

Finds the index of the first element that equal to the target value in the slice

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

### LastIndexOf

Finds the index of the last element that equal to the target value in the slice

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

### Includes

Determine whether there is an element equal to the target value in the slice

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

### Some

Determine if there is an element in the slice that satisfies the condition

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

### Every

Determines whether all elements in the slice satisfy the condition

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

## Modification Functions

**Functions under this heading operate directly on the memory space to which the slice address points**

### Insert

Insert elements at the position of index in the slice

```go
func Insert[T any](sp *[]T, index int, items ...T)
```

```go
sl := []int{1, 3, 4}
sliceutil.Insert(&sl, 1, 2)
// sl = [1 2 3 4]
```

### Remove

Removes an element from a slice at the position of index (make sure there is an element at that position)

```go
func Remove[T any](sp *[]T, index int)
```

```go
sl := []int{1, 3, 4}
sliceutil.Remove(&sl, 1)
// sl = [1 4]
```

### Push

Insert elements at the end of slice

```go
func Push[T any](sp *[]T, items ...T)
```

```go
sl := []int{1, 2, 3}
sliceutil.Push(&sl, 4)
// sl = [1 2 3 4]
```

### Pop

Pops up and returns the last element of a slice (do not operate on empty slices)

```go
func Pop[T any](sp *[]T) T
```

```go
sl := []int{1, 2, 3}
last := sliceutil.Pop(&sl)
// last = 3, sl = [1 2]
```

### Unshift

Insert elements at the start of slice

```go
func Unshift[T any](sp *[]T, items ...T)
```

```go
sl := []int{2, 3, 4}
sliceutil.Unshift(&sl, 1)
// sl = [1 2 3 4]
```

### Shift

Pops up and returns the first element of a slice (do not operate on empty slices)

```go
func Shift[T any](sp *[]T) T
```

```go
sl := []int{1, 2, 3}
first := sliceutil.Shift(&sl)
// first = 1, sl = [2 3]
```

## Reversal functions

### Reverse

Reverse the order of the elements in the slice (only copy, no memory space operation)

```go
func Reverse[T any](slice []T) []T
```

```go
sl := []int{1, 2, 3}
sl = sliceutil.Reverse(sl)
// sl = [3 2 1]
```

### ReverseSelf

Reverse the order of the elements in the slice **(operate the memory space)**

```go
func ReverseSelf[T any](slice []T)
```

```go
sl := []int{1, 2, 3}
sliceutil.ReverseSelf(sl)
// sl = [3 2 1]
```

## Conversion functions

### ToMap

Convert slice to map

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