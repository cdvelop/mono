# MonoGO

framework de desarrollo fullstack en go, 

Define una estructura de datos una vez y genera automáticamente formularios HTML, API HTTP y consultas SQL.

con soporte a webAssembly y tinygo.
- # ! ADVERTENCIA API EN DESARROLLO ! 

# ejemplo
```go	
package main

```



### por que no solo unir json gorm?
json de la librería estándar y gorm usan reflect y eso si lo llevamos a un dispositivo de bajo rendimiento se vuelve lento, si a eso le sumamos que queremos compilar a tinygo (que es la única forma de reducir el tamaño del binario resultante para webAssembly) necesitamos que el resultado sea optimo, sin dependencias externas y soporte a webAssembly + tinygo.

### monogo no usa reflect?
si lo usa pero una sola vez es como un json con asteroides. crea una imagen de cada estructura de tu programa cuando este arranca y lo mantiene en memoria para que no tenga que ser generado en tiempo de ejecución.


es un proyecto grande si, llevo trabando en el desde el 2020 pero con librerías separadas. ahora con todo lo aprendido es el momento de unir todo.

## alcances:

### html (gui)
- solo se pretende la generación de formularios básicos sin estilos css
- [ ] Generación de formularios HTML
- [ ] Validación en el lado del cliente (wasm)
- [ ] Validación en el lado del servidor (Go)

### base de datos
solo se trabajara con almacenamiento del tipo texto  que es el que todas las base de datos soportan, por ende la validación de datos se hará en el lado del servidor y el front con go + webAssembly. asi se evita configuraciones interminables en las estructuras de datos.

### json
- [ ] Generación de json


manejadores:

- ui: renderizados de formularios html
- net: api http
- db: sql dinámico
- json

## html rendering and validation

según input agregado en la etiqueta

- si el campo no contiene el atributo Input se le asigna un input de tipo text a todo tipo de dato excepto si es numérico como int, uint etc.
```go
type Person struct {
	Id    string `Legend:"Id Persona"`
	other uint 
	Name  string `Legend:"Nombre" Input:"Text()"`
	Age   uint8  `Input:"Number(min=0;max=120)"` 
}
```
- tipos de inputs soportados en ..monogo/inputs eg:
- si el o los inputs que necesitas no se encuentra en la librería puedes crear el método asociado a tu estructura de esta forma:
```go
type input interface {
	Render(tabIndex int) string
	Validate(value string) error
}
type Person struct {
	Name string `Legend:"Nombre" Input:"NewInputText()"`
}

func (p *Person) Inputs() []input {
	return p.myInputs
}


```
 un crear uno nuevo y agregarlo a la librería.

- params support with = eg: `Input:"Radio(options=m:male,f:female)"`:
    min,max,maxlength,class,placeholder,title,autocomplete,rows,cols,step,oninput,onkeyup,onchange,accept,multiple,value,options

- y si necesito que la data del input sea dinámica?
 ej: `Input:"List(dynamic)"`
	el nombre de la función debe ser igual al nombre del campo ej: `functionName(id string) []map[string]string`

- name se asigna de forma automática según el nombre del campo ej: `Age uint8 = name=age`

- dataset eg: `Input:"Text(data=price:100,dto:15%)"` html: `<input type="text" data-price="100" data-dto="15%">`

- !required (campo no obligatorio) eg: `Input:"Text(!required)"` por defecto todos los campos son obligatorios.

- hidden (campo oculto) eg: `Input:"Text(hidden)"`
- !required (campo no obligatorio) ej: `Input:"Text(!required)"` por defecto todos los campos son obligatorios.
- typing=hide (ocultar el campo cuando se escribe eg: password) ej: `Input:"Text(typing=hide)"`
- letters (letras permitidas) ej: `Input:"Text(letters)"`
- numbers (números permitidos) ej: `Input:"Text(numbers)"`
- chars (caracteres permitidos) ej : `Input:"Text(chars=':','-','+',' ')"`

* para combinar separar con ; ej: `Input:"Text(data=price:100;!required letters;numbers)"`

## Participar
si quieres participar en el proyecto puedes contactarme con un mensaje privado 


## Apóyame

Si encuentras útil este proyecto y te gustaría apoyarlo, puedes hacer una donación [aquí con paypal](https://paypal.me/cdvelop?country.x=CL&locale.x=es_XC)

Cualquier contribución, por pequeña que sea, es muy apreciada. 🙌