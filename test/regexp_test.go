package test

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	src := "/status ssss"
	match, _ := regexp.MatchString("^/status*", src)
	if match {
		fmt.Println("match")
	} else {
		fmt.Println("not match")
	}
}
