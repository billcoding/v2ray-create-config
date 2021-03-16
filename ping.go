package main

import (
	"fmt"
	"net"
	"time"
)

func telnet(host string, port int) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), time.Second*5)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_ = conn.Close()
	}
	return err == nil
}
