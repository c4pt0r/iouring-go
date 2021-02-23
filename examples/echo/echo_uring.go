package main

import (
	"fmt"
	"log"

	"github.com/hodgesds/iouring-go"
)

func ServeUring() {
	ops := []iouring.RingOption{
		iouring.WithEnterErrHandler(func(err error) {
			log.Println(err)
		}),
	}
	if *debug {
		ops = append(ops, iouring.WithDebug())
	}
	r, err := iouring.New(
		8192,
		&iouring.Params{
			Features: iouring.FeatNoDrop | iouring.FeatFastPoll,
		},
		ops...,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Listen on:", *addr)
	l, err := r.SockoptListener(
		"tcp",
		*addr,
		func(err error) {
			log.Println(err)
		},
		iouring.SOReuseport,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("hello", conn)
		go func() {
			msg := make([]byte, 5)
			for {
				_, err := conn.Read(msg)
				if err != nil {
					fmt.Println("!!!!!", conn.RemoteAddr())
					fmt.Println("!!!!!", err)
					return
				}
				_, err = conn.Write(msg)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}()
	}

}
