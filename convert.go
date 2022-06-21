package sliceutil

/**
 * Description: ToMap
 * Author: Sirw
 * DateTime: 2022/6/21 20:56
 */

func ToMap[T any, K comparable, V any](
	slice []T,
	key func(item T) K,
	value func(item T) V,
) map[K]V {
	mp := make(map[K]V)
	for _, item := range slice {
		mp[key(item)] = value(item)
	}
	return mp
}
