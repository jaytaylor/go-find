package find

import (
	"path/filepath"

	"github.com/ryanuber/go-glob"
)

type namePredicatae struct {
	pattern string
}

func (p *namePredicatae) Match(_ string, path string) (bool, error) {
	return glob.Glob(p.pattern, filepath.Base(path)), nil
}
