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

	fmt.Println("\n*** Working with Channel Synchronization ***")
	done := make(chan bool, 1)
	go channels.Worker(done)

	/* Block until we receive a notification from the worker on the channel
	If removed, the program would exit before the worker even started.
	*/
	<-done

	//go_beta.WorkWithGoBeta()
}
