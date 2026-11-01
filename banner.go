package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// loadBanner reads the banner file and returns a map:
// each character -> its 8 lines of art.
func loadBanner(name string) (map[rune][]string, error) {
	// only accept the three banner types
	if name != "standard" && name != "shadow" && name != "thinkertoy" {
		return nil, fmt.Errorf("unknown banner %q", name)
	}

	// read the whole file
	path := filepath.Join("banners", name+".txt")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read banner file %q: %v", path, err)
	}

	// fix Windows line endings, then split into lines
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(content, "\n")

	// 95 characters x 9 lines = 855 lines minimum
	if len(lines) < 855 {
		return nil, fmt.Errorf("banner file %q is incomplete: %d lines, need 855", path, len(lines))
	}

	// one map entry for every character from ' ' (32) to '~' (126)
	bannerMap := make(map[rune][]string)
	for char := ' '; char <= '~'; char++ {
		// first art line of this character
		start := int(char-' ')*9 + 1
		// take its 8 lines
		bannerMap[char] = lines[start : start+8]
	}

	return bannerMap, nil
}



