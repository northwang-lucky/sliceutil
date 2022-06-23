package sliceutil

/**
 * Description: Foreach, Map, Filter, Reduce
 * Author: Sirw
 * DateTime: 2022/6/21 20:37
 */

// Iterate over each element in the slice
func Foreach[T any](slice []T, action func(item T, index int, slice []T)) {
	for index, item := range slice {
		action(item, index, slice)
	}
}

// Constructs a new element from the elements in the slice
func Map[T, R any](
	slice []T,
	action func(item T, index int, slice []T) R,
) []R {
	processed := make([]R, 0, len(slice))
	for index, item := range slice {
		processed = append(processed, action(item, index, slice))
	}
	return processed
}

// Filter out eligible elements in a slice
func Filter[T any](
	slice []T,
	action func(value T, index int, slice []T) bool,
) []T {
	processed := make([]T, 0, len(slice))
	for index, item := range slice {
		if action(item, index, slice) {
			processed = append(processed, item)
		}
	}
	return processed
}

// Accumulate each element in the slice
func Reduce[T any](
	slice []T,
	action func(result T, item T, index int, slice []T) T,
) T {
	var result T
	for index, item := range slice {
		result = action(result, item, index, slice)
	}
	return result
}
