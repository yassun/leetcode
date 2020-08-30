func convert(s string, numRows int) string {
	slen := len(s)

	// 処理不要
	if slen == 0 || numRows == 0 {
		return ""
	}
	if numRows == 1 {
		return s
	}

	resRows := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		resRows[i] = make([]byte, slen/2+1)
	}

	colIdx, rowIdx := 0, 0
	writeIdx := 1
	for i := 0; i < slen; i++ {
		resRows[rowIdx][colIdx] = s[i]

		// 正のカウント中で一番下までいったら反転 || 負のカウントで一番上まで行ったら反転
		if (writeIdx == 1 && rowIdx == numRows-1) || (writeIdx == -1 && rowIdx == 0) {
			writeIdx = -1 * writeIdx
		}
		rowIdx += writeIdx

		// 上方向への書き込み中だけcolを移動
		if writeIdx < 0 {
			colIdx++
		}
	}

	// 結果の作成
	result := make([]byte, 0, slen)
	for i := 0; i < numRows; i++ {
		for j, _ := range resRows[i] {
			if resRows[i][j] != 0 {
				result = append(result, resRows[i][j])
			}
		}
	}

	return string(result)
}
