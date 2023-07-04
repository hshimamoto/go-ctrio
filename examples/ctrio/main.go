// MIT License Copyright (C) 2023 Hiroshi Shimamoto
package main

import (
	"io"
	"log"
	"os"

	"github.com/hshimamoto/go-ctrio"
)

func main() {
	if len(os.Args) < 3 {
		log.Println("ctrio <enc|dec> key")
		os.Exit(1)
	}
	cmd := os.Args[1]
	key := []byte(os.Args[2])
	switch cmd {
	case "enc":
		w, err := ctrio.NewWriter(key, os.Stdout)
		if err != nil {
			log.Printf("NewWriter: %v", err)
			os.Exit(16)
		}
		_, err = io.Copy(w, os.Stdin)
		if err != nil {
			log.Printf("Copy: %v", err)
			os.Exit(17)
		}
	case "dec":
		r, err := ctrio.NewReader(key, os.Stdin)
		if err != nil {
			log.Printf("NewReader: %v", err)
			os.Exit(32)
		}
		_, err = io.Copy(os.Stdout, r)
		if err != nil {
			log.Printf("Copy: %v", err)
			os.Exit(33)
		}
	default:
		log.Printf("unknown command %v", cmd)
		os.Exit(48)
	}
	return
}
