package find

type minDepthPredicate struct {
	n int
}

func (p *minDepthPredicate) Match(root string, path string) (bool, error) {
	d := depth(root, path)
	return d >= p.n, nil
}
