package main

import (
	"fmt"

	"github.com/Cyantosh0/golang/channels"
	"github.com/Cyantosh0/golang/enums"
	//go_beta "github.com/Cyantosh0/golang/go_1.18"
)

func main() {
	fmt.Println("*** Working with Enums ***")
	enums.Mon.DisplayWeekDayDetails()

	fmt.Println("\n*** Working with unbuffered channels ***")
	channels.UnbufferedChannels()

	fmt.Println("\n*** Working with buffered channels ***")
	channels.BufferedChannels()

	//go_beta.WorkWithGoBeta()
}
