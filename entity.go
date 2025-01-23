package mono

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

type entity struct {
	Name string //ej: users.user, module.product
	// Legend        string // e.g.: Person, User, Product
	IsTable   bool
	TableName string //table name db ej: user, product
	// ParentStruct any
	Fields []field
	// StructHandler *structHandler

	HtmlForm string //html form
}

func CreateEntityFromStruct(structIN ...any) []entity {
	var entities []entity
	for _, s := range structIN {
		processStruct(reflect.TypeOf(s), &entities, nil)
	}
	return entities
}

func processStruct(t reflect.Type, entities *[]entity, foreignField *field) {

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	e := entity{
		Name:      strings.ToLower(t.String()), // the package name + "." + structure name e.g.: attention.person
		TableName: G.String.SnakeCase(t.Name()),
		Fields:    []field{},
		// ParentStruct: parentStruct,
	}

	fmt.Println("processing struct:", t.Name())

	for i := 0; i < t.NumField(); i++ {
		// skip if it starts with lowercase
		if !unicode.IsUpper([]rune(t.Field(i).Name)[0]) {
			continue
		}

		rsField := t.Field(i)

		newField := field{
			Index:  uint32(len(e.Fields)),
			Name:   G.String.SnakeCase(t.Field(i).Name),
			Unique: setUnique(&rsField),
			Parent: &e,
		}

		newField.setDataBaseParams()

		var structureFrom reflect.Type

		// skip unsupported fields
		switch rsField.Type.Kind() {

		case reflect.String:

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		case reflect.Float32, reflect.Float64:

		case reflect.Bool:

		case reflect.Slice:
			// Handle one-to-many relationship
			if rsField.Type.Elem().Kind() == reflect.Struct {
				structureFrom = rsField.Type.Elem()
				processStruct(rsField.Type.Elem(), entities, &newField)
				// continue // Skip adding this field to the current table
			}

		case reflect.Struct:
			// Handle one-to-many relationship
			structureFrom = rsField.Type.Elem()
			processStruct(rsField.Type, entities, &newField)
			// continue // Skip adding this field to the current table

		default:
			fmt.Println("unsupported type:", rsField.Type.Kind())
			continue

		}
		newField.setLegend(&rsField)

		newField.setInput(structureFrom, &rsField)

		e.Fields = append(e.Fields, newField)

	}

	// Add foreign key for one-to-many relationship
	if foreignField != nil {

		fmt.Println("foreignField:", foreignField.Parent.TableName)

		e.Fields = append(e.Fields, field{
			Index:      uint32(len(e.Fields)),
			Name:       prefixNameID + foreignField.Parent.TableName,
			Legend:     "Id",
			PrimaryKey: false,
			NotNull:    true,
			ForeignKey: foreignField.Parent,
			Input:      IN.ID(),
			Parent:     &e,
		})

	}

	*entities = append(*entities, e)
}
