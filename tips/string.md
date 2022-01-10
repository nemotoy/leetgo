
# 文字列操作

```golang

var s = "abcdefg"
var want = "c"

for i, r := range s {
  if string(r) == want {
    fmt.Printf("found(index: %d)",i)
  }
}

// 文字列から先頭n文字を取得する
s := "Hello, playground"
pref := s[:5]
fmt.Println("expeted: Hello, got: ", pref)

// 文字列がどのアルファベットで構成されているかを表現する
var s = "abc"
var count [26]int

for _, r := range s {
    count[r-'a'] += 1
}
fmt.Println("expected: [1 1 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], got: ", count)

// codepoint表現
var chars [123]int
for _, r := range text {
  chars[r] += 1
}

// 文字列のあるindex要素を削除する
s := "abc"
i := 1
s = s[:i] + s[i+1:]
fmt.Println("expeted: ac, got: ", s)

// 隣接する文字列が一致しないようするために削除する要素数
var result 32
for i := 0; i < len(s)-1; i++ {
    if s[i] == s[i+1] {
        result++
    }
}

func isNumeric(r rune) bool {
	return r < '0' || r > '9'
}

// a->0, b->1 ...
func sum(s string) int {
	ret := 0
	for _, r := range s {
		// 桁数を上げるために*10
		ret = ret*10 + int(r-'a')
	}
	return ret
}

// reverse
for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
  ret[i], ret[j] = ret[j], ret[i]
}

// 回文
func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// Intから桁数を算出する
digits := 0
for n != 0 {
  n /= 10
  digits++
}

// 素数
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 最大公約数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// reverse s from string
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// substring/部分文字列
for i := 0; i < l; i++ {
	for j := i; j < l && isXxx(); j++ {
	}
}

// 文字列が整数か
num, err := strconv.Atoi(v)
if err == nil {
}

/*
  ## TODO
  - 文字列の出現回数
  - 部分文字列
  - 後方からの要素取得
  - アルファベット
*/

```
