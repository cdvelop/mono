package inputs

import (
	"log"
	"testing"
)

var (
	filePathTestData = map[string]struct {
		inputData string
		expected  string
	}{
		// "tres rutas separadas por comas":       {`.\\files\\1234\\,.\\files\\5678\\,.\\images\\ok\\`, "false"},
		"dirección correcta":                               {".\\files\\1234\\", ""},
		"dirección correcta sin punto inicio?":             {"\\files\\1234\\", errPath.Error()},
		"ruta relativa con directorios":                    {".\\files\\1234\\", ""},
		"ruta relativa sin punto de inicio":                {"\\files\\1234\\", errPath.Error()},
		"ruta absoluta en Linux":                           {"/home/user/files/", ""},
		"ruta absoluta sin punto de inicio":                {"/files/1234/", ""},
		"ruta relativa sin directorios ok?":                {".\\", ""},
		"ruta relativa sin barra final":                    {"./files/1234", ""},
		"ruta relativa con barra final":                    {"./files/1234/", ""},
		"ruta con nombre de archivo":                       {".\\files\\1234\\archivo.txt", ""},
		"ruta con nombres de directorio con guiones bajos": {".\\mi_directorio\\sub_directorio\\", ""},
		"un numero es una ruta valida?":                    {"5", ""},
		"una sola palabra es una ruta valida?":             {"ruta", ""},
		"espacios en blanco permitidos?":                   {".\\ruta con espacio en blanco\\", errorWhiteSpace},
	}
)

func Test_Check(t *testing.T) {

	for prueba, data := range filePathTestData {
		t.Run((prueba + " " + data.inputData), func(t *testing.T) {
			err := FilePath().ValidateInput(data.inputData)

			var err_str string
			if err != nil {
				err_str = err.Error()
			}

			if err_str != data.expected {
				log.Println(prueba)
				log.Fatalf("resultado: [%v] expectativa: [%v]\n%v", err, data.expected, data.inputData)
			}
		})
	}
}

// func Test_TagFilePath(t *testing.T) {
// 	tag := input.FilePath().BuildHtmlInput("1", "name", true)
// 	if tag == "" {
// 		log.Fatalln("ERROR NO TAG RENDERING ")
// 	}
// }

// func Test_GoodInputFilePath(t *testing.T) {
// 	for _, data := range input.FilePath().GoodTestData() {
// 		t.Run((data), func(t *testing.T) {
// 			if ok := input.FilePath().ValidateInput(data, false); ok != nil {
// 				log.Fatalf("resultado [%v] [%v]", ok, data)
// 			}
// 		})
// 	}
// }

// func Test_WrongInputFilePath(t *testing.T) {
// 	for _, data := range input.FilePath().WrongTestData() {
// 		t.Run((data), func(t *testing.T) {
// 			if ok := input.FilePath().ValidateInput(data, false); ok == nil {
// 				log.Fatalf("resultado [%v] [%v]", ok, data)
// 			}
// 		})
// 	}
// }
