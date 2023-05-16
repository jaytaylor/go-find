package find

import (
	"fmt"
	"os"
)

var (
	ErrorEmptyPredStat    = fmt.Errorf("emptyPredicate: lstat")
	ErrorEmptyPredOpen    = fmt.Errorf("emptyPredicate: opening")
	ErrorEmptyPredListing = fmt.Errorf("emptyPredicate: listing")
)

type emptyPredicate struct{}

func (p *emptyPredicate) Match(root string, path string) (bool, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return false, PredicateError{errType: ErrorEmptyPredStat, errMessage: err.Error()}
	}

	var (
		isCharDev   = info.Mode()&os.ModeCharDevice != 0
		isSymlink   = info.Mode()&os.ModeSymlink != 0 // True if the file is a symlink.
		isNamedPipe = info.Mode()&os.ModeNamedPipe != 0
		isSocket    = info.Mode()&os.ModeSocket != 0
	)

	if isCharDev || isSymlink || isNamedPipe || isSocket {
		return false, nil
	}

	if info.IsDir() {
		f, err := os.Open(path)
		if err != nil {
			return false, PredicateError{errType: ErrorEmptyPredOpen, errMessage: err.Error()}
		}
		dirs, err := f.ReadDir(-1)
		f.Close()
		if err != nil {
			fmt.Println(err)
			return false, PredicateError{errType: ErrorEmptyPredListing, errMessage: err.Error()}
		}
		return len(dirs) == 0, nil
	} else {
		return info.Size() == 0, nil
	}
}
