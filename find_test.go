package find

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestFind(t *testing.T) {
	basePath := "/tmp/test-go-find"
	if err := os.RemoveAll(basePath); err != nil {
		t.Fatal(err)
	}

	if err := os.MkdirAll(filepath.Join(basePath, "foo", "bar", "baz"), os.FileMode(int(0777))); err != nil {
		t.Fatal(err)
	}
	defer func() {
		os.RemoveAll(basePath)
	}()

	if err := os.MkdirAll(filepath.Join(basePath, "some", "other", "dir", "foo"), os.FileMode(int(0777))); err != nil {
		t.Fatal(err)
	}

	if err := ioutil.WriteFile(filepath.Join(basePath, "file0.txt"), nil, os.FileMode(int(0600))); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(filepath.Join(basePath, "file1.txt"), []byte("hello here 1\n"), os.FileMode(int(0600))); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(filepath.Join(basePath, "file2.txt"), []byte("hello here 2\n"), os.FileMode(int(0600))); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(filepath.Join(basePath, "foo", "file1.txt"), []byte("hello there 1\n"), os.FileMode(int(0600))); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(filepath.Join(basePath, "foo", "file2.txt"), []byte("hello there 2\n"), os.FileMode(int(0600))); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(filepath.Join(basePath, "foo", "bar", "baz", "file1.txt"), []byte("hello over there 1\n"), os.FileMode(int(0600))); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(filepath.Join(basePath, "some", "other", "file0.txt"), nil, os.FileMode(int(0600))); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(filepath.Join(basePath, "some", "other", "dir", "null.txt"), nil, os.FileMode(int(0600))); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(filepath.Join(basePath, "some", "other", "dir", "null.txt"), filepath.Join(basePath, "lnk")); err != nil {
		t.Fatal(err)
	}

	// No predicates
	{
		hits, err := NewFind(basePath).Evaluate()
		if err != nil {
			t.Fatal(err)
		}
		if expected, actual := 17, len(hits); actual != expected {
			t.Errorf("Expected num–hits=%v but actual=%v\nHits: % #v)\n", expected, actual, hits)
		}
	}

	// MinDepth
	{
		hits, err := NewFind(basePath).MinDepth(1).Name("file1.txt").Evaluate()
		if err != nil {
			t.Fatal(err)
		}
		if expected, actual := 2, len(hits); actual != expected {
			t.Errorf("Expected num–hits=%v but actual=%v\nHits: % #v)\n", expected, actual, hits)
		}
	}

	// MaxDepth
	{
		hits, err := NewFind(basePath).MaxDepth(1).Evaluate()
		if err != nil {
			t.Fatal(err)
		}
		if expected, actual := 6, len(hits); actual != expected {
			t.Errorf("Expected num–hits=%v but actual=%v\nHits: % #v)\n", expected, actual, hits)
		}
	}

	// Name (exact)
	{
		hits, err := NewFind(basePath).Name("file1.txt").Evaluate()
		if err != nil {
			t.Fatal(err)
		}
		if expected, actual := 3, len(hits); actual != expected {
			t.Errorf("Expected num–hits=%v but actual=%v\nHits: % #v)\n", expected, actual, hits)
		}
	}

	// Name (glob)
	{
		hits, err := NewFind(basePath).Name("*1.txt").Evaluate()
		if err != nil {
			t.Fatal(err)
		}
		if expected, actual := 3, len(hits); actual != expected {
			t.Errorf("Expected num–hits=%v but actual=%v\nHits: % #v)\n", expected, actual, hits)
		}
	}

	// MaxDepth + Name
	{
		hits, err := NewFind(basePath).MaxDepth(1).Name("file1.txt").Evaluate()
		if err != nil {
			t.Fatal(err)
		}
		if expected, actual := 1, len(hits); actual != expected {
			t.Errorf("Expected num–hits=%v but actual=%v\nHits: % #v)\n", expected, actual, hits)
		}
	}

	// Wholename
	{
		hits, err := NewFind(basePath).WholeName("*foo*2.txt").Evaluate()
		if err != nil {
			t.Fatal(err)
		}
		if expected, actual := 1, len(hits); actual != expected {
			t.Errorf("Expected num–hits=%v but actual=%v\nHits: % #v)\n", expected, actual, hits)
		}
	}

	// Type dir
	{
		hits, err := NewFind(basePath).Type("d").Evaluate()
		if err != nil {
			t.Fatal(err)
		}
		if expected, actual := 8, len(hits); actual != expected {
			t.Errorf("Expected num–hits=%v but actual=%v\nHits: % #v)\n", expected, actual, hits)
		}
	}

	// Type file
	{
		hits, err := NewFind(basePath).Type("d").Evaluate()
		if err != nil {
			t.Fatal(err)
		}
		if expected, actual := 8, len(hits); actual != expected {
			t.Errorf("Expected num–hits=%v but actual=%v\nHits: % #v)\n", expected, actual, hits)
		}
	}

	// Empty
	{
		hits, err := NewFind(basePath).Empty().Evaluate()
		if err != nil {
			t.Fatal(err)
		}
		if expected, actual := 4, len(hits); actual != expected {
			t.Errorf("Expected num–hits=%v but actual=%v\nHits: % #v)\n", expected, actual, hits)
		}
	}

	// Empty dirs
	{
		hits, err := NewFind(basePath).Type("d").Empty().Evaluate()
		if err != nil {
			t.Fatal(err)
		}
		if expected, actual := 1, len(hits); actual != expected {
			t.Errorf("Expected num–hits=%v but actual=%v\nHits: % #v)\n", expected, actual, hits)
		}
	}

	// Empty files
	{
		hits, err := NewFind(basePath).Type("f").Empty().Evaluate()
		if err != nil {
			t.Fatal(err)
		}
		if expected, actual := 3, len(hits); actual != expected {
			t.Errorf("Expected num–hits=%v but actual=%v\nHits: % #v)\n", expected, actual, hits)
		}
	}
}
