package main

import "fmt"

import "rsc.io/quote"

func main(){
	fmt.Println("Hello world...")

	fmt.Println("Printing Inspirational Quote from an imported quote package function:");
	fmt.Println(quote.Opt());


}