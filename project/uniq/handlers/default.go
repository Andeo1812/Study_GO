package handlers

import (
	"Modules/project/uniq/options"
	"strings"
)

func defaultHandler(lines []string, opt options.Options) (res []string) {
	if len(lines) > 0 {
		res = append(res, lines[0])
	}

	compareRule := func(left string, right string) bool {
		return left == right
	}

	if opt.RegisterNotImportant {
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
