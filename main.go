package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

var (
	flagFile = flag.String("file", "", "(Required) File from which to read newline-separated items to be sorted")
)

func main() {
	flag.Parse()

	var b []byte
	var err error
	b, err = os.ReadFile(*flagFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}
	s := strings.TrimSpace(string(b))

	items := strings.Split(s, "\n")
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })

	sort.Slice(items, func(i, j int) bool {
		fmt.Fprintln(os.Stderr, "----------")
		fmt.Fprintf(os.Stderr, "1 | %s\n2 | %s\n", items[i], items[j])
		for {
			fmt.Print("Choice: ")
			var choice string
			if _, err := fmt.Scanln(&choice); err != nil {
				fmt.Fprintf(os.Stderr, "error scanning input: %v\n", err)
			}
			if choice == "1" {
				return true
			}
			if choice == "2" {
				return false
			}
			fmt.Fprintf(os.Stderr, `Not a valid choice "%s", must be "1" or "2"\n`, choice)
		}
	})

	fmt.Println("==========")
	for _, item := range items {
		fmt.Println(item)
	}
}
