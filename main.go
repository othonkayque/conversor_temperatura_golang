package main

import "fmt"

func main() {
	fmt.Println("programa iniciado!")
	CtoF(-17.777779)
	FtoC(0)
	
}

// Converter celsius para fahrenheit
func CtoF(celsius float32) float32 {
	var resultado float32 = (celsius * 1.8) + 32
	fmt.Println(celsius, "Graus celsius são ", resultado, "Graus fahrenheit")
	return resultado
}
// Converter fahrenheit para celsius  
func FtoC(fahr float32) float32 {
	var resultado float32 = (fahr - 32)* 5/9
	fmt.Println(fahr, "fahrenheit são ", resultado, "celsius")
	return resultado
}
