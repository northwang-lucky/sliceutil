package sliceutil

/**
 * Description: Insert, Remove, Push, Pop, Unshift, Shift
 * Author: Sirw
 * DateTime: 2022/6/21 15:50
 */

// Insert element into slice at target index
func Insert[T any](sp *[]T, index int, items ...T) {
	slice := *sp
	processed := make([]T, len(slice)+len(items))
	for i := 0; i < len(slice); i++ {
		if i < index {
			processed[i] = slice[i]
		} else if i > index {
			processed[i+len(items)] = slice[i]
		} else {
			for j, newItem := range items {
				processed[i+j] = newItem
			}
			processed[i+len(items)] = slice[i]
		}
	}
	*sp = processed
}

// Remove an element from the slice
func Remove[T any](sp *[]T, index int) {
	slice := *sp
	*sp = append(slice[:index], slice[index+1:]...)
}

// Insert elements at the end of the slice
func Push[T any](sp *[]T, items ...T) {
	slice := *sp
	*sp = append(slice, items...)
}

// Pop the last element from the slice
func Pop[T any](sp *[]T) T {
	slice := *sp
	lastIndex := len(slice) - 1
	last := slice[lastIndex]
	Remove(sp, lastIndex)
	return last
}

// Insert elements at the start of the slice
func Unshift[T any](sp *[]T, items ...T) {
	slice := *sp
	*sp = append(items, slice...)
}

// Pop the first element from the slice
func Shift[T any](sp *[]T) T {
	slice := *sp
	first := slice[0]
	Remove(sp, 0)
	return first
}
