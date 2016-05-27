package git

func Clone(source, path string) {
	c := NewCommand("clone", source, path)
	c.Run()
}
