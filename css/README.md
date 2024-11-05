
## Ejemplo de uso con múltiples valores:
```go
func main() {
    stylesheet := &Stylesheet{}

    // Ejemplo de propiedades con múltiples valores
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
    fmt.Println(css)
}
```
## Este código generaría un CSS como:
```css
.btn-primary {
    font-family: Arial Helvetica sans-serif;
    background: linear-gradient(to right) blue purple;
    padding: 10px 15px;
    box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}
.card {
    border: 1px solid #ccc;
    transition: all 0.3s ease-in-out;
}
```