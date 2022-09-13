package iohandlers

import "os"

func WriteFlow(stream *os.File, lines []string) {
	for _, val := range lines {
		stream.WriteString(val)
	}
}
