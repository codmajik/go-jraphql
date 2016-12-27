package gojraphql

import (
	"encoding/json"
	"errors"
)

const (
	JQLActionMutation = "@mutate"
	JQLActionQuery    = "@query"
)

type bfld struct {
	FieldName string      `json:"@name,omitempty"`
	AliasName string      `json:"@alias,omitempty"`
	Fields    []*JQLField `json:"@fields,omitempty"`
}

type JQLField struct {
	FieldName string
	AliasName string
	Fields    []*JQLField
}

func (field *JQLField) UnmarshalJSON(b []byte) error {
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

func (j *JQLField) String() string {
	return j.FieldName
}

type JQL struct {
	ActionType string
	Action     string
	Args       map[string]interface{}
	Fields     []*JQLField
	Extra      map[string]interface{}
}

func (ql *JQL) UnmarshalJSON(b []byte) error {
	jJ := make(map[string]json.RawMessage)

	if err := json.Unmarshal(b, &jJ); err != nil {
		return err
	}

	for key, item := range jJ {
		switch key {
		case JQLActionQuery:
			fallthrough
		case JQLActionMutation:
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
			ql.Fields = make([]*JQLField, 0)
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

	return nil
}
