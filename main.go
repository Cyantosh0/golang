package main

import (
	"fmt"

	"github.com/Cyantosh0/golang/channels"
	"github.com/Cyantosh0/golang/enums"
	"github.com/Cyantosh0/golang/interfaces"
	//go_beta "github.com/Cyantosh0/golang/go_1.18"
)

func main() {
	// maps.Maps()
	interfaces.Interfaces()
	// example()
}

func example() {
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

	fmt.Println("\n*** Working with Channel Directions ***")
	firstChannel := make(chan string, 1)
	secondChannel := make(chan string, 1)
	channels.SendMessageToChannel(firstChannel, "passed message")
	channels.RecieveAndSendMessageWithInChannels(secondChannel, firstChannel)
	fmt.Println(<-secondChannel)

	fmt.Println("\n*** Working with Channel Select ***")
	channels.SelectChannel()

	fmt.Println("\n*** Working with Timeouts ***")
	channels.ChannelTimeout()

	fmt.Println("\n*** Working with Range Over Channels ***")
	channels.RangeOverChannels()

	fmt.Println("\n*** Working with Worker Pools ***")
	channels.WorkerPools()

	fmt.Println("\n*** Working with Wait Group ***")
	channels.WaitGroup()

	//go_beta.WorkWithGoBeta()
}
