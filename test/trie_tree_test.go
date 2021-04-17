package test

import (
	"fmt"
	"testing"

	"github.com/tap2joy/ChatService/utils"
)

func TestTrieTree(t *testing.T) {
	fmt.Println(">>>>> TrieTree test begin")
	tree := utils.NewTree()
	tree.BuildTree("fuck god")

	str1 := tree.Replace("test 1: fuck", "*")
	fmt.Printf("str1 = %v\n", str1)

	str2 := tree.Replace("test 2: aafuckbb", "*")
	fmt.Printf("str2 = %v\n", str2)
}
