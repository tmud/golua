
package main

import (
	"fmt"
	"luajit"
)

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
	go test_c_goruntines("local function fib(n) print('1:'..n..';') if n < 2 then return n end return fib(n - 2) + fib(n - 1) end print(fib(20))")
	test_c_goruntines("local function fib(n) print('2:'..n..';') if n < 2 then return n end return fib(n - 2) + fib(n - 1) end print(fib(20))")
}
