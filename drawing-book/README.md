# Drawing Book
* 目的のページまでどのくらいの回数めくると辿り着くかを計算する関数を用意する
* `Return `: 本の始まり、終わりで目的のページに辿り着く最短のめくりページ枚数を返す
* 1ページ目は必ず右側から始まる
* `pages`: 本のページ総数が与えられる
* `targetPage`: 目的のページが与えられる

```go
func pageCount(pages, targetPage int32) int32
```