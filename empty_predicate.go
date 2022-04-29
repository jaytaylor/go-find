package find

import (
	"fmt"
	"os"
)

type emptyPredicate struct{}

func (p *emptyPredicate) Match(root string, path string) (bool, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return false, fmt.Errorf("emptyPredicate: lstat %q: %s", path, err)
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
			return false, fmt.Errorf("emptyPredicate: opening %q: %s", info.Name(), err)
		}
		dirs, err := f.ReadDir(-1)
		f.Close()
		if err != nil {
			return false, fmt.Errorf("emptyPredicate: listing %q: %s", info.Name(), err)
		}
		return len(dirs) == 0, nil
	} else {
		return info.Size() == 0, nil
	}
}
