package command

import (
	"errors"
	"fmt"
	"reflect"

	"gopkg.in/flosch/pongo2.v3"
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

func (a CommandArgs) Float64(name string) (float64, bool) {
	i, ok := a[name]
	if !ok {
		return 0, false
	}

	v, ok := i.(float64)
	return v, ok
}

func (a CommandArgs) Int64(name string) (int64, bool) {
	switch i := a[name].(type) {
	case int:
		return int64(i), true
	case int64:
		return i, true
	case float64:
		return int64(i), true
	}

	return 0, false
}

func UpdateArgs(obj interface{}, args CommandArgs, context map[string]interface{}) error {
	t := reflect.TypeOf(obj).Elem()
	v := reflect.ValueOf(obj).Elem()
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
			err := updateValue(value, field.Name, args, context)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func updateValue(value reflect.Value, name string, args CommandArgs, context map[string]interface{}) error {
	t := value.Type().Name()
	fmt.Printf("get %s, from %#v, t %s", name, args, t)
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
		tpl, err := pongo2.FromString(v)
		if err != nil {
			return err
		}
		out, err := tpl.Execute(context)
		if err != nil {
			return err
		}
		value.SetString(out)
	default:
		return fmt.Errorf("unsupported type %s", t)

	}
	return nil
}
