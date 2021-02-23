package main

import (
	"flag"
)

var (
	testMsgLen = flag.Int("l", 26, "test message length")
	addr       = flag.String("addr", "0.0.0.0:12345", "Listen address, default: 0.0.0.0:12345")
	testType   = flag.String("t", "epoll", "test type (epoll, iouring)")

	debug = flag.Bool("d", false, "debug mode")
)

func main() {
	flag.Parse()
	if *testType == "epoll" {
		ServeEpoll()
	} else if *testType == "iouring" {
		ServeUring()
	}
}
