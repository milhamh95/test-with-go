package form

import (
	"reflect"
)

func fields(strct interface{}) field {
	rv := reflect.ValueOf(strct)

	t := rv.Type()

	tf := t.Field(0)

	return field{
		Label:       tf.Name,
		Name:        tf.Name,
		Type:        "text",
		Placeholder: tf.Name,
	}
}

type field struct {
	Label       string
	Name        string
	Type        string
	Placeholder string
}
