package handlers

import (
	"errors"
	"strconv"
	"uniq/project/uniq/options"
)

func getCountLines(lines []string, opt options.Options) ([]string, error) {
	res := make([]string, 0)

	if len(lines) == 0 {
		return res, errors.New("no data")
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

	for _, node := range sequence {
		res = append(res, strconv.Itoa(node.value)+" "+node.key)
	}

	return res, nil
}
