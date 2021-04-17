package test

import (
	"fmt"
	"testing"

	"github.com/tap2joy/ChatService/utils"
)

func TestBuildTrieTree(t *testing.T) {
	tree := utils.NewTree()
	tree.BuildTree("Hello every one, I'm go")
	tree.AddWord("simple")
	fmt.Printf("%v", tree)
}

func TestReplace(t *testing.T) {
	tree := utils.NewTree()
	tree.BuildTree("fuck god")

	str1 := tree.Replace("test 1: fuck", "*")
	fmt.Printf("str1 = %v\n", str1)

	str2 := tree.Replace("test 2: aafuckbb", "*")
	fmt.Printf("str2 = %v\n", str2)
}
