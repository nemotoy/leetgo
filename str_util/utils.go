package strutil

func CountVowels(s string) int {
	ret := 0
	for _, r := range s {
		switch r {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			ret++
		}
	}
	return ret
}
