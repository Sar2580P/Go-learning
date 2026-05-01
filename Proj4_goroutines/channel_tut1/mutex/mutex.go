package main

import (
	"fmt"
	"sync"
)

type post struct{
	view int
	mu sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup){
	// keep unlock inside defer, else it will cause deadlock if the function panics before unlocking the mutex
	defer func(){
		p.mu.Unlock()     //  unlock the mutex after incrementing the view count
		wg.Done()
	}()

	p.mu.Lock()  // lock the mutex to ensure exclusive access to the view count
	p.view++
}


func main(){
	var wg sync.WaitGroup
	myPost:= post{view: 0}

	for i:=0; i<100; i++{
		wg.Add(1)
		go myPost.inc(&wg)
	}
	wg.Wait()
	fmt.Println("Post views:", myPost.view)
}