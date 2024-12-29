package utils

func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V, len(maps))

	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}

	return result
}
