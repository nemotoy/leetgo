# leetgo

![ci](https://github.com/nemotoy/leetgo/workflows/ci/badge.svg)

learn [LeetCode](https://leetcode.com/) in Go.

- `list` directory manages problems related to the list node.

## Tips

### 文字列操作

```golang

var s = "abcdefg"
var want = "c"

for i, r := range s {
  if string(r) == want {
    fmt.Printf("found(index: %d)",i)
  }
}

```

### 配列操作

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

fmt.Println("expeted: [1 2 4 5], got: ",nums)
```

## References

- [SliceTricks](https://github.com/golang/go/wiki/SliceTricks)
