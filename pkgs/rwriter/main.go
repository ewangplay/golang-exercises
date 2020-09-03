package main

import (
	"io"
	"log"

	"github.com/ewangplay/rwriter"
)

func main() {
	var err error
	var w io.Writer

	cfg := &rwriter.RotateWriterConfig{
		Module: "test",
		Path:   ".",
	}
	w, err = rwriter.NewRotateWriter(cfg)
	if err != nil {
		log.Printf("Create rotate writer failed: %v\n", err)
		return
	}

	log.SetOutput(w)
	log.SetPrefix("HelloKitty ")
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("Hello, rotate writer!")
}
