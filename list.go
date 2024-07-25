package main

import (
	"fmt"
	"log"
)

func ListPackages() {
	output, err := ShellExec(PACMAN, "-Q")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)

}
