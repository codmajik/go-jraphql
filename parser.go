package gojraphql

import (
	"encoding/json"
	// "errors"
	"errors"
	"io"
	"strings"
)

const (
	ParserScopeSchema   = "@schema"
	ParserScopeQuery    = "@query"
	ParserScopeMutation = "@mutation"
)

var scalarTypes = map[string]bool{
	"int":   true,
	"str":   true,
	"float": true,
	"bool":  true,
}

func parse(s *Schema, desc io.Reader) error {
	s.Lock()
	defer s.Unlock()

	_schema := make(map[string]json.RawMessage)

	err := json.NewDecoder(desc).Decode(&_schema)
	if err != nil {
		println(err.Error())
		return err
	}

	for item, def := range _schema {
		println(item)
		switch item {
		case ParserScopeSchema:
			println("populate our type table", string(def))

		case ParserScopeQuery:
			queries := make(map[string]map[string]json.RawMessage)

			err = json.Unmarshal(def, &queries)
			if err != nil {
				return err
			}

			for qkey, q := range queries {
				r := &queryResolver{
					argInfo: make([]*typeInfo, 0),
				}

				rtn, ok := q["@return"]
				if !ok {
					println("return not set")
					continue
				}

				rtype, err := parseReturnType(rtn)
				if err != nil {
					println("return set", err.Error())
				}

				r.retInfo = rtype

				s.queryActions[qkey] = r
			}

			// println("handle query", string(def))
		case ParserScopeMutation:
			println("handle muation", string(def))

		default:
			println("unknown section skipping", item, def)
		}
	}

	return nil
}

type typeInfo struct {
	typeName   string
	refType    bool
	valueType  string
	isNullable bool
	isScalar   bool
	isArray    bool
	isEnum     bool

	subTypes map[string]*typeInfo
}

type returnType struct {
	*typeInfo
}

func parseReturnType(j json.RawMessage) (*returnType, error) {
	tInfo, err := parseTypeInfo(j)
	if err != nil {
		return nil, err
	}

	return &returnType{
		typeInfo: tInfo,
	}, nil
}

func parseTypeInfo(j json.RawMessage) (*typeInfo, error) {
	println(string(j))

	tInfo := &typeInfo{}
	if parseScalar(j, tInfo) == nil {
		return tInfo, nil
	}

	var arrayOfType []json.RawMessage
	if json.Unmarshal(j, &arrayOfType) == nil {

		if len(arrayOfType) != 1 {
			return nil, errors.New("expects a single item array for type defination")
		}

		if err := parseScalar(arrayOfType[0], tInfo); err != nil {
			return nil, errors.New("array type must be a scalar or pre-defined custom type")
		}

		tInfo.isArray = true
		return tInfo, nil
	}

	var objOfType map[string]json.RawMessage
	if json.Unmarshal(j, &objOfType) == nil {

		tInfo.isEnum = false
		tInfo.isScalar = false
		tInfo.isArray = false
		if _type, ok := objOfType["$type"]; ok {
			println("objectType", string(_type))
		}

		// if err := parseScalar(j, tInfo); err != nil {
		// 	return nil, errors.New("array type must be a scalar or pre-defined custom type")
		// }

		return tInfo, nil
	}

	return tInfo, nil
}

func parseScalar(r json.RawMessage, ti *typeInfo) error {
	lookupType := ""
	if err := json.Unmarshal(r, &lookupType); err != nil {
		return err
	}

	ti.isNullable = !strings.HasSuffix(lookupType, "!")
	ti.refType = strings.HasPrefix(lookupType, "$")

	ti.typeName = strings.TrimSuffix(strings.TrimPrefix(lookupType, "$"), "!")
	ti.isScalar = isScalarType(ti.typeName)

	return nil
}

func parseEnumValues(r json.RawMessage, ti *typeInfo) error {

	return nil
}

func isScalarType(name string) bool {
	println("check scalarTypes")
	return scalarTypes[name] == true
}
