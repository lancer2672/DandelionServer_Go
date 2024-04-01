package utils

func StringInSlice(s string, list []string) bool {
	for _, str := range list {
		if str == s {
			return true
		}
	}
	return false
}
