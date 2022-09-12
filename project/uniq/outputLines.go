package uniq

import (
	"os"
	"strconv"
)

var simple outputRuleType = func(flow *os.File, table *map[string]int) {
	for key, val := range *table {
		flow.WriteString(key + " " + strconv.Itoa(val) + "\n")
	}
}

type outputRuleType func(*os.File, *map[string]int)

func showLines(path string, outputRule outputRuleType, table *map[string]int) {
	var stream *os.File = os.Stdout

	if path != "" {
		stream, errCreate := os.Create(path)
		if errCreate != nil {
			panic(errCreate)
		}

		defer stream.Close()

		outputRule(stream, table)

		return
	}

	outputRule(stream, table)

}
