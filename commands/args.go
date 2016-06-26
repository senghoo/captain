package command

type CommandArgs map[string]interface{}

func (a CommandArgs) String(name string) (string, bool) {
	i, ok := a[name]
	if !ok {
		return "", false
	}

	v, ok := i.(string)
	return v, ok
}

func (a CommandArgs) Int(name string) (int, bool) {
	i, ok := a[name]
	if !ok {
		return 0, false
	}

	v, ok := i.(int)
	return v, ok
}

func (a CommandArgs) Int64(name string) (int64, bool) {
	i, ok := a[name]
	if !ok {
		return 0, false
	}

	v, ok := i.(int64)
	return v, ok
}
