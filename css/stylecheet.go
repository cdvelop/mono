package css

import "strings"

// Estructura para manejar múltiples clases
type Stylesheet struct {
	classes []*Class
}

// Método para agregar una clase al stylesheet
func (ss *Stylesheet) AddClass(Class *Class) *Stylesheet {
	// Verificar si la clase ya existe
	for _, existingClass := range ss.classes {
		if existingClass.Name == Class.Name {
			return ss
		}
	}
	ss.classes = append(ss.classes, Class)
	return ss
}

// Método para generar todo el stylesheet
// Usar un buffer pre-asignado para reducir allocations
func (ss *Stylesheet) GenerateStylesheet() string {

	// Usar strings.Builder con capacidad inicial
	var stylesheetBuilder strings.Builder

	stylesheetBuilder.WriteString(GenerateRoot())

	for _, Class := range ss.classes {
		stylesheetBuilder.WriteString(Class.GenerateCSS())
	}

	return stylesheetBuilder.String()
}
