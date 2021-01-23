package util

func InArrayString(s string, sArr []string) bool {
	for _, v := range sArr {
		if s == v {
			return true
		}
	}
	return false
}
