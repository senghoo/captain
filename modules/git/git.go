package git

func Clone(source, path string) (string, error) {
	c := NewCommand("clone", source, path)
	return c.Run()
}

func Archive(path, format, branch, file string) (string, error) {
	c := NewCommand("archive", "--format", format, branch, "-o", file)
	return c.RunInDir(path)
}
