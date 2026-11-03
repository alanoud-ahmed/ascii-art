package main

import (
	"fmt"
	"os"
)

func main() {
	// can accept text only OR text + banner name
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Fprintln(os.Stderr, "usage : go run . \"text\" [standard | shadow | thinkertoy]")
		os.Exit(1)
	}

	// saving the args
	text := os.Args[1]
	bannerName := "standard" // as a default choice
	if len(os.Args) == 3 {
		bannerName = os.Args[2]
	}

	// 1) loading the banner file
	banner, err := loadBanner(bannerName)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	// build & print the ASCII art
	ascii, err := BuildASCII(text, banner)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	fmt.Print(ascii) // print it

}
