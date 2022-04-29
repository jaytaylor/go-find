package find

import (
	"github.com/ryanuber/go-glob"
)

type wholeNamePredicate struct {
	pattern string
}

func (p *wholeNamePredicate) Match(_ string, path string) (bool, error) {
	return glob.Glob(p.pattern, path), nil
}
