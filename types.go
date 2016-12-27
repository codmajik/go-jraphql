package gojraphql

import (
	"errors"
)

//ErrTypeRequierd error when a required field/argument is ommited
var ErrTypeRequierd = errors.New("REQUIRED")

type JraphQLType interface {
	Validate(v interface{}) error
}

// Generics would have worked wonderfully here; but ....
type JraphQLInt struct {
	Required bool
	Nullable bool
}

func (j *JraphQLInt) Validate(v interface{}) error {
	_, ok := v.(int)
	if ok {
		return nil
	}

	if j.Required {
		return errors.New("Required")
	}

	if v == nil && j.Nullable {
		return nil
	}

	return errors.New("cannot be null")
}

// Generics would have worked wonderfully here; but ....
type JraphQLStr struct {
	Required bool
	Nullable bool
}

func (j *JraphQLStr) Validate(v interface{}) error {
	_, ok := v.(int)
	if ok {
		return nil
	}

	if j.Required {
		return errors.New("Required")
	}

	if v == nil && j.Nullable {
		return nil
	}

	return errors.New("cannot be null")
}
