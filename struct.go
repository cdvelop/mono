package mono

import (
	"reflect"
	"strings"
)

type structHandler struct {
	objectName string //ej: attention.user,
	structName string //ej: User
	anonymous  bool

	reflectKind  reflect.Kind
	reflectType  reflect.Type
	reflectValue reflect.Value
}

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
		s.structName = "unknown"
	}
}

func (s *structHandler) validate(structIN any, requiredPointer bool) error {
	if structIN == nil {
		return R.Err(D.Pointer, D.TheStructure, D.Is, D.Nil)
	}

	s.reflectValue = reflect.ValueOf(structIN)
	s.reflectKind = s.reflectValue.Kind()

	// Verifica si structIN es un puntero de una estructura
	if requiredPointer && s.reflectKind != reflect.Ptr {
		return R.Err(D.TheStructure, s.Name(structIN), D.IsNotOfPointerType)
	}

	// Si es un puntero, obtén el valor apuntado
	isPointer := s.reflectKind == reflect.Ptr
	if isPointer {
		s.reflectValue = s.reflectValue.Elem()
		s.reflectKind = s.reflectValue.Kind()

		if !requiredPointer {
			s.buildName()
			return R.Err(D.TheStructure, s.Name(), D.IsNotRequired, D.AsAPointer)
		}
	}

	s.buildName()

	// Verifica si structIN es una estructura
	if s.reflectKind != reflect.Struct {
		typeName := s.reflectKind.String()
		if s.reflectKind == reflect.Ptr {
			typeName = s.reflectValue.Type().String()
		}
		s.anonymous = false
		return R.Err(D.TheElement, typeName, D.IsNotOfStructureType)
	}

	// Configura anonymous basado en si es una estructura anónima
	s.anonymous = s.structName == "unknown"

	return nil
}
