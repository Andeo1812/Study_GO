package handlers

import (
	"Modules/project/uniq/options"
)

type Node struct {
	key   string
	value int
}

func ClassifierHandlers(lines []string, opt options.Options) (res []string) {
	switch {
	case opt.ShowCountStr:
		return getCount(lines, opt)
	case opt.ShowUniqStr:
		return getUniq(lines, opt)
	case opt.ShowUnUniqStr:
		return getUnUniq(lines, opt)
	default:
		return defaultHandler(lines, opt)
	}
}