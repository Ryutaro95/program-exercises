# sparse-arrays
入力文字列のコレクションとクエリー文字列のコレクションの２つを受け取る。  
各クエリー文字列について、それが入力文字列のリストで何回出現するかを調べて結果を配列で返す。



## Sample
* Input1
    ```
    3
    def
    de
    fgh
    3
    de
    lmn
    fgh
    ```
* Output1
    ```
    1
    0
    1
    ```
* Input2
    ```
    4
    aba
    baba
    aba
    xzxb
    3
    aba
    xzxb
    ab
    ```
* Output2
    ```
    2
    1
    0
    ```
* Input3
    ```
    13
    abcde
    sdaklfj
    asdjf
    na
    basdn
    sdaklfj
    asdjf
    na
    asdjf
    na
    basdn
    sdaklfj
    asdjf
    5
    abcde
    sdaklfj
    asdjf
    na
    basdn
    ```
* Output
    ```
    1
    3
    4
    3
    2
    ```