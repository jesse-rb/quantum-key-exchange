package main

import (
	"log"
	"os"
)

var info = log.New(os.Stdout, "\033[35minfo\t->\033[0m ", log.Lshortfile)
var error = log.New(os.Stderr, "\033[31merror\t->\033[0m ", log.Lshortfile)

func init() {
	info.Println("Hi...")
}

func main() {
	error.Println("Nothing here yet!")
}