package utils

func ArrayToInterface[T any](arr []T) []interface{} {
	ret := make([]any, len(arr))
	for idx, v := range arr {
		ret[idx] = v
	}
	return ret
}
