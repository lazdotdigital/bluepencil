package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	b, path, err := readFileFromArgs()
	if err != nil {
		log.Fatalf("couldn't read file from args: %v", err)
	}
	bd := bindData{
		editor: newEditor(b),
		path:   path,
	}
	initUI(bd)
}
