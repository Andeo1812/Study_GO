package handlers

import (
	"errors"
	"uniq/project/uniq/options"
)

func getUniqLines(lines []string, opt options.Options) ([]string, error) {
	if len(lines) == 0 {
		return nil, errors.New("no data")
	}

	sequence := []node{node{lines[0], 1}}
	var countNodes int = 1

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

	res := make([]string, 0)

	for _, node := range sequence {
		if node.value == 1 {
			res = append(res, node.key)
		}
	}

	return res, nil
}
