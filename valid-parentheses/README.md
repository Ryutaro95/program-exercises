# Valid Parentheses
`(), {}, []` を含む文字列 s が与えられたとき、その入力文字列が有効かどうかを判定する関数を作成する

* 入力文字列は以下の場合に有効である
    * 開きカッコは同じ種類の閉じカッコで閉じなければならない
    * 開きカッコは正しい順序で閉じなければならない
    * 全ての閉じカッコは、対応する同じ型の開きカッコを持つ

## Example

example 1:
```go
Input: s = "()"
Output: true
```

example 2:
```go
Input: s = "()[]{}"
Output: true
```

example 3:
```go
Input: s = "(]"
Output: false
```

## Constraints
* `1 <= s.length <= 10^4`
* `s` は `() [] {}` のみで構成されている


## Approach

* 問題を単純化させるために入力されるカッコが `()` のみの場合を考えてみる  
* アプローチとしては、入力された文字列を左から順に処理していき、`(` に遭遇したとき用意したカウンターをインクリメントする
    * もし、閉じカッコ `)` を見つけた場合はカウンターをデクリメントする
    * カッコが対応しているとき、ループが終わった時点でカウンターは 0 を持っているはずなので true となる、そうでない場合は false となる
* ただ、このアプローチは `()` のみを想定されているため `(}` のような入力が合った場合は true と判定されてしまう

### Stacks


![](https://leetcode.com/problems/valid-parentheses/Figures/20/20-Valid-Parentheses-Recursive-Property.png)

色分けされたセルは括弧の開始と終了を示しているため、対応している括弧を見つけるたびに、文字列からその括弧を順に取り除く。  
そうすると、有効な文字列の場合は空の文字列が残ることになる。

```go
0 1 2 3 4 5 6 7 8 9
{ [ [ ] { } ] } ( )

↓  index 2, 3は一致するので取り除く

0 1 2 3 4 5 6 7
{ [ { } ] } ( ) { }

↓  同様にindex 2, 3は一致するので取り除く

0 1 2 3 4 5
{ [ ] } ( )

↓  index 1, 2は一致するので取り除く

0 1 2 3
{ } ( )

↓  index 0, 1は一致するので取り除く

0 1
( )

↓  index 0, 1は一致するので取り除くと空文字が残る

""
```

このアプローチをスタックを使って実装する
1. スタックを初期化する
2. 文字列を順に処理する
3. 開始括弧 `(, [, {` があれば、スタックにプッシュする
4. 閉じ括弧を見つけたら スタックの一番上の要素をチェックする
   1. それが、同じタイプの開き括弧であれば、それをスタックから取り出す
   2. そうでないなら、無効な式を意味するので false を返す
5. 最終的にスタックに要素が残っている場合は無効な式を意味するので false を返す