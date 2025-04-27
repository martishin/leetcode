package word_search_ii_212

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
	word     string
}

type Position struct {
	row int
	col int
}

func buildTrie(words []string) *TrieNode {
	head := &TrieNode{}

	for _, word := range words {
		current := head
		for _, ch := range word {
			idx := ch - 'a'
			if current.children[idx] == nil {
				current.children[idx] = &TrieNode{}
			}
			current = current.children[idx]
		}
		current.word = word
	}

	return head
}

var nbs = [4]Position{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func dfs(board [][]byte, pos Position, trieNode *TrieNode, visited map[Position]struct{}, result *[]string) {
	if trieNode.word != "" {
		*result = append(*result, trieNode.word)
		trieNode.word = ""
	}

	for _, nb := range nbs {
		nbPos := Position{row: pos.row + nb.row, col: pos.col + nb.col}

		if nbPos.row >= 0 && nbPos.row < len(board) && nbPos.col >= 0 && nbPos.col < len(board[0]) {
			_, ok := visited[nbPos]
			idx := board[nbPos.row][nbPos.col] - 'a'
			nbTrieNode := trieNode.children[idx]

			if ok || nbTrieNode == nil {
				continue
			}

			visited[nbPos] = struct{}{}
			dfs(board, nbPos, nbTrieNode, visited, result)
			delete(visited, nbPos)
		}
	}
}

func findWords(board [][]byte, words []string) []string {
	trieHead := buildTrie(words)

	result := []string{}
	for r, row := range board {
		for c, ch := range row {
			idx := ch - 'a'
			trieNode := trieHead.children[idx]
			if trieNode != nil {
				pos := Position{row: r, col: c}
				visited := make(map[Position]struct{})
				visited[pos] = struct{}{}
				dfs(board, pos, trieNode, visited, &result)
			}
		}
	}

	return result
}

func Test() {
	fmt.Println(findWords([][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}, []string{"oath", "pea", "eat", "rain"}))

	fmt.Println(findWords([][]byte{
		{'a', 'b'},
		{'c', 'd'},
	}, []string{"abcb"}))
}
