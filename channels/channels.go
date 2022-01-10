package channels

import "fmt"

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
