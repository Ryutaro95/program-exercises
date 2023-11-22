# Flipping Bits

32ビットの符号なし整数が与えられる。すべてのビットを(0->1, 1->0)に反転させて、その結果を符号なし整数として返す関数を作成する。

32ビット符号なし整数:
```
0 ~ 4294967295
```

作成する関数
```go
func flippingBits(n int64) int64 {
}
```

Sample Input
```go
3 // 与えられる32ビット整数の数
2147483647 
1 
0
```

Sample Output 
```go
2147483648 
4294967294 
4294967295
```

## 考え方

まずは、引数として与えられた整数を32ビット2進数の文字列へ変換する
```go
1 -> "00000000000000000000000000000001"
```

https://go.dev/play/p/gxsC_bMRHR4
```go

func main() {
    binary := decimalToBinaryString(1)
	fmt.Printf("型: %T: %s", binary, binary)
    //=> 型: string: 00000000000000000000000000000001
}
func decimalToBinaryString(decimal int64) string {
	if decimal == 0 {
		return "00000000000000000000000000000000"
	}
	binary := ""

	for decimal > 0 {
		remainder := decimal % 2
		binary = fmt.Sprintf("%d%s", remainder, binary)
		decimal /= 2
	}

	padding := 32 - len(binary)
	for i := 0; i < padding; i++ {
		binary = fmt.Sprintf("%d%s", 0, binary)
	}
	return binary
}
```

次に、上で変換した２進数を反転させる(0->1, 1->0)
```go
"00000000000000000000000000000001" -> "11111111111111111111111111111110"
```

```go
func main() {
	binary := decimalToBinaryString(1)

	converted := ""
	for _, bin := range binary {
		if bin == '0' {
			converted = fmt.Sprintf("%s%d", converted, 1)
		} else {
			converted = fmt.Sprintf("%s%d", converted, 0)
		}
	}

    fmt.Printf("型: %T: %s", converted, converted) 
    //=> 型: string: 11111111111111111111111111111110
}    
```

反転させた2進数を符号なし32ビットの整数に変換する
```go
"11111111111111111111111111111110" -> 4294967294
```

```go
func main() {
    converted := "11111111111111111111111111111110"
    result, err := strconvert.ParseUint(convert, 2, 32)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(int64(result))
    //=> 4294967294
}


```