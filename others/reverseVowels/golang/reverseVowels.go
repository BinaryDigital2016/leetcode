package others

func reverseVowels(s string) string {
	b := []byte(s)
	i, j := 0, len(s)-1
	for i <= j {
		for i <= j && !isVowels(b[i]) {
			i++
		}
		for i <= j && !isVowels(b[j]) {
			j--
		}
		if i <= j {
			b[i], b[j] = b[j], b[i]
			i++
			j--
		}
	}
	return string(b)
}

func isVowels(c uint8) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
