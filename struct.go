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
	s.reflectKind = reflect.ValueOf(structIN).Kind()

	// Verifica si structIN es un puntero de una estructura
	if requiredPointer && s.reflectKind != reflect.Ptr {
		return R.Err(D.TheStructure, s.Name(structIN), D.IsNotOfPointerType)
	}

	// Si es un puntero, obtén el valor apuntado
	if s.reflectKind == reflect.Ptr {
		s.reflectValue = s.reflectValue.Elem()
		s.reflectKind = s.reflectValue.Kind()

		if !requiredPointer {
			s.buildName()

			// "structure " + name + " is not required as a pointer"
			// return errNoStructPtrReq(s.Name())
			return R.Err(D.TheStructure, s.Name(), D.IsNotRequired, D.AsAPointer)
		}

	}

	// Verifica si structIN es una estructura
	if s.reflectKind != reflect.Struct {
		return R.Err(D.TheElement, D.IsNotOfStructureType)
	}

	s.buildName()

	return nil
}
