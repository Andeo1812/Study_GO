package handlers

import (
	"errors"
	"uniq/project/uniq/options"
)

func getDuplicateLines(lines []string, opt options.Options) ([]string, error) {
	res := make([]string, 0)

	sequence := make([]node, 0)

	var countNodes int

	if len(lines) > 0 {
		sequence = append(sequence, node{lines[0], 1})
		countNodes++
	} else {
		return res, errors.New("no data")
	}

	compareRule := getTypeComparator(opt.RegisterNotImportant)

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

	return res, nil
}
