package main //This is a special package name that tells the compiler to look for the entry function within this file

import "fmt" //When importing a package it will show an error right away because GO doesn't allow unused imports in the codebase

func main(){
	fmt.Println("Testing")
	var intNum int32 = 327673
	fmt.Println(intNum)

	var floatNum1 float32 = 1234567.89	//32 bit floats and 64 bit floats
	var floatNum2 float64 = 1234567.89 //Its important to choose the data type to use less resources
	fmt.Println(floatNum1)
	fmt.Println(floatNum2)
}