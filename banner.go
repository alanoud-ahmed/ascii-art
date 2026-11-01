package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// loadBanner --> reads the banner file and return map of eah char to its 8 lones of arts
func loadBanner(name string) (map[rune][]string, err) {
	// only accept the three banner types
	if name != "standard" && name != "shadow" && name != "thinkertoy" {
		return nil, fmt.Errorf("unkown banner %q ", name)
	}

	// read the whole file
	data, err := os.ReadFile(filepath.Join("banners", name+".txt"))
	if err != nil {
		return nil, fmt.Errorf(" can not find file ")
	}

}
