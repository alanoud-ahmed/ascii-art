package main

import (
	"strings"
	"testing"
)

// Normal inputs: check the number of printed lines.
func TestPrintASCII(t *testing.T) {
	banner, err := loadBanner("standard")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name  string
		input string
		lines int // expected number of output lines
	}{
		{"empty", "", 0},
		{"one newline", `\n`, 1},
		{"two newlines", `\n\n`, 2},
		{"real newline", "\n", 1},
		{"word", "Hello", 8},
		{"word and newline", `Hello\n`, 9},
		{"two words", `Hello\nThere`, 16},
		{"empty line between", `Hello\n\nThere`, 17},
		{"real newline between", "Hello\nThere", 16},
		{"spaces and numbers", "1Hello 2There", 8},
		{"special characters", `{Hello & There #}`, 8},
		{"tab", "a\tb", 8},
		{"written tab", `a\tb`, 8},
		{"windows newline", "Hello\r\nThere", 16},
		{"quote and backslash", `1a"#FdwHywR&/()=\`, 8},
	}

	for _, tc := range tests {
		got, err := PrintASCII(tc.input, banner)
		if err != nil {
			t.Errorf("%s: unexpected error: %v", tc.name, err)
			continue
		}
		if n := strings.Count(got, "\n"); n != tc.lines {
			t.Errorf("%s: got %d lines, want %d", tc.name, n, tc.lines)
		}
	}
}

// All 8 rows of one line of text must have the same width (output stays aligned).
func TestRowsHaveSameWidth(t *testing.T) {
	banner, err := loadBanner("standard")
	if err != nil {
		t.Fatal(err)
	}

	got, _ := PrintASCII("Hello There", banner)
	rows := strings.Split(strings.TrimSuffix(got, "\n"), "\n")

	for i, row := range rows {
		if len(row) != len(rows[0]) {
			t.Errorf("row %d has width %d, row 0 has width %d", i, len(row), len(rows[0]))
		}
	}
}

// Every printable ASCII character must work with every banner.
func TestAllPrintableASCII(t *testing.T) {
	for _, name := range []string{"standard", "shadow", "thinkertoy"} {
		banner, err := loadBanner(name)
		if err != nil {
			t.Fatal(err)
		}

		var input strings.Builder
		for char := ' '; char <= '~'; char++ {
			input.WriteRune(char)
		}

		if _, err := PrintASCII(input.String(), banner); err != nil {
			t.Errorf("%s: %v", name, err)
		}
	}
}

// Bad characters: error and NO output at all (even if the bad one is on line 2).
func TestInvalidInput(t *testing.T) {
	banner, err := loadBanner("standard")
	if err != nil {
		t.Fatal(err)
	}

	inputs := []string{"مرحبا", "😀", "Hi\x00", "Hello\nمرحبا", `Hello\nمرحبا`}
	for _, input := range inputs {
		got, err := PrintASCII(input, banner)
		if err == nil || got != "" {
			t.Errorf("%q: expected error and no output, got %q, %v", input, got, err)
		}
	}
}

// validateRawText: valid text passes, and the error points to the bad character.
func TestValidateRawText(t *testing.T) {
	if err := validateRawText("Hello\\nThere\t!"); err != nil {
		t.Errorf("valid text rejected: %v", err)
	}

	err := validateRawText("Hi م")
	if err == nil {
		t.Fatal("expected an error for a non-ASCII character")
	}
	// "Hi " is 3 columns, so the ^ comes after 2 (indent) + 3 spaces
	if !strings.HasSuffix(err.Error(), "     ^") {
		t.Errorf("^ is not under the bad character: %q", err.Error())
	}
}

// A broken banner map must give an error, not a crash.
func TestBrokenBannerMap(t *testing.T) {
	// 'A' is missing from the map
	if _, err := PrintASCII("A", map[rune][]string{}); err == nil {
		t.Error("expected an error for a missing character")
	}

	// 'A' has only 3 rows instead of 8
	short := map[rune][]string{'A': {"a", "b", "c"}}
	if _, err := PrintASCII("A", short); err == nil {
		t.Error("expected an error for a wrong number of rows")
	}
}

// Wrong banner names must give an error.
func TestLoadBannerErrors(t *testing.T) {
	for _, name := range []string{"", "nothing", "../standard", "Standard"} {
		if _, err := loadBanner(name); err == nil {
			t.Errorf("loadBanner(%q): expected an error", name)
		}
	}
}