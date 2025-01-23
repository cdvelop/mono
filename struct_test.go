package mono

import (
	"testing"
)

func TestValidateStruct(t *testing.T) {

	const unknownStName = "unknown"

	type Person struct {
		ID   string
		Name string
		Age  uint8
	}

	testCases := []struct {
		name            string
		input           any
		requirePointer  bool
		expectPointer   bool
		expectAnonymous bool
		errorExpected   string
	}{
		{
			name:           "structure sent as parameter ok",
			input:          Person{},
			requirePointer: false,
			expectPointer:  true,
			errorExpected:  "",
		},
		{
			name:           "Person structure sent as parameter, but pointer is required",
			input:          Person{},
			requirePointer: true,
			expectPointer:  false,
			errorExpected:  R.T(D.TheStructure, "Person", D.IsNotOfPointerType),
		},
		{
			name:           "Person structure sent as required pointer ok",
			input:          &Person{},
			requirePointer: true,
			expectPointer:  true,
			errorExpected:  "",
		},
		{
			name:           "Person structure sent as pointer and not required as pointer",
			input:          &Person{},
			requirePointer: false,
			expectPointer:  false,
			errorExpected:  R.T(D.TheStructure, "Person", D.IsNotRequired, D.AsAPointer),
		},
		{
			name:            "anonymous structure sent as pointer ok",
			input:           &struct{ Name string }{Name: "john"},
			requirePointer:  true,
			expectPointer:   true,
			expectAnonymous: true,
			errorExpected:   "",
		},
		{
			name:            "anonymous structure sent as parameter ok",
			input:           struct{ Name string }{Name: "john"},
			requirePointer:  false,
			expectPointer:   true,
			expectAnonymous: true,
			errorExpected:   "",
		},
		{
			name:            "anonymous structure sent as parameter but pointer is required",
			input:           struct{ Name string }{Name: "john"},
			requirePointer:  true,
			expectPointer:   false,
			expectAnonymous: true,
			errorExpected:   R.T(D.TheStructure, unknownStName, D.IsNotOfPointerType),
		},
		{
			name:           "map sent, error expected",
			input:          map[string]int{"data": 1},
			requirePointer: false,
			expectPointer:  false,
			errorExpected:  R.T(D.TheElement, "map", D.IsNotOfStructureType),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			var response structHandler

			err := response.validate(tc.input, tc.requirePointer)
			var errStr string
			if err != nil {
				errStr = err.Error()
			}

			if errStr != tc.errorExpected {
				t.Fatalf("expected error %v, got %v", tc.errorExpected, errStr)
			}

			if tc.expectAnonymous != response.anonymous {
				t.Fatalf("expected anonymous %v, got %v", tc.expectAnonymous, response.anonymous)
			}
		})
	}
}

func BenchmarkValidateStruct(b *testing.B) {
	type Person struct {
		ID   string
		Name string
		Age  uint8
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var response structHandler
		response.validate(&Person{}, true)
	}
}
