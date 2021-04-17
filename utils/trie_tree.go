package utils

import "strings"

type TrieTree struct {
	Root *TrieNode
}

type TrieNode struct {
	Children map[string]*TrieNode
	IsEnd    bool
}

func NewTree() *TrieTree {
	return &TrieTree{Root: NewNode()}
}

func NewNode() *TrieNode {
	node := &TrieNode{Children: make(map[string]*TrieNode), IsEnd: false}
	return node
}

// 根据文字内容构建字典树
func (tree *TrieTree) BuildTree(content string) {
	trimContent := strings.Trim(content, " ")
	words := strings.Split(trimContent, " ")
	for _, word := range words {
		tree.AddWord(word)
	}
}

// 往字典树中添加单词
func (tree *TrieTree) AddWord(word string) {
	tree.Root.addWord(word)
}

// 往指定节点添加单词
func (node *TrieNode) addWord(word string) {
	characters := []rune(word)
	character := string(characters[0])

	var isEnd bool
	var restWord string
	if len(characters) == 1 {
		isEnd = true
		restWord = ""
	} else {
		isEnd = false
		restWord = string(characters[1:])
	}

	if node.Children == nil {
		node.Children = make(map[string]*TrieNode)
	}

	if _, ok := node.Children[character]; !ok {
		node.Children[character] = &TrieNode{nil, isEnd}
	} else {
		if isEnd {
			node.Children[character].IsEnd = isEnd
		}
	}

	if len(restWord) > 0 {
		node.Children[character].addWord(restWord)
	}
}

func (node *TrieNode) isEnd() bool {
	return node.IsEnd
}

func (tree *TrieTree) Replace(text string, character string) string {
	var (
		parent  = tree.Root
		current *TrieNode
		runes   = []rune(text)
		length  = len(runes)
		left    = 0
		found   bool
	)

	for position := 0; position < len(runes); position++ {
		current, found = parent.Children[string(runes[position])]

		if !found || (!current.isEnd() && position == length-1) {
			parent = tree.Root
			position = left
			left++
			continue
		}

		if current.isEnd() && left <= position {
			for i := left; i <= position; i++ {
				runes[i] = []rune(character)[0]
			}
		}

		parent = current
	}

	return string(runes)
}
