package gojraphql

import (
	"encoding/json"
)

type ResolveContext struct {
	Action    string
	Fields    []*QueryField
	Arguments map[string]interface{}
}

// ResolverFunc Callback function for Query resolution
type ResolverFunc func(*ResolveContext) (interface{}, error)

type typeDesc struct {
	scaler    bool
	isArray   bool
	valueType string
	subTypes  []string
}

type retDesc struct {
	scaler    bool
	isArray   bool
	valueType string
	arguments []string
}

type schemaDesc struct {
	retType *retDesc
	args    []*typeDesc
}

// Execute using schema validate and reolve specified query
func (s *Schema) Execute(q []byte) (interface{}, error) {
	s.RLock() // the schema can't be modified while we are with it
	defer s.RUnlock()

	r := make(map[string]*Query)
	err := json.Unmarshal([]byte(q), &r)
	if err != nil {
		return nil, err
	}

	return nil, nil

}
