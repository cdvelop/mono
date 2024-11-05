package css

import (
	"fmt"
	"strings"
	"testing"
)

func TestCSSGeneration(t *testing.T) {
	// Caso de prueba 1: Generación básica de CSS
	t.Run("Basic CSS Generation", func(t *testing.T) {
		stylesheet := &Stylesheet{}

		buttonClass := NewCSSClass("btn-primary").
			AddProperty("font-family", "Arial", "Helvetica", "sans-serif").
			AddProperty("background", "linear-gradient(to right)", "blue", "purple").
			AddProperty("padding", "10px", "15px").
			AddProperty("box-shadow", "0 2px 4px rgba(0,0,0,0.1)")

		cardClass := NewCSSClass("card").
			AddProperty("border", "1px", "solid", "#ccc").
			AddProperty("transition", "all", "0.3s", "ease-in-out")

		stylesheet.AddClass(buttonClass).AddClass(cardClass)

		css := stylesheet.GenerateStylesheet()

		// Verificaciones de contenido
		expectedParts := []string{
			".btn-primary {",
			"font-family: Arial Helvetica sans-serif;",
			"background: linear-gradient(to right) blue purple;",
			"padding: 10px 15px;",
			"box-shadow: 0 2px 4px rgba(0,0,0,0.1);",
			".card {",
			"border: 1px solid #ccc;",
			"transition: all 0.3s ease-in-out;",
		}

		for _, part := range expectedParts {
			if !strings.Contains(css, part) {
				t.Errorf("Generated CSS missing expected part: %s", part)
			}
		}
	})

	// Caso de prueba 2: Verificación de orden de propiedades
	t.Run("Property Order Preservation", func(t *testing.T) {
		Class := NewCSSClass("test-order")

		// Propiedades en un orden específico
		Class.AddProperty("display", "flex")
		Class.AddProperty("justify-content", "center")
		Class.AddProperty("align-items", "center")

		css := Class.GenerateCSS()

		// Verificar orden de las propiedades
		cssLines := strings.Split(css, "\n")

		expectedOrder := []string{
			".test-order {",
			"    display: flex;",
			"    justify-content: center;",
			"    align-items: center;",
			"}",
		}

		for i, expectedLine := range expectedOrder {
			if i < len(cssLines) && strings.TrimSpace(cssLines[i]) != strings.TrimSpace(expectedLine) {
				t.Errorf("Unexpected line order. Expected %s, got %s",
					expectedLine, cssLines[i])
			}
		}
	})

	// Caso de prueba 3: Múltiples valores
	t.Run("Multiple Values", func(t *testing.T) {
		Class := NewCSSClass("multi-value")

		Class.AddProperty("background", "linear-gradient(45deg)", "red", "blue")

		css := Class.GenerateCSS()

		expectedValue := "background: linear-gradient(45deg) red blue;"
		if !strings.Contains(css, expectedValue) {
			t.Errorf("Failed to generate correct multiple value property. Got: %s", css)
		}
	})

	// Caso de prueba 4: Stylesheet completo
	t.Run("Full Stylesheet Generation", func(t *testing.T) {
		stylesheet := &Stylesheet{}

		buttonClass := NewCSSClass("button").
			AddProperty("color", "white")

		cardClass := NewCSSClass("card").
			AddProperty("border", "1px", "solid", "black")

		stylesheet.AddClass(buttonClass)
		stylesheet.AddClass(cardClass)

		css := stylesheet.GenerateStylesheet()

		expectedClasses := []string{
			".button {",
			"color: white;",
			".card {",
			"border: 1px solid black;",
		}

		for _, expectedClass := range expectedClasses {
			if !strings.Contains(css, expectedClass) {
				t.Errorf("Stylesheet missing expected Class: %s", expectedClass)
			}
		}
	})
}

// Ejemplo de ejecución de benchmark para rendimiento
func BenchmarkCSSGeneration(b *testing.B) {
	stylesheet := &Stylesheet{}

	// Preparar un conjunto de clases
	for i := 0; i < 100; i++ {
		Class := NewCSSClass(fmt.Sprintf("Class-%d", i))
		Class.AddProperty("color", "black")
		Class.AddProperty("margin", "10px")
		stylesheet.AddClass(Class)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stylesheet.GenerateStylesheet()
	}
}
