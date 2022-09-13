package handlers

import (
	"Modules/project/uniq/options"
	"strings"
)

func getUnUniq(lines []string, opt options.Options) (res []string) {
	sequence := make([]Node, 0)

	var countNodes int

	if len(lines) > 0 {
		sequence = append(sequence, Node{lines[0], 1})
		countNodes++
	}

	compareRule := func(left string, right string) bool {
		return left == right
	}

	if opt.RegisterNotImportant {
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
