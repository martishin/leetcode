package word_search_ii_212

import "fmt"

type TrieNodeM struct {
	children [26]*TrieNodeM
	word     string
}

var nbsM = [4][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func buildTrieM(words []string) *TrieNodeM {
	head := &TrieNodeM{}
	for _, word := range words {
		node := head
		for _, ch := range word {
			idx := ch - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &TrieNodeM{}
			}
			node = node.children[idx]
		}
		node.word = word
	}
	return head
}

func findWordsM(board [][]byte, words []string) []string {
	trieHead := buildTrieM(words)
	result := []string{}

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			idx := board[row][col] - 'a'
			if trieHead.children[idx] != nil {
				dfsM(board, row, col, trieHead, &result)
			}
		}
	}

	return result
}

func dfsM(board [][]byte, row, col int, node *TrieNodeM, result *[]string) {
	ch := board[row][col]
	idx := ch - 'a'
	node = node.children[idx]

	if node == nil {
		return
	}

	if node.word != "" {
		*result = append(*result, node.word)
		node.word = "" // Avoid duplicate results
	}

	board[row][col] = '#' // mark as visited

	for _, d := range nbsM {
		newRow := row + d[0]
		newCol := col + d[1]
		if newRow >= 0 && newRow < len(board) && newCol >= 0 && newCol < len(board[0]) && board[newRow][newCol] != '#' {
			dfsM(board, newRow, newCol, node, result)
		}
	}

	board[row][col] = ch // restore
}

func TestM() {
	fmt.Println(findWordsM([][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}, []string{"oath", "pea", "eat", "rain"}))

	fmt.Println(findWordsM([][]byte{
		{'a', 'b'},
		{'c', 'd'},
	}, []string{"abcb"}))
}
