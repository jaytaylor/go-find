package find

import (
	"regexp"
)

type Find struct {
	Paths      []string
	predicates predicates
}

func NewFind(paths ...string) *Find {
	find := &Find{
		Paths: paths,
	}
	return find
}

func (finder *Find) Evaluate() ([]string, error) {
	results := []string{}
	for _, path := range finder.Paths {
		hits, err := finder.predicates.Evaluate(path)
		if err != nil {
			return nil, err
		}
		results = append(results, hits...)
	}
	return results, nil
}

func (finder *Find) MinDepth(n int) *Find {
	c := &minDepthPredicate{
		n: n,
	}
	finder.predicates = append(finder.predicates, c)
	return finder
}

func (finder *Find) MaxDepth(n int) *Find {
	c := &maxDepthPredicate{
		n: n,
	}
	finder.predicates = append(finder.predicates, c)
	return finder
}

func (finder *Find) Type(t string) *Find {
	c := &typePredicate{
		t: t,
	}
	finder.predicates = append(finder.predicates, c)
	return finder
}

func (finder *Find) Name(pattern string) *Find {
	c := &namePredicatae{
		pattern: pattern,
	}
	finder.predicates = append(finder.predicates, c)
	return finder
}

func (finder *Find) WholeName(pattern string) *Find {
	c := &wholeNamePredicate{
		pattern: pattern,
	}
	finder.predicates = append(finder.predicates, c)
	return finder
}

func (finder *Find) Regex(expr *regexp.Regexp) *Find {
	c := &regexPredicate{
		expr: expr,
	}
	finder.predicates = append(finder.predicates, c)
	return finder
}

func (finder *Find) Empty() *Find {
	c := &emptyPredicate{}
	finder.predicates = append(finder.predicates, c)
	return finder
}

func (finder *Find) Mount() *Find {
	fsType := getFileSystemType(finder.Paths...)
	c := &mountPredicate{
		path:   finder.Paths,
		fsType: fsType,
	}
	finder.predicates = append(finder.predicates, c)
	return finder
}
