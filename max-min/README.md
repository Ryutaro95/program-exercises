# HackerRank - Max Min

Tag: Go
Priority: Middle
Status: Not started
Updated At: 2023年11月26日 2:45

- Problem Solving (Basic) の問題

[https://www.hackerrank.com/challenges/one-month-preparation-kit-angry-children/problem?isFullScreen=true&h_l=interview&playlist_slugs[]=preparation-kits&playlist_slugs[]=one-month-preparation-kit&playlist_slugs[]=one-month-week-two](https://www.hackerrank.com/challenges/one-month-preparation-kit-angry-children/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=preparation-kits&playlist_slugs%5B%5D=one-month-preparation-kit&playlist_slugs%5B%5D=one-month-week-two)

## 概要

与えられた整数配列 `arr` からサイズ `k` の連続した要素を選び、その中で最大と最小の差である絶対値(不公平さ)を計算する

Example:

```go
Input:
	arr = [1, 2, 3, 4, 10, 20, 30, 40, 100, 200]
	k = 4
Output
	3
```

考え方としては以下。

アルゴリズム: スライディングウィンドウ

```go
{} <- の範囲がwindow

[{1, 2, 3, 4}, 10, 20, 30, 40, 100, 200]
	3 = max: 4 - min:1

[1, {2, 3, 4, 10}, 20, 30, 40, 100, 200]
	8 = max: 10 - min: 2

[1, 2, {3, 4, 10, 20}, 30, 40, 100, 200]
	17 = max: 20 - min: 3

[1, 2, 3, {4, 10, 20, 30}, 40, 100, 200]
	26 = 	max: 30 - min: 4

[1, 2, 3, 4, {10, 20, 30, 40}, 100, 200]
	30 = max: 40 - min: 10	

[1, 2, 3, 4, 10, {20, 30, 40, 100}, 200]
	80 = max: 100 - min: 20	

[1, 2, 3, 4, 10, 20, {30, 40, 100, 200}]
	170 = max: 200 - min: 30

3 = math.Min(3, 8, 17, 26, 30, 80, 170)
となるので、3が正しいOutputになる
```

## 実装

- [Ryutaro95/program-exercises/max-min - GitHub](https://github.com/Ryutaro95/program-exercises/blob/main/max-min/main.go)
- [The Go PlayGround](https://go.dev/play/p/VwAo5abgQNE)

```go
func maxMin(k int32, arr []int32) int32 {
	unfairness := int32(math.MaxInt32)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	for i := 0; i < len(arr)-(int(k)-1); i++ {
		min := arr[i]
		max := arr[int(k)+i-1]
		difference := max - min
		if unfairness > difference {
			unfairness = difference
		}
	}

	return unfairness
}
```

### 解説

```go
unfairness := int32(math.MaxInt32)
```

`math.MaxInt32` で初期化している。 `math.MaxInt32` はint32型の最大値を表しており、これは後で最小の不公平さを保持する際に、どらちらが最小か比較するために初期値として使用される。

[pkg.go.dev/math#MaxInt32](https://pkg.go.dev/math#MaxInt32)

```go
sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
```

`arr` を昇順にソートしている。これにより後で `k` サイズの範囲を取りやすくなる

[pkg.go.dev/sort#Slice](https://pkg.go.dev/sort#Slice)

```go
for i := 0; i < len(arr)-(int(k)-1); i++ {
```

`len(arr) - (int(k) - 1)` 未満の範囲でループしている。この範囲は `k` サイズの範囲を保つためのインデックス範囲

```go
min := arr[i]
max := arr[int(k)+i-1]
```

現在の範囲の最小値と最大値を設定している。 `min` は `i` から、 `max` は `k` 個進んだ位置から取得っしている。

既に上の処理でソートしているので、 `k` 範囲の先頭が最小値で `k` の末尾が最大値で有ることが保証されている。

```go
difference := max - min

if unfairness > difference {
	unfairness = difference
}
```

現在の範囲の最大値と最小値の差を計算して、 `unfairness > difference` でこれまでの最小不公平さをより小さいものに更新している 。