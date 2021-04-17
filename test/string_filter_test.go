package test

import (
	"fmt"
	"testing"

	"github.com/tap2joy/ChatService/utils"
)

func TestStringFilter(t *testing.T) {
	fmt.Println(">>>>> StringFilter test begin")

	filter := utils.NewStringFilter()
	before := "goo Dxxxg afuck god hell wwwankerww bunny fucker"
	after := filter.Replace(before)
	fmt.Printf("before replace: %s\n", before)
	fmt.Printf(" after replace: %s\n", after)
}
