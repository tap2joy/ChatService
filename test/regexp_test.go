package test

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	fmt.Println(">>>>> Regexp test begin")

	src := "/status ssss"
	match, _ := regexp.MatchString("^/status*", src)
	if match {
		fmt.Println("regexp match")
	} else {
		fmt.Println("regexp not match")
	}
}
