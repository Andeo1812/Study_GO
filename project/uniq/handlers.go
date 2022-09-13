package uniq

import (
	"strconv"
	"strings"
)

func showDefault(lines []string, opt options) (res []string) {
	if len(lines) > 0 {
		res = append(res, lines[0])
	}

	compareRule := func(left string, right string) bool {
		return left == right
	}

	if opt.registerNotImportant {
		compareRule = func(left string, right string) bool {
			return strings.EqualFold(left, right)
		}
	}

	for i := 1; i < len(lines)-1; i++ {
		prev := lines[i-1]
		cur := lines[i]

		if !compareRule(cur, prev) {
			res = append(res, cur)
		}
	}

	return res
}

type Node struct {
	key   string
	value int
}

func showAll(lines []string, opt options) (res []string) {
	sequence := []Node{}

	var countNodes int

	if len(lines) > 0 {
		sequence = append(sequence, Node{lines[0], 1})
		countNodes++
	}

	compareRule := func(left string, right string) bool {
		return left == right
	}

	if opt.registerNotImportant {
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
		res = append(res, strconv.Itoa(node.value)+" "+node.key)
	}

	return res
}

func showUniq(lines []string, opt options) (res []string) {
	sequence := []Node{}

	var countNodes int

	if len(lines) > 0 {
		sequence = append(sequence, Node{lines[0], 1})
		countNodes++
	}

	compareRule := func(left string, right string) bool {
		return left == right
	}

	if opt.registerNotImportant {
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
			res = append(res, node.key)
		}
	}

	return res
}

func showUnUniq(lines []string, opt options) (res []string) {
	sequence := []Node{}

	var countNodes int

	if len(lines) > 0 {
		sequence = append(sequence, Node{lines[0], 1})
		countNodes++
	}

	compareRule := func(left string, right string) bool {
		return left == right
	}

	if opt.registerNotImportant {
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
			res = append(res, node.key)
		}
	}

	return res
}
