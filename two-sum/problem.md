# Two Sum
整数配列 `nums` と整数 `target` が与えられたとき、 配列内の2つの数値合計が `target` になる値の添字を返す。

入力の値は1つの解答を必ず持っており、解が存在しないパターンを考慮しなくてもよく、添字もどのような順番でもよい。


## Example 
Input1:
```go
nums := [2,7,11,15]
target := 9

nums[0] + nums[1] == target
```
Output1:
```go
return [0, 1]
```

Input2:
```go
nums := [3,2,4]
target := 6

nums[1] + nums[2] == target
```
Output2:
```go
return [1, 2]
```

Input3:
```go
nums := [3,3]
target := 6

nums[0] + nums[1] == target
```
Output3:
```go
return [0, 0]
```

## Constraints
```go
2 <= nums.length <= 10^4
-10^9 <= nums[i] <= 10^9
-10^9 <= target <= 10^9
```

## Approach
過去の履歴を格納するmapを用意して、 `for range` で各値がキーとしてmapに含まれているかチェックする。

値が含まれていなければ `target - nums[i]` の値をキーとしてmapに格納する。  
`n` とマップ保管されているキーが一致すれば、valueとして残していた以前のindex番号が取得できるので、現在のindex番号と合わせてreturnできる
```go
// nums = [2,7,11,15]
// target = 9
for currentIdx, n := range nums {
    if prevIdx, ok := m[n]; ok {
        return []int{prevIdx, currentIdx}
    }
    // 9 - 2 = 7
    m[target-n] = i
    // m{7: 1}
}
```

## Reversion

こっちのほうが早くなった
```go
for i, n := range nums {
    if v, ok := m[target-n]; ok {
        return []int{v, i}
    }
    m[n] = i
}
```