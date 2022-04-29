package find

import (
	"os"
	"path/filepath"
)

type predicate interface {
	Match(root string, path string) (bool, error)
}

type predicates []predicate

func (ps predicates) Evaluate(root string) ([]string, error) {
	results := []string{}
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		allMatched := true
		for _, p := range ps {
			matched := false
			if matched, err = p.Match(root, path); err != nil {
				return err
			}
			if !matched {
				allMatched = false
				break
			}
		}
		if allMatched {
			results = append(results, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return results, nil
}
