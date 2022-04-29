package find

import (
	"path/filepath"
	"testing"
)

func TestDepth(t *testing.T) {
	sep := string(filepath.Separator)
	testCases := []struct {
		root     string
		path     string
		expected int
	}{
		{
			root:     "",
			path:     "",
			expected: 0,
		},
		{
			root:     "",
			path:     filepath.Join("foo", "bar", "baz", "a", "b", "c"),
			expected: 6,
		},
		{
			root:     "",
			path:     filepath.Join("foo", "bar", "baz", "a", "b", "c") + sep,
			expected: 6,
		},
		{
			root:     filepath.Join("foo", "bar", "baz"),
			path:     filepath.Join("foo", "bar", "baz", "a", "b", "c"),
			expected: 3,
		},
		{
			root:     filepath.Join("foo", "bar", "baz") + sep,
			path:     filepath.Join("foo", "bar", "baz", "a", "b", "c"),
			expected: 3,
		},
	}
	for i, testCase := range testCases {
		if expected, actual := testCase.expected, depth(testCase.root, testCase.path); actual != expected {
			t.Errorf("[testCase=%va] Expected depth=%v but result=%v for % #v", i, expected, actual, testCase)
		}
		testCase.root = sep + testCase.root
		testCase.path = sep + testCase.path
		if expected, actual := testCase.expected, depth(testCase.root, testCase.path); actual != expected {
			t.Errorf("[testCase=%vb] Expected depth=%v but result=%v for % #v", i, expected, actual, testCase)
		}
	}
}
