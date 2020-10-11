package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Skal åpne en fil ...")
	f, err := os.Create("sample.txt")
	if err != nil {
		fmt.Println("Kunne ikke lage filen. Panikk!")
		panic(err)
	}
	byte_array := []byte{1, 2, 3, 4, 5}
	g, err := f.Write(byte_array)
	if err != nil {
		fmt.Println("Kunne ikke skrive til denne filen. Panikk!")
		panic(err)
	}
	fmt.Printf("Skre %d bytes til filen\n", g)
}
