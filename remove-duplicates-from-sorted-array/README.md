# Remove Duplicates From Sorted Array
整数配列 `nums` が非減少順にソートされている場合、重複を取り除く。  
重複を取り除いた後の配列の長さを返す

また、`nums` は重複が取り除かれた状態にしておく。  
そして以下のアサーションにクリアする形になっていれば解答は問題ないと判定される

```go
nums := []int{...}
expectedNums := []int{...}

k := removeDuplicates(nums)

assert k == len(exectedNums)
for i := 0; i < k; i++ {
    assert nums[i] == expectedNums[i];
}
```

## Example

example1:
```go
Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
```

example2:
```go
Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
```

## Approach
* Two Pointer
高速に進むポインタと遅く進むポインタを用意する
```
s := 高速ポインタ
l := 低速ポインタ
```

低速ポインタは `nums` の要素が入れ替わった際に進み、高速ポインタはループが繰り返されるたびに進む
```go
nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

l := 0
s := 1

if nums[l]: 0 != nums[s]: 0
    値をチェックして異なっていれば、nums[l+1]に nums[s]の要素を入れる
else 
    値が同じであれば何もしない
これを繰り返した結果を以下に示す

```

```go
初期値:
nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
l := 0
s := 1

step1:
    nums[l] 0 != nums[s] 0
    s++

    l := 0
    s := 2
    変更なし
    nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

step2:
    nums[l] 0 != nums[s] 1
    nums[l+1] = num[s]
    l++****
    s++

    l := 1
    s := 3
        nums[s]に書き換える
             ↓
    []int{0, 1, 1, 1, 1, 2, 2, 3, 3, 4}

step3:
    nums[l] 1 != nums[s] 1
    s++

    l := 1
    s := 4
    変更なし
    []int{0, 1, 1, 1, 1, 2, 2, 3, 3, 4}

step4:
    nums[l] 1 != nums[s] 1
    s++

    l := 1
    s := 5
    変更なし
    []int{0, 1, 1, 1, 1, 2, 2, 3, 3, 4}

step5:
    nums[l] 1 != nums[s] 2
    nums[l+1] = num[s]
    l++
    s++

    l := 2
    s := 6
            nums[s]に書き換える
                ↓
    []int{0, 1, 2, 1, 1, 2, 2, 3, 3, 4}

step6:
    nums[l] 2 != nums[s] 2
    l++
    s++

    l := 2
    s := 7
    何もしない
    []int{0, 1, 2, 1, 1, 2, 2, 3, 3, 4}

step6:
    nums[l] 2 != nums[s] 3
    nums[l+1] = num[s]
    l++
    s++

    l := 3
    s := 8
                 nums[s]に書き換える
                   ↓
    []int{0, 1, 2, 3, 1, 2, 2, 3, 3, 4}

これを繰り返していくとnumsの中身は以下になる
[0 1 2 3 4 2 2 3 3 4]
結果は4移行を無視(なんでも良いため) [0 1 2 3 4] となっていれば成功
```
