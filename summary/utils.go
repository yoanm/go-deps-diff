package summary

func inOrderMapIteratorHelper[K comparable, V any](theMap map[K]V, orderList []K, processor func(key K, val V)) {
	for _, key := range orderList {
		val, exists := theMap[key]
		if !exists {
			continue
		}

		processor(key, val)
	}
}
