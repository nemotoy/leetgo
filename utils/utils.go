package utils

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

func IsVowel(r byte) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func IsNumeric(r rune) bool {
	return r < '0' || r > '9'
}

func IsPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ReverseSubString(s string, a, b int) string {
	runes := []rune(s[a : b+1])
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return s[:a] + string(runes) + s[b+1:]
}

func ReplaceSubString(s string, a, b int, p string) string {
	return s[:a] + p + s[b+1:]
}

func ReplaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func BiggerLexicographical(a, b string) string {
	if a > b {
		return a
	} else if a < b {
		return b
	}
	return ""
}

// 対象indexの文字コード差分を返す
func DiffCodeByIndex(a, b string, i int) rune {
	aa, bb := a[i], b[i]
	if aa < bb {
		return rune((bb - aa + 26) % 26)
	}
	return rune((aa - bb + 26) % 26)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x < 1 {
		return -x
	}
	return x
}

func Sum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// 整数の桁数
func NumOfDigit(n int) int {
	cnt := 0
	for n > 0 {
		n /= 10
		cnt++
	}
	return cnt
}

// 約数
func NumOfDivisor(n int) int {
	divisor := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			divisor++
		}
	}
	return divisor
}

// 公約数
func CommonDivisor(a, b int) int {
	min := a
	if a > b {
		min = b
	}
	cnt := 0
	for i := 1; i <= min; i++ {
		if a%i == 0 && b%i == 0 {
			cnt++
		}
	}
	return cnt
}

func SumAsStr(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += int(s[i]) - '0'
	}
	return sum
}

// 最大公約数
func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

// 階乗
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

// 回文
func IsPalindromeInt(x int) bool {
	if x < 0 {
		return false
	}
	tmp := x
	rev := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		rev = rev*10 + pop
	}
	return rev == tmp
}

// keyに該当するindexを返す二分探索。aはソート済み配列。該当しない場合は-1を返す。
func BinarySearchInts(a []int, key int) int {
	left, right := 0, len(a)-1
	for left <= right {
		mid := left + (right-left)/2
		if a[mid] == key {
			return mid
		} else if a[mid] > key {
			right = mid - 1
		} else if a[mid] < key {
			left = mid + 1
		}
	}
	return -1
}

// 累乗
// math.Pow10(5))
