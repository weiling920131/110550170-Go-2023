package main

import (
	"fmt"
	"math/big"
	"strconv"
	"syscall/js"
)

func CheckPrime(this js.Value, args []js.Value) interface{} {
	// TODO: Check if the number is prime
	str := js.Global().Get("value").Get("value").String()
	num, _ := strconv.ParseInt(str, 10, 64)
	
	ans := big.NewInt(num)
	isPrime := ans.ProbablyPrime(0)

	if (isPrime){
		js.Global().Get("answer").Set("innerText", "It's prime")
	}else {
		js.Global().Get("answer").Set("innerText", "It's not prime")
	}
	
	return str
}

func registerCallbacks() {
	// TODO: Register the function CheckPrime
	js.Global().Set("CheckPrime", js.FuncOf(CheckPrime))
}

func main() {
	fmt.Println("Golang main function executed")
	registerCallbacks()

	//need block the main thread forever
	select {}
}
