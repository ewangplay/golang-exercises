package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ewangplay/rwriter"
)

func main() {
	var err error
	var w io.Writer

	cfg := &rwriter.Config{
		Module: "kitty",
		Path:   ".",
	}
	w, err = rwriter.NewRotateWriter(cfg)
	if err != nil {
		fmt.Printf("Create rotate writer failed: %v\n", err)
		os.Exit(1)
	}

	log.SetOutput(w)
	msg := "Hello, rotate writer!"
	log.Println(msg)

	filename := "kitty.log"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Read log file failed: %v\n", err)
		os.Exit(1)
	}
	if !strings.Contains(string(content), msg) {
		fmt.Printf("The log file %s should contains '%s'", filename, msg)
		os.Exit(1)
	}
	os.Remove(filename)
}
