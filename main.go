package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/lazdotdigital/bluepencil/editor"
)

func main() {
	b, path, err := readFileFromArgs()
	if err != nil {
		log.Fatalf("couldn't read file from args: %v", err)
	}
	bd := bindData{
		Editor: editor.New(b),
		path:   path,
	}
	initUI(bd)
}
