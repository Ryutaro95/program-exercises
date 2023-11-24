# Remove Element
整数配列 nums と整数 val を指定すると、nums 内の val の出現箇所をすべてその場で削除します。要素の順序は変更される場合があります。次に、val と等しくない要素の数を nums で返します。  
val に等しくない nums 内の要素の数が k であると考えてください。受け入れられるようにするには、次のことを行う必要があります。

* nums の最初の k 要素に val と等しくない要素が含まれるように配列 nums を変更します。 nums の残りの要素は、nums のサイズと同様に重要ではありません。
* return k

```go
func removeElement(nusm []int, val int) int
```

審査員は次のコードを使用してソリューションをテストします。
```c
int[] nums = [...]; // Input array
int val = ...; // Value to remove
int[] expectedNums = [...]; // The expected answer with correct length.
                            // It is sorted with no values equaling val.

int k = removeElement(nums, val); // Calls your implementation

assert k == expectedNums.length;
sort(nums, 0, k); // Sort the first k elements of nums
for (int i = 0; i < actualLength; i++) {
    assert nums[i] == expectedNums[i];
}

```

## Example
example1:
```
Input: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]
```

example2:
```
Input: nums = [0,1,2,2,3,0,4,2], val = 2
Output: 5, nums = [0,1,4,0,3,_,_,_]
```

## Approach
### Two Pointer
* Time complexity: O(n)
* Space complexity: O(1)