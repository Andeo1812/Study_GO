package calc

import (
	"bufio"
	"os"
)

func input() string {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	return scanner.Text()
}

func output(line string) {
	os.Stdout.WriteString(line)
}
