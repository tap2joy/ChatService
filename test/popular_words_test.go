package test

import (
	"fmt"
	"testing"

	"github.com/tap2joy/ChatService/utils"
)

func TestPopularWords(t *testing.T) {
	fmt.Println("TestPopularWords begin")
	pop := utils.NewPopularWords()
	pop.AddLog("pytest-django tries to automatically add your project to the Python path by looking for a manage.py file and adding its path to the Python path.")
	topWord := pop.GetPopularWord()
	fmt.Printf("popular word is: %s\n", topWord)
	pop.Reset()
	topWord = pop.GetPopularWord()
	fmt.Printf("after reset popular word is: %s\n", topWord)
}
