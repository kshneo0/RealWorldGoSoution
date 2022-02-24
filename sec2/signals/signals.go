package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func CatchSig(ch chan os.Signal, done chan bool) {
	sig := <-ch
	fmt.Println("\nsigreceived:", sig)

	switch sig {
	case syscall.SIGINT:
		fmt.Println("handling a SIGINT now!")
	case syscall.SIGTERM:
		fmt.Println("handling a SIGTERM in an entirely differenr way!")
	default:
		fmt.Println("unexcepted signal received")
	}
	done <- true
}

func main() {
	signals := make(chan os.Signal)
	done := make(chan bool)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go CatchSig(signals, done)
	fmt.Println("Press ctrl-c to terminate...")
	<-done
	fmt.Println("Done!")
}
