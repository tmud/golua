package main

import (
	"fmt"
	"io"
	"log"
	"luajit"
	"net"
	"shutdown"
	"time"
)

const version = "(dev)"

func test_c_goruntines(code string) {
	L := luajit.NewState()
	defer L.Close()

	L.OpenLibs()

	lines := []string{
		code,
	}
	var err error

	for _, line := range lines {
		err = L.LoadString(line)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}
		err = L.Run()
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}
	}
}

func main() {
	fmt.Println("Starting server " + version)
	// start network server
	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		return
	}
	go handleListener(listener)
	shutdown.Wait()
	//go test_c_goruntines("local function fib(n) print('1:'..n..';') if n < 2 then return n end return fib(n - 2) + fib(n - 1) end print(fib(20))")
	//test_c_goruntines("local function fib(n) print('2:'..n..';') if n < 2 then return n end return fib(n - 2) + fib(n - 1) end print(fib(20))")
}

func handleListener(listener net.Listener) error {
	defer listener.Close()
	for {
		connection, err := listener.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()
	for {
		_, err := io.WriteString(connection, time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
