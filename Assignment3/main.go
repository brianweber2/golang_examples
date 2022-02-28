package main

import (
	"fmt"
	"os"
)

func main() {
	a := os.Args
	filename := a[1]

	fileContents, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error when opening file", filename, err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	fileContents.Read(bs)
	fmt.Println(string(bs))

	fileContents.Close()
}
