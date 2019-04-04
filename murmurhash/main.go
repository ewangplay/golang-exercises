package main

import (
	"log"

	"github.com/go-ego/murmur"
)

func main() {
	var str = "github.com1"

	sum32 := murmur.Sum32(str)
	log.Println("hash32...", sum32)

	sum32 = murmur.Sum32(str, 0)
	log.Println("hash32...", sum32)

	hash32 := murmur.Murmur3([]byte(str))
	log.Println("hash32...", hash32)

	hash32 = murmur.Murmur3([]byte(str), 1)
	log.Println("hash32...", hash32)
}
