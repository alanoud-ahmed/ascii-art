package main
 
import (
"strings"
"testing"
)
 
// Test 1: normal inputs must print the right number of lines.
func TestValidInput(t *testing.T) {
banner, err := loadBanner("standard")
if err != nil {
t.Fatal(err)
}
 
tests := []struct {
input string
lines int // expected number of lines
}{
{"", 0},
{`\n`, 1},
{"Hello", 8},
{`Hello\n`, 9},
{`Hello\nThere`, 16},
{`Hello\n\nThere`, 17},
{"1Hello 2There", 8},
{"{Hello & There #}", 8},
}
 
for _, test := range tests {
result, err := BuildASCII(test.input, banner)
if err != nil {
t.Errorf("%q: unexpected error: %v", test.input, err)
continue
}
 
got := strings.Count(result, "\n")
if got != test.lines {
t.Errorf("%q: got %d lines, expected %d", test.input, got, test.lines)
}
}
}
 
// Test 2: unsupported characters must return an error and print nothing.
func TestInvalidInput(t *testing.T) {
banner, err := loadBanner("standard")
if err != nil {
t.Fatal(err)
}
 
inputs := []string{"مرحبا", "😀", `Hello\nمرحبا`}
 
for _, input := range inputs {
result, err := BuildASCII(input, banner)
if err == nil {
t.Errorf("%q: expected an error", input)
}
if result != "" {
t.Errorf("%q: expected no output", input)
}
}
}
 
// Test 3: a wrong banner name must return an error.
func TestWrongBannerName(t *testing.T) {
_, err := loadBanner("nothing")
if err == nil {
t.Error("expected an error for a wrong banner name")
}
}