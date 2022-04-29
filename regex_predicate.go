package find

import (
	"regexp"
)

type regexPredicate struct {
	expr *regexp.Regexp
}

func (p *regexPredicate) Match(_ string, path string) (bool, error) {
	return p.expr.MatchString(path), nil
}
