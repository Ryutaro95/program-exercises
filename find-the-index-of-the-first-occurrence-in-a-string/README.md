# Find the index of the first occurrence in a string
文字列内で最初に出現するインデックスを検索する関数を作成する。  
2 つの文字列 haystack と needle を指定すると、haystack 内で最初に出現した needle のインデックスを返します。needle が haystack の一部でない場合は -1 を返します。

```go
func strStr(haystack string, needle string) int
```

## Example
* needle = "sad" は haystack = "sadbutsad" の一部であり、"sad" が最初に出現したのは haystack[0] なので `0` が返ることを期待している
```
Input: haystack = "sadbutsad", needle = "sad"
Output: 0
```

* needle = "leeto" は haystack = "leetcode"の一部ではないのそれを示す `-1` を返すことを期待している
```
Input: haystack = "leetcode", needle = "leeto"
Output: -1
```

## Constraints
* `1 <= haystack.length, needle.length <= 10^4`
* `haystack` と　`needle` は英小文字のみで構成されている



## Approach1: ボツ
needle が空になったかどうかで判定する
haystack[0] == needle[0]
上記のように順にループを回していく
trueであれば、index番号を記録しておく
最終的に len(needle) == 0  であれば、記録しておいたindex番号を返す
len(needle) != 0 であれば、-1を返す
```
haystack: "leetcode"
needle:   "leeto"
"l" == "l" -> "eeto"
"e" == "e" -> "eto"
"e" == "e" -> "to"
"t" == "t" -> "o"
"c" == "o" -> "o"
return -1

haystack: "sadbutsad"
needle:   "sad"
"s" == "s" -> "ad"
"a" == "a" -> "d"
"d" == "d" -> ""
return 0


haystack: "sabbutsad"
needle:   "sad"
"s" == "s" -> "ad"
"a" == "a" -> "d"
"b" == "d" -> "d"
"b" == "d" -> "d"
```
上記のような場合は、このアルゴリズムでは対応しきれない

## Approach2:
needle文字数を取得して、haystack[0:len(needle)] == needleのように
順に比較していく、一致していなければ、haystack[1:len(needle)]ようにずらしていけば見つかる？

```
haystack: "leetcode"
needle:   "leeto"
needleLength := len(needle): 5
```