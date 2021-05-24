# 配列操作

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
- [slices-intro](https://blog.golang.org/slices-intro)
- [slices](https://blog.golang.org/slices)
- [effective_go.html#slices](https://golang.org/doc/effective_go.html#slices)
- [spec#Appending_and_copying_slices](https://golang.org/ref/spec#Appending_and_copying_slices)
