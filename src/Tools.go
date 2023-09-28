package PROJETRED

func IsAlpha(s string) bool {
	for _, element := range s {
		if !(65 <= element && element <= 90 || 97 <= element && element <= 122 || element == 45) {
			return false
		}
	}
	return true
}
