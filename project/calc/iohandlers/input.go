package iohandlers

import (
	"bufio"
	"os"
	"strings"
)

func Input() string {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	noSpacesExpr := strings.ReplaceAll(scanner.Text(), " ", "")

	return noSpacesExpr
}
