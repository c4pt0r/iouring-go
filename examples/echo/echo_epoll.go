package main

import "net"

func ServeEpoll() {
	if lsn, err := net.Listen("tcp", *addr); err == nil {
		for {
			conn, err := lsn.Accept()
			if err != nil {
				break
			}
			go func() {
				defer conn.Close()
				msg := make([]byte, *testMsgLen)
				for {
					_, err := conn.Read(msg)
					if err != nil {
						break
					}
					_, err = conn.Write(msg)
					if err != nil {
						break
					}
				}
			}()
		}
	}
}
