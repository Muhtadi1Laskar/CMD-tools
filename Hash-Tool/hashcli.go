package main

import (
	"flag"
	"fmt"
	"hash-tool/hashs"
)

func main() {
	data := flag.String("data", "", "the data to hash")
	hashName := flag.String("hash", "", "the hash function name")
	flag.Parse()

	hash, err := hashs.HashData(*data, *hashName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", hash)
}