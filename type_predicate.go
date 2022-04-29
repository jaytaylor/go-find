package find

import (
	"fmt"
	"os"
)

type typePredicate struct {
	t string
}

func (p *typePredicate) Match(_ string, path string) (bool, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return false, fmt.Errorf("typePredicate: lstat %q: %s", path, err)
	}

	var (
		isCharDev   = info.Mode()&os.ModeCharDevice != 0
		isSymlink   = info.Mode()&os.ModeSymlink != 0 // True if the file is a symlink.
		isNamedPipe = info.Mode()&os.ModeNamedPipe != 0
		isSocket    = info.Mode()&os.ModeSocket != 0
	)

	switch p.t {
	case "c": // Unix character device.
		return isCharDev, nil
	case "d": // Directory
		return info.IsDir() && !isSymlink && !isNamedPipe && !isSocket, nil
	case "f": // Regular file
		return !info.IsDir() && !isSymlink && !isNamedPipe && !isSocket, nil
	case "l": // Symbolic link
		return isSymlink, nil
	case "p": // Named pipe
		return isNamedPipe, nil
	case "s": // Socket
		return isSocket, nil
	}
	return false, p.validate()
}

func (p *typePredicate) validate() error {
	if p.t == "d" || p.t == "f" || p.t == "l" || p.t == "p" || p.t == "s" {
		return nil
	}
	return fmt.Errorf("unrecognized value for -type: %s", p.t)
}
