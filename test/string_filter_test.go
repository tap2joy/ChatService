package test

import (
	"fmt"
	"testing"

	"github.com/tap2joy/ChatService/utils"
)

func TestLoadFile(t *testing.T) {
	filter := utils.NewStringFilter()
	filter.LoadFile()
}

func TestStringFilter(t *testing.T) {
	filter := utils.NewStringFilter()
	filter.LoadFile()
	before := "goo Dxxxg afuck god hell wwwankerww bunny fucker"
	after := filter.Replace(before)
	fmt.Printf("before replace: %s\n", before)
	fmt.Printf(" after replace: %s\n", after)
}
