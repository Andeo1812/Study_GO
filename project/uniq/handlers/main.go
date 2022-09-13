package handlers

import (
	"uniq/project/uniq/options"
)

func ClassifierHandlers(lines []string, opt options.Options) ([]string, error) {
	switch {
	case opt.ShowCountStr:
		return getCountLines(lines, opt)
	case opt.ShowUniqStr:
		return getUniqLines(lines, opt)
	case opt.ShowDuplicateStr:
		return getDuplicateLines(lines, opt)
	default:
		return defaultHandlerLines(lines, opt)
	}
}
