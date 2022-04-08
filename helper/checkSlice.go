package helper

func CheckSlice(arr []int, id int) bool {
	for _, item := range arr {
		if item == id {
			return true
		}
	}
	return false
}
