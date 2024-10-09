package godi

import (
	"errors"
	"reflect"
	"strings"
)

var (
	// errStAnonymous = errors.New("anonymous structure not supported")
	errStNil = errors.New("the structure pointer is nil")
)

func errNotStruct(name string) error {
	return errors.New("the element " + name + " is not of structure type")
}

func errNoStructPtr(name string) error {
	return errors.New("the structure " + name + " is not of pointer type")
}

func errNoStructPtrReq(name string) error {
	return errors.New("structure " + name + " is not required as a pointer")
}

type structHandler struct {
	objectName string //ej: attention.user,
	structName string //ej: User
	anonymous  bool

	reflectKind  reflect.Kind
	reflectType  reflect.Type
	reflectValue reflect.Value
}

const unknownStName = "unknown"

func (s *structHandler) Name(stIN ...any) string {

	for _, st := range stIN {
		s.structName = reflect.TypeOf(st).Name()
	}

	s.checkAnonymousName()

	return s.structName
}

// construir nombre estructura
func (s *structHandler) buildName() {
	s.reflectType = s.reflectValue.Type()
	s.structName = s.reflectType.Name()
	s.objectName = strings.ToLower(s.reflectType.String())

	s.checkAnonymousName()

}

func (s *structHandler) checkAnonymousName() {
	if s.structName == "" {
		s.anonymous = true
		s.structName = unknownStName
	}
}

func (s *structHandler) validate(structIN any, requiredPointer bool) error {
	if structIN == nil {
		return errStNil
	}

	s.reflectValue = reflect.ValueOf(structIN)
	s.reflectKind = reflect.ValueOf(structIN).Kind()

	// Verifica si structIN es un puntero de una estructura
	if requiredPointer && s.reflectKind != reflect.Ptr {
		return errNoStructPtr(s.Name(structIN))
	}

	// Si es un puntero, obtén el valor apuntado
	if s.reflectKind == reflect.Ptr {
		s.reflectValue = s.reflectValue.Elem()
		s.reflectKind = s.reflectValue.Kind()

		if !requiredPointer {
			s.buildName()
			return errNoStructPtrReq(s.Name())
		}

	}

	// Verifica si structIN es una estructura
	if s.reflectKind != reflect.Struct {
		return errNotStruct(s.reflectKind.String())
	}

	s.buildName()

	return nil
}
