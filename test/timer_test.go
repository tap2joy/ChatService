package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/tap2joy/ChatService/utils"
)

func TestTimer(t *testing.T) {
	utils.StartTimer(5, "2021-01-01 19:14:30", "", func() {
		fmt.Println("Hello", time.Now())
	})
	select {}
}
