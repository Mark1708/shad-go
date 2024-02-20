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
	fp := os.Args[1:]          // file's paths
	wc := make(map[string]int) // words count
	var w []string

	for _, filePath := range fp {
		file, err := os.Open(filePath)
		check(err)

		fs := bufio.NewScanner(file)
		fs.Split(bufio.ScanLines)

		for fs.Scan() {
			row := fs.Text()
			if count, found := wc[row]; found {
				nr := count + 1
				wc[row] = nr
				if nr == 2 {
					w = append(w, row)
				}
			} else {
				wc[row] = 1
			}
		}
		check(file.Close())
	}

	for _, word := range w {
		fmt.Printf("%d\t%s\n", wc[word], word)
	}
}
