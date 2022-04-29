package find

import (
	"path/filepath"
	"strings"
)

// depth calculates the relative depth of a path.
func depth(root string, path string) int {
	root = filepath.Clean(root)
	path = filepath.Clean(path)
	sep := string(filepath.Separator)
	if root != sep && !strings.HasSuffix(root, sep) {
		root += sep
	}
	if path != sep && strings.HasSuffix(path, sep) {
		path = strings.TrimRight(path, sep)
	}
	diff := strings.TrimPrefix(path, root)
	pieces := strings.Split(diff, sep)
	d := len(pieces)
	if d == 1 {
		return 0
	}
	return d
}
