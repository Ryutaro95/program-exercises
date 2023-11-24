# Merge Two Sorted Lists

以下2つのリストを１つのソート済みリストとして連結して返す関数を作成

```go
list1: 1 -> 2 -> 4
list2: 1 -> 3 -> 4

return 1 -> 1 -> 2 -> 3 -> 4 -> 4
```

## 関数とデータ構造
```go
type ListNode struct {
    Val int
    Next *ListNode
}

func mergetTwoLists(list1, list2 *ListNode) *ListNodh
```

## Constraints
* 両リストのノード数は0..50の範囲内
* `-100 <= Node.Val <= 100`
* list1,2は昇順でソートされた状態
