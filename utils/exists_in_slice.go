package utils

func Exists(s string, slice []string) bool {
	for _, cur := range slice {
		if cur == s {
			return true
		}
	}
	return false
}
