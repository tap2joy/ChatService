package utils

import (
	"bufio"
	"fmt"
	"os"
)

type StringFilter struct {
	TrieTree *TrieTree
}

func NewStringFilter() *StringFilter {
	filter := &StringFilter{
		TrieTree: NewTree(),
	}
	err := filter.LoadFile()
	if err != nil {
		fmt.Printf("load filter.txt failed, err = %v\n", err)
	}
	fmt.Println("create string filter success")
	return filter
}

func (filter *StringFilter) LoadFile() error {
	dir, _ := os.Getwd()

	fullpath := fmt.Sprintf("%s/config/filter.txt", dir)
	// fptr := flag.String("filterpath", dir+"/config/filter.txt", "the file path to read from")
	// flag.Parse()

	f, err := os.Open(fullpath)
	if err != nil {
		return err
	}

	defer f.Close()

	i := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		i++
		word := s.Text()
		//fmt.Printf("%d: %s\n", i, word)
		filter.TrieTree.AddWord(word)
	}

	err = s.Err()
	if err != nil {
		return err
	}

	fmt.Printf("load %s/config/filter.txt success\n", dir)

	return nil
}

func (filter *StringFilter) Replace(content string) string {
	return filter.TrieTree.Replace(content, "*")
}
