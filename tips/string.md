
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
/*
  ## TODO
  - 文字列の出現回数
  - 部分文字列
  - 後方からの要素取得
  - reverseString
  - アルファベット
*/

```
