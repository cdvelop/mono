package godi

import (
	"fmt"
	"reflect"
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
	return []*Address{
		{Id: "1", Street: "123 Main St", City: "New York", ZipCode: "10001"},
		{Id: "2", Street: "456 Elm St", City: "Los Angeles", ZipCode: "90001"},
		{Id: "3", Street: "789 Oak St", City: "Chicago", ZipCode: "60601"},
	}
}

type Person struct {
	Id        string    `Legend:"Id"`
	Name      string    `Legend:"Nombre" Input:"Text()"`
	Date      uint8     `Input:"Date()"`
	Gender    string    `Input:"RadioGender()"`
	Phone     string    `Input:"Phone()"`
	Addresses []Address `Legend:"Direcciones" Input:"List()"` //foreign key expected
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
			{Index: 1, Name: "name", Legend: "Nombre", Input: inputs.Text("name=name"), Parent: gotPerson},
			{Index: 2, Name: "date", Input: inputs.Date(), Parent: gotPerson},
			{Index: 3, Name: "gender", Input: inputs.RadioGender(), Parent: gotPerson},
			{Index: 4, Name: "phone", Input: inputs.Phone(), Parent: gotPerson},
			{Index: 5, Name: "addresses", Legend: "Direcciones", Input: inputs.List(structureAddressFrom, "name=addresses"), Parent: gotPerson},
		},
	}

	expectedAddress := &entity{
		Name:      "godi.address", // package name + struct name
		TableName: "address",
		IsTable:   true,
		// ReflectValue: gotAddress.ReflectValue,
		Fields: []field{
			{Index: 0, Name: "id_address", Legend: "Id", PrimaryKey: true, Unique: true, Input: inputs.ID(), Parent: gotAddress},
			{Index: 1, Name: "street", Legend: "Calle", Input: inputs.Text("name=street"), Parent: gotAddress},
			{Index: 2, Name: "city", Legend: "Ciudad", Input: inputs.Text("name=city"), Parent: gotAddress},
			{Index: 3, Name: "zip_code", Legend: "Código Postal", Input: inputs.Text("name=zip_code"), Parent: gotAddress},
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

	expectedForm := `<form>
 	<div>
 		<label for="id_person">Id</label>
 		<input type="text" id="id_person" name="id_person" readonly>
 	</div>
 	<div>
 		<label for="name">Nombre</label>
 		<input type="text" id="name" name="name">
 	</div>
 	<div>
 		<label for="date">Date</label>
 		<input type="date" id="date" name="date">
 	</div>
 	<div>
 		<label>Gender</label>
 		<input type="radio" id="gender_male" name="gender" value="male">
 		<label for="gender_male">Male</label>
 		<input type="radio" id="gender_female" name="gender" value="female">
 		<label for="gender_female">Female</label>
 	</div>
 	<div>
 		<label for="phone">Phone</label>
 		<input type="tel" id="phone" name="phone">
 	</div>
 	<div>
 		<label for="addresses">Direcciones</label>
 		<select id="addresses" name="addresses" multiple>
 			<option value="1">Address 1</option>
 			<option value="2">Address 2</option>
 			<option value="3">Address 3</option>
 		</select>
 	</div>
 </form>`

	fmt.Println("****")
	gotForm := expectedPerson.FormRender()

	if expectedForm != gotForm {
		// fmt.Println(gotAddress)
		t.Fatalf("\n❌ %s are not equal\n\nexpected:\n%v\n\n❌got:\n%v\n\n", expectedPerson.Name+" FORM", expectedForm, gotForm)
	}

}
