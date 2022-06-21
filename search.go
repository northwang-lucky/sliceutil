package sliceutil

/**
 * Description: Find, FindIndex, IndexOf, LastIndexOf, Includes, Some, Every
 * Author: Sirw
 * DateTime: 2022/6/21 20:35
 */

// Find the first eligible element in a slice
func Find[T any](
	slice []T,
	action func(item T, index int, slice []T) bool,
) (t T, hasFound bool) {
	for index, item := range slice {
		if action(item, index, slice) {
			return item, true
		}
	}
	return t, false
}

// Find the index of the first eligible element in the slice
func FindIndex[T any](
	slice []T,
	action func(item T, index int, slice []T) bool,
) int {
	for index, item := range slice {
		if action(item, index, slice) {
			return index
		}
	}
	return -1
}

// Find the index of the first element in the slice
// that is equal to the target value
func IndexOf[T comparable](slice []T, target T) int {
	for index, item := range slice {
		if item == target {
			return index
		}
	}
	return -1
}

// Find the index of the last element in the slice
// that is equal to the target value
func LastIndexOf[T comparable](slice []T, target T) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if target == slice[i] {
			return i
		}
	}
	return -1
}

// Determine if there is an element in the slice equal to the target value
func Includes[T comparable](slice []T, target T) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

// Determine if there is an element that matches the condition in the slice
func Some[T any](
	slice []T,
	action func(item T, index int, slice []T) bool,
) bool {
	for index, item := range slice {
		if action(item, index, slice) {
			return true
		}
	}
	return false
}

// Determine whether each element in the slice meets the condition
func Every[T any](
	slice []T,
	action func(item T, index int, slice []T) bool,
) bool {
	for index, item := range slice {
		if !action(item, index, slice) {
			return false
		}
	}
	return true
}
