package main

import (
	"fmt"
	_"math/rand"
	"time"
)


// recieving data using channel
func processNum(numChan chan int) {

	// fmt.Println("Processing number", <-numChan)

	for num := range numChan{
		fmt.Println("Processing number", num)
		time.Sleep(500 * time.Millisecond)
	}
}

// sending data to go routine using channel
func sum(result chan int, num1, num2 int) {
	total := num1 + num2
	result <- total
}


// go-routine synchronization using channels
func task(done chan bool){
	defer func() {done <- true}()  // signal completion when function exits
	fmt.Println("processing...")
	// done <- true  // never reached if process fails
}

// <-chan string means the channel is only for receiving data
// chan<- bool means the channel is only for sending data
func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() {done <- true}()  // signal completion when function exits
	for email := range emailChan {
		fmt.Println("Sending email to", email)
		time.Sleep(100 * time.Millisecond)  // simulate time taken to send email
	}
}


func main() {

	/* ☀️ Deadlock example */  // will cause a deadlock because the main goroutine is trying to send a message to the channel, but there is no other goroutine receiving from it.
	// messageChan := make(chan string)
	// messageChan <- "ping"  // blocking

	// msg := <-messageChan
	// fmt.Println(msg)

	/* ☀️ Sending data to go routine using channel */
	// numChan:= make(chan int)
	// go processNum(numChan)
	
	// for {
	// 	numChan <- rand.Intn(100)
	// }


	/* ☀️ Recieving data using channel */
	// resultChan := make(chan int)
	// go sum(resultChan, 5, 10)
	// res:= <-resultChan
	// fmt.Println("Result:", res)   // blocking


	/* ☀️ Signaling completion with channels */
	// doneChan := make(chan bool)
	// go task(doneChan)
	// <- doneChan  // blocking: wait for task to signal completion


	/* ☀️ Buffered channels */
	// emailChan := make(chan string, 100)  // buffered channel with capacity of 100
	// doneChan := make(chan bool)
	// go emailSender(emailChan, doneChan)

	// for i:=0; i<5; i++{
	// 	emailChan <- fmt.Sprintf("%d@example.com", i)
	// }

	// // need to close the buffer channels
	// close(emailChan)  // AVOID DEADLOCK: close channel to signal no more emails will be sent
	// <- doneChan  // wait for emailSender to finish


	/* ☀️ Multiple Channels */
	chan1:= make(chan int)
	chan2:= make(chan string)
	
	go func() {   // inline go-routine
		chan1<-10   // will not cause a deadlock because there is another goroutine receiving from chan1
	}()

	go func() {
		chan2<-"pong"
	}()

	for i:=0; i<2; i++{    // multiple channels
		select {
		case chan1Val := <-chan1:
			fmt.Println("recieved data from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("recieved data from chan2", chan2Val)
		}
	}
}