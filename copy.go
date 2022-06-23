package sliceutil

// Copy Slice
func Copy[T any](slice []T) []T {
	copied := make([]T, len(slice))
	copy(copied, slice)
	return copied
}

// Cut, return a copy
func Cut[T any](slice []T, start int, end int) []T {
	copied := Copy(slice)
	return copied[start:end]
}

// Cut from first element, return a copy
func CutFromFirst[T any](slice []T, end int) []T {
	copied := Copy(slice)
	return copied[:end]
}

// Cut to last element, return a copy
func CutToLast[T any](slice []T, start int) []T {
	copied := Copy(slice)
	return copied[start:]
}
