package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func echo(conn net.Conn, shout string, delay time.Duration)  {
	fmt.Fprintln(conn, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", strings.ToLower(shout))
}

func handle(conn net.Conn) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		echo(conn, input.Text(), 1 * time.Second)
	}
	conn.Close()
}
