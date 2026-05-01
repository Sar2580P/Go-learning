package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// ctx.Done will trigger first
func doWork(ctx context.Context, wg *sync.WaitGroup){
	select {
	case <-time.After(2*time.Second):
		fmt.Println("WorkDone")
	case <-ctx.Done():
		fmt.Println("Canceled:", ctx.Err())
	}
	wg.Done()
}


func main(){
	ctx, cancel:= context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() // cancel the context to release resources when done with main


	// start a goroutine to simulate processing
	var wg sync.WaitGroup
	wg.Add(1)
	go doWork(ctx, &wg)
	
	wg.Wait()

	// EXPECTED: ctx would get done, so processing would be ignored and func returns
}