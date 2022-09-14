package iohandlers

import (
	"fmt"
	"os"
)

func Output(number float64) {
	os.Stdout.WriteString(fmt.Sprintf("%.2f", number) + "\n")
}
