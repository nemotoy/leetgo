# Tips in Go

## 文字列操作

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
*/

```

## 配列操作

Goでは配列ではなくスライス。

```golang

var nums = []int{1,2,3,4,5}
var size = len(nums)

// 配列の各要素への処理
// numsをfor rangeで処理することも可能だが、indexの操作ができない
// numsの要素数は5で、長さは5。iは0から4まで処理する。
for i := 0; i < size; i++ {
  v := nums[i]
  if v == 100 {
    return v
  }
  i++
}

// 配列から指定のindexの要素を削除する
nums := []int{1, 2, 3, 4, 5}
i := 2
nums = append(nums[:i], nums[i+1:]...)
fmt.Println("expeted: [1 2 4 5], got: ", nums)

// スライスから最後の要素を削除する
nums := []int{1, 2, 3, 4, 5}
nums = nums[:len(nums)-1]
fmt.Println("expeted: [1 2 3 4], got: ", nums)
```

## References

- [SliceTricks](https://github.com/golang/go/wiki/SliceTricks)
