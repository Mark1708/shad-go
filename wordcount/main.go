//go:build !solution

package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	wc := make(map[string]int) // words count

	for _, filePath := range os.Args[1:] {
		file, err := os.Open(filePath)
		check(err)

		fs := bufio.NewScanner(file)
		fs.Split(bufio.ScanLines)

		for fs.Scan() {
			word := fs.Text()
			wc[word]++
		}
		check(file.Close())
	}

	for word, count := range wc {
		if count < 2 {
			continue
		}
		fmt.Printf("%v\t%v\n", count, word)
	}
}
