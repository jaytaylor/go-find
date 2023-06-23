package find

import (
	"fmt"
	"os"
	"path/filepath"
)

type predicate interface {
	Match(root string, path string) (bool, error)
}

type PredicateError struct {
	errType    error
	errMessage string
}

func (p PredicateError) Error() string {
	return fmt.Sprintf("%q: %s", p.errType, p.errMessage)
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
				if pe, ok := err.(PredicateError); ok {
					// ignore the directory if encountered a different filesystem
					if pe.errType == ErrorFSType {
						return filepath.SkipDir
					} else {
						fmt.Fprintf(os.Stderr, "error: %s\n", pe.errMessage)
						return nil
					}
				}
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
