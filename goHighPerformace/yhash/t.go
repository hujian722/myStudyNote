package main

import (
	"fmt"
	"log"
)

func main() {
	log.SetFlags(log.Llongfile)
	f3()
}
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}
func f3() {
	i := 1
	defer fmt.Printf("1:: %v\n", i)
	i = 2
	defer fmt.Printf("2:: %v\n", i)
	i = 8
	defer fmt.Printf("3:: %v\n", i)
	fmt.Printf("4:: %v\n", i)
}
