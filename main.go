package main

import (
	"fmt"

	helper "github.com/rizalnurizalludin/belajar-module-golang/v2"
)

func main() {
	var name string

	fmt.Println("Masukkan Nama")
	fmt.Scanln(&name)

	helper.SayHello(name, "Rizal")
}
