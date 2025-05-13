package utils

func FilterList[T any](list []T, test func(T) bool) (ret []T) {
	for _, item := range list {
		if test(item) {
			ret = append(ret, item)
		}
	}
	return
}
