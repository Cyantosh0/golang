package channels

import (
	"fmt"
	"time"
)

func UnbufferedChannels() {
	messages := make(chan string)

	go func() { messages <- "apple" }()

	msg := <-messages

	fmt.Println("Message : ", msg)
}

func BufferedChannels() {
	messages := make(chan string, 2)

	messages <- "apple"
	messages <- "banana"

	fmt.Println(<-messages) // Prints : apple
	fmt.Println(<-messages) // Prints : banana
	// fmt.Println(<-messages) // fatal error: all goroutines are asleep - deadlock!
}

func Worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true // Send a value to notify that we’re done
}

func SendMessageToChannel(firstChannel chan<- string, msg string) {
	firstChannel <- msg
}

func RecieveAndSendMessageWithInChannels(secondChannel chan<- string, firstChannel <-chan string) {
	msg := <-firstChannel
	secondChannel <- msg
}

func SelectChannel() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(3 * time.Second)
		c2 <- "two"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c3 <- "three"
	}()

	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case msg3 := <-c3:
			fmt.Println("received", msg3)
		}
	}
}

func ChannelTimeout() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "This is channel 1"
	}()

	select {
	case msg := <-c1:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout for channel 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "This is channel 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout for channel 2")
	}
}

// Its possible to close a non-empty channel but still have the remaining values be received
func RangeOverChannels() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
