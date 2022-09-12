package uniq

import (
	"os"
	"strconv"
	"strings"
)

func showDefault(flow *os.File, lines []string, registerNotImportant bool) {
	if len(lines) > 0 {
		flow.WriteString(lines[0] + "\n")
	}

	compareRule := func(left string, right string) bool {
		return left == right
	}

	if !registerNotImportant {
		compareRule = func(left string, right string) bool {
			return strings.EqualFold(left, right)
		}
	}

	for i := 1; i < len(lines)-1; i++ {
		prev := lines[i-1]
		cur := lines[i]

		if !compareRule(cur, prev) {
			flow.WriteString(cur + "\n")
		}
	}
}

type Node struct {
	key   string
	value int
}

func showAll(flow *os.File, lines []string, registerNotImportant bool) {
	sequence := []Node{}

	var countNodes int

	if len(lines) > 0 {
		sequence = append(sequence, Node{lines[0], 1})
		countNodes++
	}

	compareRule := func(left string, right string) bool {
		return left == right
	}

	if !registerNotImportant {
		compareRule = func(left string, right string) bool {
			return strings.EqualFold(left, right)
		}
	}

	for i := 1; i < len(lines); i++ {
		prev := lines[i-1]
		cur := lines[i]

		if !compareRule(cur, prev) {
			sequence = append(sequence, Node{cur, 1})
			countNodes++
		} else {
			sequence[countNodes-1].value++
		}
	}

	for _, node := range sequence {
		flow.WriteString(strconv.Itoa(node.value) + " " + node.key + "\n")
	}
}

func showUniq(flow *os.File, lines []string, registerNotImportant bool) {
	sequence := []Node{}

	var countNodes int

	if len(lines) > 0 {
		sequence = append(sequence, Node{lines[0], 1})
		countNodes++
	}

	compareRule := func(left string, right string) bool {
		return left == right
	}

	if !registerNotImportant {
		compareRule = func(left string, right string) bool {
			return strings.EqualFold(left, right)
		}
	}

	for i := 1; i < len(lines); i++ {
		prev := lines[i-1]
		cur := lines[i]

		if !compareRule(cur, prev) {
			sequence = append(sequence, Node{cur, 1})
			countNodes++
		} else {
			sequence[countNodes-1].value++
		}
	}

	for _, node := range sequence {
		if node.value == 1 {
			flow.WriteString(node.key + "\n")
		}
	}
}

func showUnUniq(flow *os.File, lines []string, registerNotImportant bool) {
	sequence := []Node{}

	var countNodes int

	if len(lines) > 0 {
		sequence = append(sequence, Node{lines[0], 1})
		countNodes++
	}

	compareRule := func(left string, right string) bool {
		return left == right
	}

	if !registerNotImportant {
		compareRule = func(left string, right string) bool {
			return strings.EqualFold(left, right)
		}
	}

	for i := 1; i < len(lines); i++ {
		prev := lines[i-1]
		cur := lines[i]

		if !compareRule(cur, prev) {
			sequence = append(sequence, Node{cur, 1})
			countNodes++
		} else {
			sequence[countNodes-1].value++
		}
	}

	for _, node := range sequence {
		if node.value != 1 {
			flow.WriteString(node.key + "\n")
		}
	}
}
