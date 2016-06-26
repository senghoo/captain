package command

import (
	"errors"
	"fmt"
	"reflect"
)

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

func UpdateArgs(obj interface{}, args CommandArgs) error {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	fieldNum := t.NumField()
	if fieldNum == 0 {
		return errors.New("Not Contains any field")
	}

	for i := 0; i < fieldNum; i++ {
		field := t.Field(i)
		switch field.Tag.Get("command") {
		case "-":
			continue
		default:
			value := v.Field(i)
			updateValue(value, field.Name, args)
		}
	}

	return nil
}

func updateValue(value reflect.Value, name string, args CommandArgs) error {
	t := value.Type().Name()
	switch t {
	case "int", "int64":
		v, ok := args.Int64(name)
		if !ok {
			return fmt.Errorf("%s not exist", name)
		}
		value.SetInt(v)
	case "string":
		v, ok := args.String(name)
		if !ok {
			return fmt.Errorf("%s not exist", name)
		}
		value.SetString(v)
	default:
		return fmt.Errorf("unsupported type %s", t)

	}

}
