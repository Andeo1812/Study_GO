package handlers

import (
	"strconv"
	"uniq/project/uniq/options"
)

func getCountLines(lines []string, opt options.Options) (res []string) {
	sequence := make([]node, 0)

	var countNodes int

	if len(lines) > 0 {
		sequence = append(sequence, node{lines[0], 1})
		countNodes++
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
		res = append(res, strconv.Itoa(node.value)+" "+node.key)
	}

	return
}
