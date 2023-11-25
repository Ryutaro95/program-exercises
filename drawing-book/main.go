package main

func pageCount(pages, targetPage int32) int32 {
	// 本の始まりから目的のページまでのページ数を計算
	// 目的のページが偶数なら targetPage / 2 で目的のページまでたどり着くめくり枚数が得られる
	// 奇数の場合も、割り算の結果が整数として切り捨てられるため、上と同様の結果が得られる
	fromFront := targetPage / 2

	// 本の終わりから目的のページまでのページ数を計算
	// もし、本の1ページ目が右側ではなく左からの場合は、以下がFromFrontになる
	// (本のページ総数 / 2) - fromFront で本の終わりから目的のページまで辿り着くのめくり枚数が得られる
	fromBack := (pages / 2) - fromFront

	// 最終的な目的ページまでのめくり枚数を格納する
	var turns int32
	// 本の始まりまたは、終わりどちらからめくるのが最短か比較する
	if fromFront < fromBack {
		turns = fromFront
	} else {
		turns = fromBack
	}
	return turns
}

func main() {
}
