package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"unicode/utf8"
)

func main() {
	var filename string
	var output string

	flag.StringVar(&filename, "file", "", "Filename to make utf-8 compatible.")
	flag.StringVar(&output, "out", "out.txt", "Where to store the output.")
	flag.Parse()

	if _, err := os.Stat(filename); err != nil {
		log.Fatal(err)
	}

	fHandle, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fHandle.Close()

	ofHandle, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}
	defer ofHandle.Close()

	fmt.Println("Converting...")

	chunk := make([]byte, 10240)
	for {
		_, err := fHandle.Read(chunk)
		if err == io.EOF {
			break
		}

		utf8line := make([]byte, 10240)
		if !utf8.Valid(chunk) {
			for _, cbyte := range chunk {
				if utf8.Valid([]byte{cbyte}) && cbyte != 0 {
					utf8line = append(utf8line, cbyte)
				}
			}
		} else {
			utf8line = chunk
		}
		ofHandle.Write(utf8line)
	}
	fmt.Printf("Finished! The UTF-8 compatible version of the file is stored at \"%s\".\n", output)
}
