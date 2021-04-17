package test

import (
	"fmt"
	"testing"

	"github.com/tap2joy/ChatService/utils"
)

func TestOnlineTimeFormat(t *testing.T) {
	duration1 := uint32(60)
	time1 := utils.FormatOnlineTime(duration1)
	fmt.Printf("duration %d: %s\n", duration1, time1)

	duration2 := uint32(677360)
	time2 := utils.FormatOnlineTime(duration2)
	fmt.Printf("duration %d: %s\n", duration2, time2)
}
