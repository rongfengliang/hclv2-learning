package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var wait sync.WaitGroup
	ticker := time.NewTicker(time.Second * 5)
	wait.Add(1)
	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("some thing")
			}
		}
	}()
	wait.Wait()
}
