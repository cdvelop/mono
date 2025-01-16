package godi

import (
	"reflect"
	"strings"
	"testing"

	"github.com/cdvelop/godi/inputs"
)

type Address struct {
	Id      string `Legend:"Id"`
	Street  string `Legend:"Calle" Input:"Text()"`
	City    string `Legend:"Ciudad" Input:"Text()"`
	ZipCode string `Legend:"Código Postal" Input:"Text()"`
}

func (a *Address) DataSource() any {
	// database connected or mock eg:
	options := make(map[string]string)
	options["1"] = "123 Main St, New York"
	options["2"] = "456 Elm St, Los Angeles"
	options["3"] = "789 Oak St, Chicago"
	return options
}

type Person struct {
	Id        string    `Legend:"Id"`
	Name      string    `Legend:"Nombre" Input:"Text()"`
	Date      uint8     `Input:"Date()"`
	Gender    string    `Input:"RadioGender()"`
	Phone     string    `Input:"Phone()"`
	Addresses []Address `Legend:"Direcciones" Input:"Select()"` //foreign key expected
}

func TestBuildEntity(t *testing.T) {
	got := CreateEntityFromStruct(&Person{})

	if got == nil {
		t.Fatalf("\n❌Entity is nil")
	}

	if len(got) != 2 {
		t.Fatalf("\n❌Entity length is not 2")
	}

	gotAddress := &got[0]
	gotPerson := &got[1]

	structureAddressFrom := reflect.TypeOf(Address{})

	expectedPerson := &entity{
		Name:      "godi.person", // package name + struct name
		TableName: "person",
		IsTable:   true,
		// ReflectValue: gotPerson.ReflectValue,
		Fields: []field{
			{Index: 0, Name: "id_person", Legend: "Id", PrimaryKey: true, Unique: true, Input: inputs.ID(), Parent: gotPerson},
			{Index: 1, Name: "name", Legend: "Nombre", Input: inputs.Text("name=name", "entity=godi.person", "legend=Nombre"), Parent: gotPerson},
			{Index: 2, Name: "date", Input: inputs.Date("entity=godi.person"), Parent: gotPerson},
			{Index: 3, Name: "gender", Input: inputs.RadioGender("entity=godi.person"), Parent: gotPerson},
			{Index: 4, Name: "phone", Input: inputs.Phone(), Parent: gotPerson},
			{Index: 5, Name: "addresses", Legend: "Direcciones", Input: inputs.Select(structureAddressFrom, "name=addresses", "entity=godi.person", "legend=Direcciones"), Parent: gotPerson},
		},
	}

	expectedAddress := &entity{
		Name:      "godi.address", // package name + struct name
		TableName: "address",
		IsTable:   true,
		// ReflectValue: gotAddress.ReflectValue,
		Fields: []field{
			{Index: 0, Name: "id_address", Legend: "Id", PrimaryKey: true, Unique: true, Input: inputs.ID(), Parent: gotAddress},
			{Index: 1, Name: "street", Legend: "Calle", Input: inputs.Text("name=street", "entity=godi.address", "legend=Calle"), Parent: gotAddress},
			{Index: 2, Name: "city", Legend: "Ciudad", Input: inputs.Text("name=city", "entity=godi.address", "legend=Ciudad"), Parent: gotAddress},
			{Index: 3, Name: "zip_code", Legend: "Código Postal", Input: inputs.Text("name=zip_code", "entity=godi.address", "legend=Código Postal"), Parent: gotAddress},
			{Index: 4, Name: "id_person", Legend: "Id", NotNull: true, ForeignKey: expectedPerson, Input: inputs.ID(), Parent: gotAddress},
		},
	}
	// fmt.Println("expectedPerson", expectedPerson)

	if err := compareFieldsInStruct(reflect.ValueOf(expectedPerson).Elem(), reflect.ValueOf(gotPerson).Elem()); err != nil {
		t.Fatalf("\n❌Entity Person %v", err)
	}

	if err := compareFieldsInStruct(reflect.ValueOf(expectedAddress).Elem(), reflect.ValueOf(gotAddress).Elem()); err != nil {
		// fmt.Println(gotAddress)
		t.Fatalf("\n❌Entity Address %v", err)
	}

	compareFormParts(t, expectedPerson)
}

func compareFormParts(t *testing.T, entity *entity) {
	form := entity.FormRender()

	// Check form opening
	expectedOpen := `<form name="godi.person" class="form-distributed-fields" autocomplete="off" spellcheck="false">`
	if !strings.Contains(form, expectedOpen) {
		t.Fatalf("\nIncorrect form opening\nExpected: %v\nGot: %v", expectedOpen, form[:len(expectedOpen)])
	}

	// Check form closing
	expectedClose := `</form>`
	if !strings.HasSuffix(form, expectedClose) {
		t.Fatal("Incorrect form closing")
	}

	// Check each field
	for index, field := range entity.Fields {
		fieldHTML := field.Input.Render(index)
		if !strings.Contains(form, fieldHTML) {
			t.Fatalf("\nIncorrect field: %s\nExpected:\n%v\n", field.Name, fieldHTML)
		}
	}
}
