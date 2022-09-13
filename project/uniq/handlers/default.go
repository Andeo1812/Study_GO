package handlers

import (
	"strings"
	"uniq/project/uniq/options"
)

func defaultHandlerLines(lines []string, opt options.Options) (res []string) {
	if len(lines) > 0 {
		res = append(res, lines[0])
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

	for i := 1; i < len(lines)-1; i++ {
		prev := lines[i-1]
		cur := lines[i]

		if compareRule(cur, prev) != 0 {
			res = append(res, cur)
		}
	}

	return
}
