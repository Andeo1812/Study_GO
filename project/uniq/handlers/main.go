package handlers

import (
	"uniq/project/uniq/options"
)

type node struct {
	key   string
	value int
}

func ClassifierHandlers(lines []string, opt options.Options) (res []string) {
	switch {
	case opt.ShowCountStr:
		return getCountLines(lines, opt)
	case opt.ShowUniqStr:
		return getUniqLines(lines, opt)
	case opt.ShowUnUniqStr:
		return getUnUniqLines(lines, opt)
	default:
		return defaultHandlerLines(lines, opt)
	}
}
