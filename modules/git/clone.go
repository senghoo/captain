package git

func Clone(source, path string) (string, error) {
	c := NewCommand("clone", source, path)
	return c.Run()
}
