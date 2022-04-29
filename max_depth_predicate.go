package find

type maxDepthPredicate struct {
	n int
}

func (p *maxDepthPredicate) Match(root string, path string) (bool, error) {
	d := depth(root, path)
	return d <= p.n, nil
}
