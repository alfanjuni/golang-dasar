package main

import "fmt"

func main() {
	var hai *string
	hai = nil
	if hai != nil && *hai == "hello" {
		fmt.Println("BENAR")
	} else {
		fmt.Println("SALAH")
	}

}
