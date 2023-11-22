# Daigonal Difference
正方行列が与えられたとき、その対角線の和の差の絶対値を計算する

例: 
```
1 2 3
4 5 6
9 8 9  
```

step1: 左から右への対角線＝1＋5＋9＝15。  
step2: 右から左への対角線 = 3 + 5 + 9 = 17.  
step3: その差の絶対値は｜15 - 17｜ = 2.  

Sample Input
```go
3
11 2 4
4 5 6
10 8 -12
```

step1:
```go
11
    5
        -12
(11 + 5 + -12) = 4
```

step2: 
```go
        4
    5
10
(4 + 5 + 10) = 19
```

step3:
```go
result := math.Abs(float64(4 - 19))
// result => 15
```


Sample Output
```go
15
```