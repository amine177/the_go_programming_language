package main


import (
	"fmt"
	"os"
	"strings"
	"time"
)


func main() {
	var s = "" 
	var sep = " "

	
	t := time.Now().UnixNano()
	for _, arg := range os.Args[1:] {
		s += arg + sep
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(time.Now().UnixNano() - t)

	t = time.Now().UnixNano()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Now().UnixNano() - t)
}
