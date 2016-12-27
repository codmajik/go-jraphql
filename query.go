package gojraphql

import (
	"encoding/json"
	"errors"
)

const (
	QueryActionMutation = "@mutate"
	QueryActionQuery    = "@query"
)

type bfld struct {
	FieldName string        `json:"@name,omitempty"`
	AliasName string        `json:"@alias,omitempty"`
	Fields    []*QueryField `json:"@fields,omitempty"`
}

type QueryField struct {
	FieldName string
	AliasName string
	Fields    []*QueryField
}

func (field *QueryField) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &field.FieldName)
	if err == nil {
		return nil
	}

	bfield := &bfld{}
	err = json.Unmarshal(b, &bfield)

	field.FieldName = bfield.FieldName
	field.AliasName = bfield.AliasName
	field.Fields = bfield.Fields
	return err
}

func (j *QueryField) String() string {
	return j.FieldName
}

type Query struct {
	ActionType string
	Action     string
	Args       map[string]interface{}
	Fields     []*QueryField
	Extra      map[string]interface{}
}

func (ql *Query) UnmarshalJSON(b []byte) error {
	jJ := make(map[string]json.RawMessage)

	if err := json.Unmarshal(b, &jJ); err != nil {
		return err
	}

	for key, item := range jJ {
		switch key {
		case QueryActionQuery:
			fallthrough
		case QueryActionMutation:
			if ql.ActionType != "" {
				return errors.New("CONFUSED: which action @query or mutation. make up you mind")
			}

			ql.ActionType = key
			if err := json.Unmarshal(item, &ql.Action); err != nil {
				return errors.New(key + " must be a string")
			}

		case "@args":
			ql.Args = make(map[string]interface{})
			if err := json.Unmarshal(item, &ql.Args); err != nil {
				return errors.New(key + " must be an array")
			}

		case "@fields":
			ql.Fields = make([]*QueryField, 0)
			if err := json.Unmarshal(item, &ql.Fields); err != nil {
				return errors.New("@fields must be a map of string and/or field descriptions")
			}

		default:
			if ql.Extra == nil {
				ql.Extra = make(map[string]interface{})
			}
			var i interface{}
			if err := json.Unmarshal(item, &i); err != nil {
				return err
			}

			ql.Extra[key] = i
		}
	}

	if ql.ActionType == "" {
		return errors.New("CONFUSED: which action @query or mutation. none specified")
	}
	return nil
}
