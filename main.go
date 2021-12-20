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
	editor := editor.New(b)
	executor := makeExecutor(editor)
	addons, err := loadAddonsFromEnv(executor)
	if err != nil {
		log.Fatalf("couldn't load addons: %v", err)
	}
	bd := bindData{
		Editor: editor,
		path:   path,
		addons: addons,
	}
	if err := initUI(bd); err != nil {
		log.Fatalf("couldn't initialize web view: %v", err)
	}
}
