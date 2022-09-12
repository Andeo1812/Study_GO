package uniq

import (
	"os"
	"strconv"
)

type outputRuleType func(*os.File, *[]string)

var showDefault outputRuleType = func(flow *os.File, lines *[]string) {
	countLines := len(*lines)

	if countLines > 0 {
		flow.WriteString((*lines)[0] + "\n")
	}

	for i := 1; i < countLines-1; i++ {
		prev := (*lines)[i-1]
		cur := (*lines)[i]

		if cur != prev {
			flow.WriteString(cur + "\n")
		}
	}
}

var showAll outputRuleType = func(flow *os.File, lines *[]string) {
	countLines := len(*lines)

	sequence := []string{}

	repetitionTable := map[string]int{}

	if countLines > 0 {
		sequence = append(sequence, (*lines)[0])

		repetitionTable[(*lines)[0]] = 1
	}

	for i := 1; i < countLines-1; i++ {
		prev := (*lines)[i-1]
		cur := (*lines)[i]

		if _, keyExist := repetitionTable[cur]; !keyExist {
			repetitionTable[cur] = 1
		} else {
			repetitionTable[cur] += 1
		}

		if cur != prev {
			sequence = append(sequence, cur)
		}
	}

	for _, val := range sequence {
		flow.WriteString(strconv.Itoa(repetitionTable[val]) + " " + val + "\n")
	}
}

var showUniq outputRuleType = func(flow *os.File, lines *[]string) {
	countLines := len(*lines)

	sequence := []string{}

	repetitionTable := map[string]int{}

	if countLines > 0 {
		sequence = append(sequence, (*lines)[0])

		repetitionTable[(*lines)[0]] = 1
	}

	for i := 1; i < countLines-1; i++ {
		prev := (*lines)[i-1]
		cur := (*lines)[i]

		if _, keyExist := repetitionTable[cur]; !keyExist {
			repetitionTable[cur] = 1
		} else {
			repetitionTable[cur] += 1
		}

		if cur != prev {
			sequence = append(sequence, cur)
		}
	}

	for _, val := range sequence {
		if repetitionTable[val] == 1 {
			flow.WriteString(val + "\n")
		}
	}
}

var showUnUniq outputRuleType = func(flow *os.File, lines *[]string) {
	countLines := len(*lines)

	sequence := []string{}

	repetitionTable := map[string]int{}

	if countLines > 0 {
		sequence = append(sequence, (*lines)[0])

		repetitionTable[(*lines)[0]] = 1
	}

	for i := 1; i < countLines-1; i++ {
		prev := (*lines)[i-1]
		cur := (*lines)[i]

		if _, keyExist := repetitionTable[cur]; !keyExist {
			repetitionTable[cur] = 1
		} else {
			repetitionTable[cur] += 1
		}

		if cur != prev {
			sequence = append(sequence, cur)
		}
	}

	for _, val := range sequence {
		if repetitionTable[val] != 1 {
			flow.WriteString(val + "\n")
		}
	}
}
