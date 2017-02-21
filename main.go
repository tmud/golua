
package main

import (
	"fmt"
	"luajit"
)

func main() {

	L := luajit.NewState()
	defer L.Close()

	L.OpenLibs()

	lines := []string{
	   "local function fib(n) if n < 2 then return n end  return fib(n - 2) + fib(n - 1) end print(fib(35))",
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
