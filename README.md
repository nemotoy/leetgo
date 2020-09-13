# leetgo

![ci](https://github.com/nemotoy/leetgo/workflows/ci/badge.svg)

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

```golang

var nums = []int{1,2,3,4,5}
var size = len(nums)

// 配列の各要素への処理
// numsをfor rangeで処理することも可能だが、incrementの操作ができない
for i := 0; i < size; i++ {
  v := nums[i]
  if v == 100 {
    return v
  }
  i++
}

```

## References

- [SliceTricks](https://github.com/golang/go/wiki/SliceTricks)
