package mono

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSnakeCase(t *testing.T) {

	testCases := []struct {
		input string
		want  string
	}{
		{"NameFirst", "name_first"},
		{"nameFirst", "name_first"},
		{"NameFirstSecond", "name_first_second"},
		{"_other_name", "_other_name"},
		{"123Test", "123_test"},
		{"", ""},
		{"ALLCAPS", "allcaps"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			got := G.String.SnakeCase(tc.input)
			if got != tc.want {
				t.Fatalf("snakeCase(%q) = %q; want %q", tc.input, got, tc.want)
			}
		})
	}
}

func compareFieldsInStruct(expected, got reflect.Value) error {
	t := expected.Type()

	for i := 0; i < expected.NumField(); i++ {
		field1 := expected.Field(i)
		field2 := got.Field(i)
		fieldName := t.Field(i).Name

		if field1.Kind() == reflect.Slice {
			if field1.Len() != field2.Len() {
				return fmt.Errorf("%s: slice lengths differ. Expected %d, got %d", fieldName, field1.Len(), field2.Len())
			}

			for j := 0; j < field1.Len(); j++ {
				elem1 := field1.Index(j)
				elem2 := field2.Index(j)

				if elem1.Kind() == reflect.Struct {
					if err := compareFieldsInStruct(elem1, elem2); err != nil {
						return fmt.Errorf("%s[%d]: %v", fieldName, j, err)
					}
				} else {
					if err := compareValues(elem1, elem2, fmt.Sprintf("%s[%d]", fieldName, j)); err != nil {
						return err
					}
				}
			}
		} else if field1.Kind() == reflect.Struct {
			if err := compareFieldsInStruct(field1, field2); err != nil {
				return fmt.Errorf("%s: %v", fieldName, err)
			}
		} else {
			if err := compareValues(field1, field2, fieldName); err != nil {
				return err
			}
		}
	}

	return nil
}

func compareValues(expected, got reflect.Value, fieldName string) error {
	if !reflect.DeepEqual(expected.Interface(), got.Interface()) {
		return fmt.Errorf("%s are not equal\n\nexpected:\n%v\n\ngot:\n%v\n\n", fieldName, expected.Interface(), got.Interface())
	}
	return nil
}
