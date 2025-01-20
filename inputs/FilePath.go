package inputs

// options:
// "multiple"
// accept="image/*"
// title="Imágenes jpg"
func FilePath(params ...any) *filePath {

	new := &filePath{
		input: input{
			attributes: attributes{
				htmlName:   "file",
				customName: "FilePath",
			},
			permitted: permitted{
				Letters:    true,
				Tilde:      false,
				Numbers:    true,
				Characters: []rune{'\\', '/', '.', '_'},
				Minimum:    1,
				Maximum:    100,
				StartWith: &permitted{Letters: true,
					Numbers:    true,
					Characters: []rune{'.', '_', '/'}},
			},
		},
	}
	new.Set(params)

	return new
}

type filePath struct {
	input
}

func (f filePath) GoodTestData() (out []string) {
	return []string{
		".\\misArchivos",
		".\\todos\\videos",
	}
}

func (f filePath) WrongTestData() (out []string) {
	out = []string{
		"\\-",
		"///.",
	}
	return
}
