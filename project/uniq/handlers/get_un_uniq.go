package handlers

import (
	"strings"
	"uniq/project/uniq/options"
)

func getUnUniqLines(lines []string, opt options.Options) (res []string) {
	sequence := make([]node, 0)

	var countNodes int

	if len(lines) > 0 {
		sequence = append(sequence, node{lines[0], 1})
		countNodes++
	}

	compareRule := strings.Compare

	if opt.RegisterNotImportant {
		compareRule = func(left string, right string) int {
			if strings.EqualFold(left, right) {
				return 0
			}

			return -1
		}
	}

	for i := 1; i < len(lines); i++ {
		prev := lines[i-1]
		cur := lines[i]

		if compareRule(cur, prev) != 0 {
			sequence = append(sequence, node{cur, 1})
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
