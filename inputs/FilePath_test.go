package inputs

import (
	"log"
	"testing"
)

func Test_FilePath(t *testing.T) {

	var (
		filePathTestData = map[string]struct {
			inputData string
			expected  string
		}{
			"correct path":                                {".\\files\\1234\\", ""},
			"correct path without starting point?":        {"\\files\\1234\\", Lang.Err(D.DoNotStartWith, '\\').Error()},
			"absolute path in Linux":                      {"/home/user/files/", ""},
			"absolute path without starting point":        {"/files/1234/", ""},
			"relative path without directories ok?":       {".\\", ""},
			"relative path with final slash":              {"./files/1234/", ""},
			"path with filename":                          {".\\files\\1234\\archivo.txt", ""},
			"path with directory names using underscores": {".\\mi_directorio\\sub_directorio\\", ""},
			"is a number a valid path?":                   {"5", ""},
			"is a single word a valid path?":              {"path", ""},
			"white spaces allowed?":                       {".\\path with white space\\", Lang.T(D.WhiteSpace, D.NotAllowed)},
		}
	)

	for prueba, data := range filePathTestData {
		t.Run((prueba + " " + data.inputData), func(t *testing.T) {
			err := FilePath().Validate(data.inputData)

			var err_str string
			if err != nil {
				err_str = err.Error()
			}

			if err_str != data.expected {
				log.Println(prueba)
				log.Fatalf("got: [%v] expected: [%v]\n%v", err, data.expected, data.inputData)
			}
		})
	}
}

// func Test_TagFilePath(t *testing.T) {
// 	tag := input.FilePath().Render("1", "name", true)
// 	if tag == "" {
// 		log.Fatalln("ERROR NO TAG RENDERING ")
// 	}
// }

// func Test_GoodInputFilePath(t *testing.T) {
// 	for _, data := range input.FilePath().GoodTestData() {
// 		t.Run((data), func(t *testing.T) {
// 			if ok := input.FilePath().Validate(data, false); ok != nil {
// 				log.Fatalf("resultado [%v] [%v]", ok, data)
// 			}
// 		})
// 	}
// }

// func Test_WrongInputFilePath(t *testing.T) {
// 	for _, data := range input.FilePath().WrongTestData() {
// 		t.Run((data), func(t *testing.T) {
// 			if ok := input.FilePath().Validate(data, false); ok == nil {
// 				log.Fatalf("resultado [%v] [%v]", ok, data)
// 			}
// 		})
// 	}
// }
