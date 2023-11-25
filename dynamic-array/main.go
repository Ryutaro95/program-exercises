package main

func dynamicArray(n int32, queries [][]int32) []int32 {
	// var arr [][]int32
	// var lastAnswer int32
	// var answer []int32
	// var idx int32
	//
	// for i := 0; i < n; i++ {
	// 	arr[i] = int32{}
	// }
	// for _, query := range queries {
	// 	q, x, y := query[0], query[1], query[2]
	// 	idx := (x ^ lastAlastAnswer) % n
	//
	// 	if q == 1 {
	// 		arr[idx] = y
	// 	}
	//
	// }
	// for _, query := range queries {
	// 	idx = (query[1] ^ lastAnswer) % n
	// 	y := query[2]
	// 	if query[0] == 1 {
	// 		arr[idx] = y
	// 	} else {
	// 		lastAnswer = arr[idx][y%len(arr[idx])]
	// 		answer = append(answer, lastAnswer)
	// 	}
	// }

	return []int32{}
}

/*
* Queryは2種類あり、文字列の配列として与えられるので、それを解析する
* 1. Qeury: 1 x y
*   1. idx = ((x ^ lastAnswer) % n) とする
*   2. 整数 y を arr[idx] に加える
* 2. Query: 2 x y
*   1. idx = ((x ^ lastAnswer) % n)とする
*   2. arr[idx][y % size(arr[idx])] の値をlastAnswerに代入する
*   3. lastAnswer の新しい値を answer配列に格納する
*
* * 関数の説明
* n: arr[idx] を初期化する空の配列の数
* string: queries[q]: 空白で区切られた3つの整数を含むクエリー文字列
* Returns
*   int[]: 各タイプの2つのクエリ結果
*
* Input
* 2 5
* 1 0 5
* 1 1 7
* 1 0 3
* 2 1 0
* 2 1 1
* Output
* 7
* 3
*
* n = 2
* lastAnswer = 0
* arr[0] = []
* arr[1] = []
* arr[(0 ^ 0) % 2]に5を加える = arr[0]
*
* lastAnswer = 0
* arr [0] = [5]
* arr [1] = []
* arr[(1 ^ 0) % 2]に7を加える = arr[1]
*
* lastAnswer = 0
* arr[0] = [5]
* arr[1] = [7]
* arr[(0 ^ 0) % 2]に3を加える = arr[0]
*
* lastAnswer = 0
* arr[0] = [5, 3]
* arr[1] = [7]
* arr[(1 ^ 0) % 2] = arr[1] のインデックス0の値を lastAnswer に代入して、lastAnswerを表示する
*
* lastAnswer = 7
* arr[0] = [5, 3]
* arr[1] = [7]
* arr[(1 ^ 7) % 2] = arr[0] のインデックス1の値を lastAnswerに代入して、lastAnswerを表示する
 */

func main() {}
