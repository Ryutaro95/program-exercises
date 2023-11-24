# Longest Common Prefix
文字列の配列から最長の共通接頭辞文字列を見つける関数を書く。  
共通の接頭辞がない場合は、空の文字列 "" を返す。

## Example

```
Input: strs = ["flower","flow","flight"]
Output: "fl"
```

```
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
```

## Binary Search Approach
Binary Searchは、ソートされた配列やリスト内で特定の要素を見つけるための効率的なアルゴリズム。要素がソートされていることが前提のため、対象の要素を効率的に検索することが可能。

1. 中央の要素を選択する
2. 中央の要素と目標が一致するか
3. 目標が中央要素より小さい場合は左側を探索
4. 目標が中央要素より大きい場合は右側を探索
5. 要素が見つからない場合は終了

Time: O(log n)

このアルゴリズムを使ってこの問題を解決する  

文字列の中央文字が共通文字列ではない場合、文字列の後半は共通でないことが分かるので検索太陽から破棄する。
文字列の中央文字が共通文字列の場合、0..mid までは共通文字列なので、より長い共通接頭辞を見つけようとする。そのため前半を破棄する。

```go
"flower", "flow", "flight"
```

![](https://leetcode.com/media/original_images/14_lcp_binary_search.png)
