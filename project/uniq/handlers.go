package uniq

import (
	"os"
	"strconv"
	"strings"
)

type outputRuleType func(*os.File, []string, bool)

var showDefault outputRuleType = func(flow *os.File, lines []string, registerNotImportant bool) {
	if len(lines) > 0 {
		flow.WriteString(lines[0] + "\n")
	}

	for i := 1; i < len(lines)-1; i++ {
		prev := lines[i-1]
		cur := lines[i]

		if !registerNotImportant {
			if cur != prev {
				flow.WriteString(cur + "\n")
			}
		} else {
			if !strings.EqualFold(cur, prev) {
				flow.WriteString(cur + "\n")
			}
		}
	}
}

type Node struct {
	key   string
	value int
}

var showAll outputRuleType = func(flow *os.File, lines []string, registerNotImportant bool) {
	sequence := []Node{}

	var countNodes int

	if len(lines) > 0 {
		sequence = append(sequence, Node{lines[0], 1})
		countNodes++
	}

	for i := 1; i < len(lines); i++ {
		prev := lines[i-1]
		cur := lines[i]

		if !registerNotImportant {
			if cur != prev {
				sequence = append(sequence, Node{cur, 1})
				countNodes++
			} else {
				sequence[countNodes-1].value++
			}
		} else {
			if !strings.EqualFold(cur, prev) {
				sequence = append(sequence, Node{cur, 1})
				countNodes++
			} else {
				sequence[countNodes-1].value++
			}
		}
	}

	for _, node := range sequence {
		flow.WriteString(strconv.Itoa(node.value) + " " + node.key + "\n")
	}
}

var showUniq outputRuleType = func(flow *os.File, lines []string, registerNotImportant bool) {
	sequence := []Node{}

	var countNodes int

	if len(lines) > 0 {
		sequence = append(sequence, Node{lines[0], 1})
		countNodes++
	}

	for i := 1; i < len(lines); i++ {
		prev := lines[i-1]
		cur := lines[i]

		if !registerNotImportant {
			if cur != prev {
				sequence = append(sequence, Node{cur, 1})
				countNodes++
			} else {
				sequence[countNodes-1].value++
			}
		} else {
			if !strings.EqualFold(cur, prev) {
				sequence = append(sequence, Node{cur, 1})
				countNodes++
			} else {
				sequence[countNodes-1].value++
			}
		}
	}

	for _, node := range sequence {
		if node.value == 1 {
			flow.WriteString(node.key + "\n")
		}
	}
}

var showUnUniq outputRuleType = func(flow *os.File, lines []string, registerNotImportant bool) {
	countLines := len(lines)

	sequence := []Node{}

	var countNodes int

	if countLines > 0 {
		sequence = append(sequence, Node{lines[0], 1})
		countNodes++
	}

	for i := 1; i < countLines; i++ {
		prev := lines[i-1]
		cur := lines[i]

		if !registerNotImportant {
			if cur != prev {
				sequence = append(sequence, Node{cur, 1})
				countNodes++
			} else {
				sequence[countNodes-1].value++
			}
		} else {
			if !strings.EqualFold(cur, prev) {
				sequence = append(sequence, Node{cur, 1})
				countNodes++
			} else {
				sequence[countNodes-1].value++
			}
		}
	}

	for _, node := range sequence {
		if node.value != 1 {
			flow.WriteString(node.key + "\n")
		}
	}
}
