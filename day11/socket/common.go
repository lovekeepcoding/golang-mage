package socket

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var sigChan = make(chan os.Signal, 10)

func DealSigPipe() {
	signal.Notify(sigChan, syscall.SIGPIPE)
	for {
		select {
		case sig := <-sigChan:
			fmt.Printf("receive signal %d\n", sig)
		}
	}
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal error: %s\n", err.Error()) //stdout是行缓冲的，他的输出会放在一个buffer里面，只有到换行的时候，才会输出到屏幕。而stderr是无缓冲的，会直接输出
		os.Exit(1)
	}
}

type (
	Request struct {
		A int
		B int
	}
	Response struct {
		Sum int
	}
)
