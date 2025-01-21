package mono

import (
	"reflect"
	"strings"
	"testing"

	"github.com/cdvelop/mono/inputs"
)

type Address struct {
	Id      string
	Street  string `Legend:"Calle" Input:"Text()"`
	City    string `Legend:"Ciudad" Input:"Text()"`
	ZipCode string `Legend:"Código Postal" Input:"Text()"`
}

func (a *Address) DataSource() any {
	// database connected or mock eg:
	return []*Address{
		{Id: "1", Street: "123 Main St", City: "New York", ZipCode: "10001"},
		{Id: "2", Street: "456 Elm St", City: "Los Angeles", ZipCode: "90001"},
		{Id: "3", Street: "789 Oak St", City: "Chicago", ZipCode: "60601"},
	}
}

type Person struct {
	Id        string
	Name      string
	BirthDate string
	Gender    string
	Phone     string
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
		Name:      "mono.person", // package name + struct name
		TableName: "person",
		IsTable:   true,
		// ReflectValue: gotPerson.ReflectValue,
		Fields: []field{
			{Index: 0, Name: "id", Legend: "Id", PrimaryKey: true, Unique: true, Input: inputs.ID("entity=mono.person", "legend=Id"), Parent: gotPerson},
			{Index: 1, Name: "name", Legend: "Nombre", Input: inputs.Text("name=name", "entity=mono.person", "legend=Nombre"), Parent: gotPerson},
			{Index: 2, Name: "birth_date", Legend: "Fecha De Nacimiento", Input: inputs.Date("name=birth_date", "entity=mono.person", "legend=Fecha De Nacimiento"), Parent: gotPerson},
			{Index: 3, Name: "gender", Legend: "Género", Input: inputs.RadioGender("entity=mono.person", "legend=Género"), Parent: gotPerson},
			{Index: 4, Name: "phone", Legend: "Teléfono", Input: inputs.Phone(), Parent: gotPerson},
			{Index: 5, Name: "addresses", Legend: "Direcciones", Input: inputs.Select(structureAddressFrom, "name=addresses", "entity=mono.person", "legend=Direcciones"), Parent: gotPerson},
		},
	}

	expectedAddress := &entity{
		Name:      "mono.address", // package name + struct name
		TableName: "address",
		IsTable:   true,
		// ReflectValue: gotAddress.ReflectValue,
		Fields: []field{
			{Index: 0, Name: "id", Legend: "Id", PrimaryKey: true, Unique: true, Input: inputs.ID("entity=mono.address", "legend=Id"), Parent: gotAddress},
			{Index: 1, Name: "street", Legend: "Calle", Input: inputs.Text("name=street", "entity=mono.address", "legend=Calle"), Parent: gotAddress},
			{Index: 2, Name: "city", Legend: "Ciudad", Input: inputs.Text("name=city", "entity=mono.address", "legend=Ciudad"), Parent: gotAddress},
			{Index: 3, Name: "zip_code", Legend: "Código Postal", Input: inputs.Text("name=zip_code", "entity=mono.address", "legend=Código Postal"), Parent: gotAddress},
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
	formOriginal := entity.FormRender()

	formWithOutSpaces := strings.ReplaceAll(formOriginal, " ", "")

	// Check form opening
	expectedOpenOriginal := `<form name="mono.person" class="form-distributed-fields" autocomplete="off" spellcheck="false">`

	expectedOpenWithOutSpaces := strings.ReplaceAll(expectedOpenOriginal, " ", "")

	if !strings.Contains(formWithOutSpaces, expectedOpenWithOutSpaces) {
		t.Fatalf("\nIncorrect form opening\nExpected: %v\nGot: %v", expectedOpenOriginal, formOriginal[:len(expectedOpenOriginal)])
	}

	// Check form closing
	if !strings.HasSuffix(formWithOutSpaces, `</form>`) {
		t.Fatal("Incorrect form closing")
	}

	// Check each field
	for index, field := range entity.Fields {
		expectedTagFieldOriginal := field.Input.Render(index)

		expectedTagFieldWithOutSpaces := strings.ReplaceAll(expectedTagFieldOriginal, " ", "")

		if !strings.Contains(formWithOutSpaces, expectedTagFieldWithOutSpaces) {
			t.Fatalf("\nIncorrect field: %s\nExpected:\n%v\n❌Form:\n%v", field.Name, expectedTagFieldOriginal, formOriginal)
		}
	}
}
