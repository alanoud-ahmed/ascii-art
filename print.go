package main

import (
	"fmt"
	"strings"
)

const tabSpaces = "  " // 2 spaces for \t

func PrintASCII(text string, banner_map map[rune][]string) error {

	if text == "" { // empty input
		return nil
	}

	// convert newlines into \n
	text = strings.ReplaceAll(text, "\r\n", "\n")
	text = strings.ReplaceAll(text, "\n", "\\n")

	// convert tabs \t into 2 spaces
	text = strings.ReplaceAll(text, "\\t", tabSpaces)
	text = strings.ReplaceAll(text, "\t", tabSpaces)

	// split text by \n
	lines_split := strings.Split(text, "\\n")

	// handling only newline inputs to remove any extra empty element
	if strings.ReplaceAll(text, "\\n", "") == "" {
		lines_split = lines_split[:len(lines_split)-1]
	}

	// check chars before printing
	err := validateChars(lines_split, banner_map)
	if err != nil {
		return err
	}

	// PRINTING
	for _, line := range lines_split { //row of the entire word (line = "Hi")

		if line == "" { // print a blank line for empty input lines
			fmt.Println()
			continue // skip to next input line
		}

		for row := 0; row < 8; row++ { // loop through 8 rows (rows 0 to 7 of "Hi")

			for _, char := range line { // loop through each char ('H', then 'i')
				fmt.Print(banner_map[char][row]) // print current char row
			}

			fmt.Println() // move to next row
		}
	}

	return nil
}

// check all characters before printing
func validateChars(lines_split []string, banner_map map[rune][]string) error {

	for line_index, line := range lines_split { // loop through each input line

		column := 0
		for _, char := range line { // loop through each character
			column++ // count character position, starting from 1

			// check ASCII range
			if char < ' ' || char > '~' {
				return fmt.Errorf("unsupported character %q at line %d, column %d", char, line_index+1, column)
			}

			// check char exists in banner map
			rows, ok := banner_map[char]
			if ok == false {
				return fmt.Errorf("missing character %q in banner at line %d, column %d", char, line_index+1, column)
			}

			// check character has exactly 8 rows
			if len(rows) != 8 {
				return fmt.Errorf("character %q must have 8 rows at line %d, column %d", char, line_index+1, column)
			}
		}
	}

	return nil
}
