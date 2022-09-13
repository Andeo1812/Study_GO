package handlers

import (
	"errors"
	"uniq/project/uniq/options"
)

func defaultHandlerLines(lines []string, opt options.Options) ([]string, error) {
	res := make([]string, 0)

	if len(lines) == 0 {
		return res, errors.New("no data")
	}

	res = append(res, lines[0])

	compareRule := getTypeComparator(opt.RegisterNotImportant)

	for i := 1; i < len(lines)-1; i++ {
		prev := lines[i-1]
		cur := lines[i]

		if compareRule(cur, prev) != 0 {
			res = append(res, cur)
		}
	}

	return res, nil
}
