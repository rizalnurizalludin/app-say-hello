package main

import (
	"fmt"

	belajar_module "github.com/rizalnurizalludin/belajar-module-golang/v2"
)

func main() {
	var name, opr string
	var a,b int

	fmt.Println("Masukkan Nama")
	fmt.Scanln(&name)

	belajar_module.SayHello(name, "Rizal")

	fmt.Println("Masukkan Angka Pertama")
	fmt.Scanln(&a)
	
	fmt.Println("masukkan salah satu")
	fmt.Println("'+' , '-' , '/' , '*'")
	fmt.Scanln(&opr)
	
	fmt.Println("Masukkan Angka Kedua")
	fmt.Scanln(&b)
	
	result, _ := belajar_module.Calculate(a,b,opr)
	fmt.Println(result)
}
