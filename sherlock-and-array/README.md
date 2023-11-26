# Sharlock and Array

```go
func BalancedSums(arr []int32) string {
	var total int32
	for _, num := range arr {
		total += num
	}

	var leftSum int32
	for i := 0; i < len(arr); i++ {
		if leftSum == total-leftSum-arr[i] {
			return "YES"
		}
		leftSum += arr[i]
	}
	return "NO"
}
```


```go
var leftSum int32
if leftSum == total-leftSum-arr[i] {
    return "YES"
}
```
上の条件が成り立つときそれは、配列 arr がある位置 `i` で左側の合計が右側の合計と等しいという状況になる。  
`leftSum` は位置 `i` までの左側の合計、 `total` は配列全体の合計、 `arr[i]` は位置 `i` の要素

例えば `[1, 2, 3, 3]` という配列を受け取ったとする。 `total` は 9、 `leftSum` は 0 `i` も 0の状態。  
このとき、右側の合計を知りたければ `total - leftSum` とすることで、右側の合計を知ることができる。ただ、ある要素 `i` を基準とした左右の合計の計算が必要なため、`arr[i]` の値は取り除く必要がある。これを考慮すると以下なる。

```go
var leftSum int32
var rightSum := total - leftSum - arr[i]

if leftSum == rightSum {
    return "YES"
}
```
