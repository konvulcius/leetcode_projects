package main

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i, str := range board {
		visited[i] = make([]bool, len(str))
	}
	searchingWord := []byte(word)
	fillingBox := make([]byte, 0, len(searchingWord))

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if backTrack(board, visited, searchingWord, &fillingBox, x, y, 0) {
				return true
			}
		}
	}
	return false
}

func backTrack(board [][]byte, visited [][]bool, word []byte, box *[]byte, x, y, cur int) bool {
	if len(*box) == len(word) {
		return true
	}
	if x < 0 || y < 0 || y >= len(board)  || x >= len(board[0]) || visited[y][x] {
		return false
	}

	if y < len(board) && x < len(board[y]) {
		if board[y][x] == word[cur] {
			*box = append(*box, word[cur])
			cur++
			visited[y][x] = true
			if backTrack(board, visited, word, box, x+1, y, cur) ||
				backTrack(board, visited, word, box, x-1, y, cur) ||
				backTrack(board, visited, word, box, x, y+1, cur) ||
				backTrack(board, visited, word, box, x, y-1, cur) {
				return true
			}
			*box = (*box)[:cur-1]
			visited[y][x] = false
		}
	}
	return false
}
