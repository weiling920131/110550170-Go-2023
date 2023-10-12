package main

import "fmt"
import "strconv"

func main() {
	var n int64

	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	result := Sum(n)
	fmt.Println(result)
}

func Sum(n int64) string {
	// TODO: Finish this function
	var i int64
	var value int64 = 1
	var s = "1"
	for i=2;i<=n;i++{
		if i%7 == 0{
			continue
		}
		value = value + i
		s = s +"+"+strconv.Itoa(int(i))
	}
	s = s + "=" + strconv.Itoa(int(value))
	return s
}
