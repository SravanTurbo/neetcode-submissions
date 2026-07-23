func isValidSudoku(board [][]byte) bool {
    for i:=0; i<=8; i++ {
        rowMap := make(map[byte]bool)
        colMap := make(map[byte]bool)
        for j:=0; j<=8; j++{
            if string(board[i][j]) != "." {
                if _, ok := rowMap[board[i][j]]; ok{
                    return false
                }else{
                    rowMap[board[i][j]] = true
                }
            }

            if string(board[j][i]) != "." {
                if _, ok := colMap[board[j][i]]; ok{
                    return false
                }else{
                    colMap[board[j][i]] = true
                }
            }
        }
    }

    isValid9 := func(x, y [2]int) bool {
        gridMap := make(map[byte]bool)
        for i:=x[0]; i<=x[1]; i++{
            for j:=y[0]; j<=y[1]; j++{
                if string(board[i][j]) == "."{
                    continue
                }

                if _, ok := gridMap[board[i][j]]; ok{
                    return false
                }else{
                    gridMap[board[i][j]] = true
                }
            }
        }
        return true
    }

    limits := [][2]int{{0,2}, {3,5}, {6,8}}
    for i := range limits {
        for j := range limits {
            if !isValid9(limits[i], limits[j]) {
                return false
            }
        }
    }

    return true
}