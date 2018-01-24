package main

import "fmt"

func main1() {
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a // a copy of a is assigned to b
	fmt.Printf("%p \n", &a)
	fmt.Printf("%p \n", &b)
	b[0] = "Singapore"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)
}
func main2() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] //creates a slice from a[1] to a[3]
	fmt.Printf("%p \n", &a[1])
	fmt.Printf("%p \n", b)
	fmt.Println(b)
}
func main3() {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:]
	fmt.Printf("%p \n", &fruitarray[1])
	fmt.Printf("%p \n", fruitslice[0])
	fruitslice = append(fruitslice, "xxx", "yyy")
	fmt.Printf("%p \n", fruitslice[0])
	fmt.Printf("length of slice %d capacity %d ", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
	fmt.Println()
}
func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}
func main4() {
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}
func main() {
	a := []byte("f")
	fmt.Println(len(a), cap(a))
	b := append(a, []byte("bar")...)
	c := append(a, []byte("baz")...)

	fmt.Println(string(a), string(b), string(c))
}
