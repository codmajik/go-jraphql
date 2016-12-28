package gojraphql

import (
	"errors"
	"io"
	"sync"
)

type queryResolver struct {
	cb      ResolverFunc
	retInfo *returnType
	argInfo []*typeInfo
}

type Schema struct {
	sync.RWMutex
	knownTypes    []interface{}
	queryActions  map[string]*queryResolver
	mutateActions map[string]*queryResolver
}

func NewSchema(r io.Reader) (*Schema, error) {
	s := &Schema{
		queryActions:  make(map[string]*queryResolver),
		mutateActions: make(map[string]*queryResolver),
	}

	if err := parse(s, r); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Schema) SetQueryResolver(label string, h ResolverFunc) error {
	r, ok := s.queryActions[label]
	if !ok {
		return errors.New("Cannot find query: " + label)
	}

	r.cb = h
	return nil
}

func (s *Schema) SetMutationResolver(label string, h ResolverFunc) error {
	r, ok := s.mutateActions[label]
	if !ok {
		return errors.New("Cannot find mutation: " + label)
	}

	r.cb = h
	return nil
}
