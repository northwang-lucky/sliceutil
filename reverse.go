package sliceutil

/**
 * Description: Reverse, ReverseSelf
 * Author: Sirw
 * DateTime: 2022/6/21 15:36
 */

// Reverse slice (no side effects)
func Reverse[T any](slice []T) []T {
	processed := make([]T, 0, len(slice))
	for i := len(slice) - 1; i >= 0; i-- {
		processed = append(processed, slice[i])
	}
	return processed
}

// Reverse slice
func ReverseSelf[T any](slice []T) {
	i, j := 0, len(slice)-1
	for i < j {
		tmp := slice[i]
		slice[i] = slice[j]
		slice[j] = tmp
		i++
		j--
	}
}
