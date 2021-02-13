package markdown

import "fmt"

// Element Markdown Document Elements interface
type Element interface {
	Markdown() string
}

// Document Markdown document Entity
type Document struct {
	Elements []Element
}

// CreateDocument Creates a new Markdown document
func CreateDocument() *Document {
	doc := Document{
		Elements: make([]Element, 0),
	}

	return &doc
}

// Sprint Gets the Markdown document in a string
func (d *Document) Sprint() string {
	result := ""
	for _, element := range d.Elements {
		result += element.Markdown()
	}
	return result
}

// Print Prints the Markdown document to the console
func (d *Document) Print() {
	fmt.Print(d.Sprint())
}

// Add Adds an Markdown element to the document
func (d *Document) Add(element Element) {
	d.Elements = append(d.Elements, element)
}

// AddHorizontalRule Adds an Horizontal Rule to the document
func (d *Document) AddHorizontalRule() {
	d.Elements = append(d.Elements, CreateHorizontalRule())
}

func insertSpaces(number int) string {
	result := ""
	for i := 0; i < number; i++ {
		result += " "
	}
	return result
}

func insertChar(number int, value string) string {
	result := ""
	for i := 0; i < number; i++ {
		result += value
	}
	return result
}
