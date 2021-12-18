package main

import (
	"log"
)

func main() {
	b, path, err := readFileFromArgs()
	if err != nil {
		log.Fatalf("couldn't read file from args: %v", err)
	}
	bd := bindData{
		newEditor(b),
	}
	initUI(path, bd)
}
