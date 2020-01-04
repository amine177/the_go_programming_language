package main


import (
	"fmt"
	"os"
)


func main() {
	var s string
	var i int

	for i, s = range os.Args[1:] {
		fmt.Printf("Args[%v] = %v\n", i, s)
	}
}
