package util

func RemoveValueFromIntSlice(s []int64, value int64) []int64 {
	var new []int64
	for _, v := range s {
		if v != value {
			new = append(new, v)
		}
	}
	return new
}

func RemoveIndex(s []int64, index int64) []int64 {
	ret := make([]int64, 0)
	ret = append(ret, s[:index]...)
	if len(ret) == 0 {
		return ret
	}
	return append(ret, s[index+1:]...)
}
