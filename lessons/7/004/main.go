package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

type workerList struct {
	ch <-chan int
	// stop chan struct{}
	stop <-chan os.Signal
	wg   sync.WaitGroup
}

func (w *workerList) Run() {
	for i := 0; i < 10; i++ {
		w.wg.Add(1)
		go func(a int) {
			defer w.wg.Done()
			for {
				select {
				case t := <-w.ch:
					fmt.Println("worker", a, ":", t)
				case <-w.stop:
					fmt.Println("exiting", a)
					return
				}
			}
		}(i)
	}
}

func (w *workerList) Stop() {
	// close(w.stop)
	w.wg.Wait()
}

func main() {
	someChan := make(chan int)
	notify := make(chan os.Signal)
	signal.Notify(notify, os.Interrupt)
	w := &workerList{
		ch:   someChan,
		stop: notify,
	}
	w.Run()
	someChan <- 42
	w.Stop()
	// time.Sleep(50 * time.Nanosecond)
}
