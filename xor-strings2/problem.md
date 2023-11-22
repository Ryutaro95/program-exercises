# XOR strings2

Input:
```go
10101
00101
```

Output:
```go
10000
```

## XOR
どちらかが true の場合にのみ true を返す
```
| A | B | A XOR B | bool  |
|---|---|---------|-------|
| 0 | 0 |    0    | false |
| 0 | 1 |    1    | true  |
| 1 | 0 |    1    | true  |
| 1 | 1 |    0    | false |
```