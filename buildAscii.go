package main

import (
	"fmt"
	"strings"
)

const tabSpaces = "  " // 2 spaces for \t

func BuildASCII(text string, banner_map map[rune][]string) (string, error) {

	if text == "" { // empty input
		return "", nil
	}

	// convert newlines into \n
	text = strings.ReplaceAll(text, "\r\n", "\n")

	// check the text exactly as typed, before anything gets replaced
	err := validateRawText(text)
	if err != nil {
		return "", err
	}

	//text = strings.ReplaceAll(text, `\n`, "\n")
	text = expandEscapes(text) // convert \n, \t and \\ into their correct chars

	// convert tabs \t into 2 spaces
	text = strings.ReplaceAll(text, "\t", tabSpaces)

	// split text by \n
	lines_split := strings.Split(text, "\n")

	// handling only newline inputs to remove any extra empty element
	if strings.ReplaceAll(text, "\n", "") == "" {
		lines_split = lines_split[:len(lines_split)-1]
	}

	// check chars before building the output
	err = validateChars(lines_split, banner_map)
	if err != nil {
		return "", err
	}

	// BUILDING THE OUTPUT
	var result strings.Builder
	for _, line := range lines_split { //row of the entire word (line = "Hi")

		if line == "" { // add a blank line for empty input lines
			result.WriteString("\n")
			continue // skip to next input line
		}

		for row := 0; row < 8; row++ { // loop through 8 rows (rows 0 to 7 of "Hi")

			for _, char := range line { // loop through each char ('H', then 'i')
				result.WriteString(banner_map[char][row]) // add current char row
			}

			result.WriteString("\n") // move to next row
		}
	}

	return result.String(), nil
}

// finds first unsupported character in the text as typed and points to it with ^
func validateRawText(text string) error {
	var shown strings.Builder // text as displayed
	col := 0                  // ^ column
	for _, char := range text {
		if char != '\t' && char != '\n' && (char < ' ' || char > '~') {
			return fmt.Errorf("%q shows unsupported character %q\n  %s%c\n  %s^", text, char, shown.String(), char, strings.Repeat(" ", col))
		}
		switch char {
		case '\t':
			shown.WriteString(`\t`)
			col += 2
		case '\n':
			shown.WriteString(`\n`)
			col += 2
		default:
			shown.WriteRune(char)
			col++
		}
	}
	return nil
}

// check all characters before building the output
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

// handles escape sequences after bash has processed the quotes
func expandEscapes(text string) string {
	var result strings.Builder

	for i := 0; i < len(text); i++ {

		if text[i] == '\\' && i+1 < len(text) { // check if current is \ and a next character exists

			switch text[i+1] { // check which character comes after \
			case 'n':
				result.WriteByte('\n') // \n or \\\n (odd \) -> newline
				i++
				continue

			case 't':
				result.WriteString(tabSpaces) // \t (odd \) -> 2 spaces
				i++
				continue

			case '\\': // each \\ pair becomes a single \
				result.WriteByte('\\') // \\ -> \    \\\\ -> \\
				i++
				continue

			case '"':
				result.WriteByte('"') // \" -> " (remove \ and keep ") ex: \"hi\" -> "hi"
				i++                   // skip the " we already handled
				continue              // move to next
			}
		}

		result.WriteByte(text[i]) // write normal characters as they are
	}

	return result.String() // return the new string
}

// single quotes '': bash changes nothing
// double quotes "": bash halves \\ first
