package main

import "time"

func main() {
	go func() {
		Loop()
	}()

	for {
		time.Sleep(time.Second * 10)
	}
}
