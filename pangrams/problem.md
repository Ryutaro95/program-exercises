# Pamgrams
パングラムとはアルファベットの全ての文字を含む文字列のこと。  
文章が与えられたとき、それが英語のアルファベットのパングラムであるかどうかを判定する。  
大文字、小文字は無視して、パングラムであるかどうかを返す。


Example:
```go
s := "The quick brown fox jumps over the lazy dog"
```
この文字列は英語のアルファベットをすべて含むので、パングラムであると判断できるため `"pangram"` という文字列を返す。  
もし、この文字列がパングラムではないと判定された場合は `"not pangram"` という文字列を返す。


Sample Input:
```go
"We promptly judged antique ivory buckles for the next prize"
```

Sample Output:
```go
"pangram"
```
上の文字列には全ての文字列が含まれているためパングラムとなる

Sample Input2:
```go
"We promptly judged antique ivory buckles for the prize"
```

Sample Output2:
```go
"not pangram"
```
上の文字列には `"x"` が含まれていないので、 パングラムではないとなる

## a..z までのアルファベットを作成する
```go
func main() {
    alphabets := []string{}
    for str := 'a'; str <= 'z'; str++ {
        alphabets = append(alphabets, string(str))
    }

    fmt.Println(alphabets)
}
// [a b c d e f g h i j k l m n o p q r s t u v w x y z]
```