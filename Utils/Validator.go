package Utils

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

const Type = "type:"
const Map = "map"
const String = "string"
const Number = "number"
const Boolean = "boolean"

type Element struct {
	Name string       `json:"name"`
	Kind reflect.Kind `json:"Kind"`
}

func Validate(InputOBJ interface{}, Structure interface{}) bool {

	InV := reflect.ValueOf(InputOBJ)
	StV := reflect.ValueOf(Structure)

	var InVElements = getElementsForInputOBJ(InV)
	var StVElements = getElementsForStructure(StV)

	return compareElements(InVElements, StVElements)
}

func getElementsForStructure(OBJ reflect.Value) map[string]Element {
	var elements = make(map[string]Element)
	if OBJ.Kind() == reflect.Map {
		for _, key := range OBJ.MapKeys() {
			e, _ := getElement(key.String(), OBJ.MapIndex(key))
			elements[e.Name] = e
		}
	}
	return elements
}

func getElementsForInputOBJ(OBJ reflect.Value) map[string]Element {
	var elements = make(map[string]Element)
	if OBJ.Kind() == reflect.Map {
		for _, key := range OBJ.MapKeys() {
			v := fmt.Sprintf("%v", OBJ.MapIndex(key))
			elements[key.String()] = Element{
				Name: key.String(), Kind: getOBJKind(v),
			}
		}
	}
	return elements
}

func compareElements(ASeries map[string]Element, BSeries map[string]Element) bool {

	if len(ASeries) != len(BSeries) {
		return false
	}

	for key, value := range ASeries {
		element, ok := BSeries[key]
		if !ok {
			return false
		} else {
			if element.Kind != value.Kind {
				return false
			}
		}
	}
	return true
}

func getElement(name string, v reflect.Value) (Element, error) {
	str := fmt.Sprintf("%v", v.Interface())
	k := parseType(str)
	return Element{
		Name: name, Kind: k,
	}, nil
}

func parseType(str string) reflect.Kind {
	return getStructureKind(str)
}

func getStructureKind(str string) reflect.Kind {
	var index = strings.Index(str, Type)
	var _type = str[index+len(Type) : len(str)-1]

	switch _type {
	case String:
		return reflect.String
	case Boolean:
		return reflect.Bool
	case Number:
		return reflect.Int
	default:
		return reflect.Invalid
	}
}

func getOBJKind(str string) reflect.Kind {
	if isInt(str) {
		return reflect.Int
	} else if isBool(str) {
		return reflect.Bool
	} else {
		return reflect.String
	}
}

func isInt(s string) bool {
	if len(s) < 1 {
		return false
	}
	s = strings.TrimPrefix(s, "-")
	return unicode.IsDigit(rune(s[0]))
}

func isBool(str string) bool {
	return str == "true" || str == "false"
}
