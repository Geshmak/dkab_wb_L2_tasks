package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var timeout = flag.Duration("timeout", 10*time.Second, "timeout")
	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		fmt.Println("need 2 parametrs")
		return
	}
	var host = args[0]
	var port = args[1]

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), *timeout)
	if err != nil {
		time.Sleep(*timeout)
		log.Fatalf("go-telnet: %s", err)
	}

	c := make(chan os.Signal, 1)
	errch := make(chan error, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go req(conn, errch, c)
	go resp(conn, errch, c)

	select {
	case <-c:
		conn.Close()
	case err = <-errch:
		if err != nil {
			log.Fatalf("go-telnet: %s", err)
		}
	}
}

func req(conn net.Conn, errch chan<- error, c chan<- os.Signal) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				c <- syscall.Signal(syscall.SIGQUIT)
				return
			}
			errch <- err
		}

		fmt.Fprintf(conn, text+"\n")
	}
}

func resp(conn net.Conn, errch chan<- error, c chan<- os.Signal) {
	for {
		reader := bufio.NewReader(conn)
		text, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				c <- syscall.Signal(syscall.SIGQUIT)
				return
			}
			errch <- err
		}

		fmt.Print(text)
	}
}
