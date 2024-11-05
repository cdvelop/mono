package css

import (
	"fmt"
	"strings"
)

// Estructura de clase CSS que mantiene el orden
type Class struct {
	Name       string   //eg: "normal", "border", "width-auto"
	Properties []string //eg: "width: 100%", "min-width: 100%"
}

// Constructor de clase CSS
func NewCSSClass(name string) *Class {
	return &Class{
		Name:       name,
		Properties: []string{},
	}
}

// Método para agregar una propiedad CSS con múltiples valores
func (c *Class) AddProperty(key string, values ...string) *Class {
	property := fmt.Sprintf("%s: %s", key, strings.Join(values, " "))
	c.Properties = append(c.Properties, property)
	return c
}

// getClassName retorna el nombre de la clase para usar en HTML
func (c *Class) GetClassName() string {
	return fmt.Sprintf(".%s", c.Name)
}

// Método para generar el CSS respetando el orden de inserción
func (c *Class) GenerateCSS() string {
	var cssBuilder strings.Builder

	// Pre-estimar tamaño
	estimatedSize := len(c.Name) + 50
	for _, prop := range c.Properties {
		estimatedSize += len(prop) + 20
	}
	cssBuilder.Grow(estimatedSize)

	cssBuilder.WriteString(fmt.Sprintf(".%s {\n", c.Name))

	for _, prop := range c.Properties {
		cssBuilder.WriteString(fmt.Sprintf("    %s;\n", prop))
	}

	cssBuilder.WriteString("}\n")

	return cssBuilder.String()
}
