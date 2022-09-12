package uniq

import (
	"os"
	"strconv"
)

type outputRuleType func(*os.File, *map[string]int)

var showDefault outputRuleType = func(flow *os.File, table *map[string]int) {
	for key, _ := range *table {
		flow.WriteString(key + "\n")
	}
}

var showAll outputRuleType = func(flow *os.File, table *map[string]int) {
	for key, val := range *table {
		flow.WriteString(strconv.Itoa(val) + " " + key + "\n")
	}
}

var showUniq outputRuleType = func(flow *os.File, table *map[string]int) {
	for key, val := range *table {
		if val == 1 {
			flow.WriteString(key + "\n")
		}
	}
}

var showUnUniq outputRuleType = func(flow *os.File, table *map[string]int) {
	for key, val := range *table {
		if val != 1 {
			flow.WriteString(key + "\n")
		}
	}
}

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
