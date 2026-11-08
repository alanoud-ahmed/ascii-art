
# bash tests/double_quotes.sh
echo $'\033[1;31mDOUBLE QUOTES\n\n\033[0m'

# NEWLINES
echo "NEWLINES"
go run . "\\n" | cat -n -e           # expected: 1 blank line ($)
go run . "\\\\n" | cat -n -e         # expected: \n (backslash + n)
go run . "\\\\\\n" | cat -n -e       # expected: \ (8 rows), then 1 blank line ($) -> 9 rows
go run . "\\\\\\\\n" | cat -n -e     # expected: \\n (2 backslashes + n)
go run . "\\\\\\\\\n" | cat -n -e     # expected: \\ (8 rows), then 1 blank line ($) -> 9 rows

# TABS
echo $'\n\nTABS'
go run . "\\t" | cat -n -e           # expected: 2 spaces
go run . "\\\\t" | cat -n -e         # expected: \t (backslash + t)
go run . "\\\\\\t" | cat -n -e       # expected: \ followed by 2 spaces
go run . "\\\\\\\\t" | cat -n -e     # expected: \\t (2 backslashes + t)

# BACKSLASHES
echo $'\n\nBACKSLASHES'
go run . "\\" | cat -n -e           # expected: \ 1 backslash
go run . "\\\\" | cat -n -e         # expected: \ 1 backslash
go run . "\\\\\\" | cat -n -e       # expected: \\ 2 backslashes
go run . "\\\\\\\\" | cat -n -e     # expected: \\ 2 backslashes

# QUOTATION MARKS
echo $'\n\nQUOTES'
go run . "" | cat -n -e              # expected: no output (empty argument)
go run . "\"\"" | cat -n -e          # expected: "" (2 quotation marks)
go run . "\"hi\"" | cat -n -e        # expected: "hi", quotes included
go run . "\"\"hi\"\"" | cat -n -e    # expected: ""hi""
go run . "\\\"hi\\\"" | cat -n -e    # expected: "hi"
go run . "hi\\\"hi\\\"hi" | cat -n -e # expected: hi"hi"hi

# EMPTY LINES
echo $'\n\nEMPTY LINES'
go run . "" | cat -n -e                 # expected: no output
go run . "\\n\\n" | cat -n -e           # expected: 2 blank lines ($ on lines 1 and 2)
go run . "\\n\\n\\n" | cat -n -e        # expected: 3 blank lines ($ on lines 1, 2, and 3)
go run . "Hello\\n" | cat -n -e         # expected: Hello (8 rows) + 1 blank line ($ on line 9)
go run . "\\nHello" | cat -n -e         # expected: 1 blank line ($) + Hello (8 rows)
go run . "Hello\\n\\nWorld" | cat -n -e # expected: Hello (8 rows) + 1 blank line + World (8 rows)

# SPACES
echo $'\n\nSPACES'
go run . " " | cat -n -e               # expected: 8 blank rows for 1 space
go run . "   " | cat -n -e             # expected: 8 blank rows for 3 spaces
go run . "Hello   World" | cat -n -e   # expected: Hello + 3 spaces + World

# ACTUAL NEWLINES AND TABS
echo $'\n\nACTUAL NEWLINES AND TABS'
go run . $'Hello\nWorld' | cat -n -e    # expected: Hello + World on separate lines, 16 rows
go run . $'Hello\tWorld' | cat -n -e    # expected: Hello + 2 spaces + World, 8 rows

# INVALID CHARACTERS
echo $'\n\nINVALID CHARACTERS'
go run . "Hello😊" | cat -n -e    # expected: error for unsupported character '😊'
go run . "Helloمرحبا" | cat -n -e # expected: error for unsupported character 'م'

# UNKNOWN ESCAPES
echo $'\n\nUNKNOWN ESCAPES'
go run . "\\x" | cat -n -e # expected: \x (backslash + x)
go run . "\\r" | cat -n -e # expected: \r (backslash + r)
go run . "\\q" | cat -n -e # expected: \q (backslash + q)

echo $'\n\n'
go run . "\x" | cat -n -e # expected: \x (backslash + x)
go run . "\r" | cat -n -e # expected: \r (backslash + r)
go run . "\q" | cat -n -e # expected: \q (backslash + q)
