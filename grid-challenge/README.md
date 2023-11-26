# Grid Challenge

`a ~ z` までの範囲からからなる、正方形のグリッドが与えられる。各行のアルファベットを昇順に並び替える。  
そのとき、列が上から下にかけて昇順でアルファベット順に並んでいるかを判定する。  
return "YES" or "NO"

```go
e b a c d
f g h i j
o l m k n
t r p q s
x y w u v
```
上の例のとき `e -> f -> o -> t -> x` と次の列 `b -> g -> l -> r -> y` はアルファベット順なので問題ない。



## Constraints
* `1 ≦ t ≦ 100`
* `1 ≦ n ≦ 100`

## 実装
```go
func gridChallenge(grid []string) string {
	for i := 1; i < len(grid); i++ {
		prevLine := []rune(grid[i-1])
		currentLine := []rune(grid[i])
		sort.Slice(prevLine, func(i, j int) bool { return prevLine[i] < prevLine[j] })
		sort.Slice(currentLine, func(i, j int) bool { return currentLine[i] < currentLine[j] })

		for j := 0; j < len(currentLine); j++ {
			if prevLine[j] > currentLine[j] {
				return "NO"
			}
		}
	}
	return "YES"
}
```


```go
for i := 1; i < len(grid); i++ {
```

ループのスタートを予め1つ進めておくことで、現在の行( `grid[1]` ) と一つ前の行 ( `grid[0]` )を一回のループで比較することができる。   

また、制約を見ると `1 ≦ n ≦ 100` とありgridスライス内の文字列が1つの可能性がある。  
この場合にスタートを0として始めてしまうと、存在しないindex番号を参照してしまう。そのため、予め 1 としておくことで、 `len(grid) == 1` のスライスが渡ってきたときはループに入らずそのまま `YES` を返すことができる。


```go
for i := 1; i < len(grid); i++ {
    prevLine := []rune(grid[i-1])
    currentLine := []rune(grid[i])
    sort.Slice(prevLine, func(i, j int) bool { return prevLine[i] < prevLine[j] })
    sort.Slice(currentLine, func(i, j int) bool { return currentLine[i] < currentLine[j] })
```
まずソートするために、文字列をルーンスライスに変換している。  
そして、ループに入った後に各行をソートしている。ループに入る前に全行をソートしても良いかなと思うが、前半にアルファベット順ではないことが判定された場合、ソートした分が無駄になると考えたので、必要な分だけループ内でソートする。

ただ、全ての列がアルファベット順である場合、全行を事前にソートしておくよりも処理回数が多くなってしまうので、検討が必要。今回のケースは最大行数が100なので、どちらを選んでも影響は少なさそう

```go
for i := 1; i < len(grid); i++ {
    // ...
    for j := 0; j < len(currentLine); j++ {
        if prevLine[j] > currentLine[j] {
            return "NO"
        }
    }
}
```
index `j` で各列のアルファベット(rune)を比較している。一つ前のアルファベットのほうが大きい場合、この `grid` 列はアルファベットではないので、 `NO` を返す。  
この判定が全て `false` つまりアルファベット順である場合は全てのループから抜けて　`YES` を返す。
