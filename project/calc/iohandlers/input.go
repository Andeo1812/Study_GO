package iohandlers

import (
	"bufio"
	"os"
)

func Input() string {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	return scanner.Text()
}
