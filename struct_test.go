package godi

import (
	"testing"
)

func TestValidateStruct(t *testing.T) {

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
			errorExpected:  errNoStructPtr("Person").Error(),
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
			errorExpected:  errNoStructPtrReq("Person").Error(),
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
			errorExpected:   errNoStructPtr(unknownStName).Error(),
		},
		{
			name:           "map sent, error expected",
			input:          map[string]int{"data": 1},
			requirePointer: false,
			expectPointer:  false,
			errorExpected:  errNotStruct("map").Error(),
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
				t.Errorf("expected error %v, got %v", tc.errorExpected, errStr)
			}

			if tc.expectAnonymous != response.anonymous {
				t.Errorf("expected anonymous %v, got %v", tc.expectAnonymous, response.anonymous)
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
