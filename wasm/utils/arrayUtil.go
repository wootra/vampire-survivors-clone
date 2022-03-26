package utils

func Search(arr []uint64, search ...uint64) bool {
	for _, id := range arr {
		for _, match := range search {
			if id == match {
				return true
			}
		}
	}
	return false
}
