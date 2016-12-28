package gojraphql

// int | str | float | bool | enum | object
import (
	"errors"
)

//ErrTypeRequierd error when a required field/argument is ommited
var ErrTypeRequierd = errors.New("REQUIRED")

type SchemaItemType interface {
	Validate(v interface{}) error
}

// Generics would have worked wonderfully here; but ....
type Int struct{}

func (j *Int) Validate(v interface{}) error {
	_, ok := v.(int)
	if ok {
		return nil
	}

	return errors.New("NotAString")
}

// Generics would have worked wonderfully here; but ....
type Str struct{}

func (j *Str) Validate(v interface{}) error {
	_, ok := v.(string)
	if ok {
		return nil
	}
	return errors.New("cannot be null")
}

type Enum struct {
	Values map[string]interface{}
}

func (e *Enum) Validate(v interface{}) error {
	key, ok := v.(string)
	if !ok {
		return errors.New("NOT valid enum value must be string")
	}

	_, ok = e.Values[key]
	if ok {
		return nil
	}

	return errors.New("Invalid enum value")
}
