package markdown

import "strings"

// Heading Entity
type Heading struct {
	Indentation int
	ID          string
	Text        *FormatedText
}

// CreateHeader Creates a Markdown Document Header
func (d *Document) CreateHeader() *Heading {
	formatedText := FormatedText{}
	heading := Heading{
		Indentation: 1,
		Text:        &formatedText,
	}

	d.Elements = append(d.Elements, &heading)
	return &heading
}

// H1 Creates an H1 header
func (h *Heading) H1(value string) {
	h.Indentation = 1
	h.Text.Value = value
}

// H2 Creates an H2 header
func (h *Heading) H2(value string) {
	h.Indentation = 2
	h.Text.Value = value
}

// H3 Creates an H3 header
func (h *Heading) H3(value string) {
	h.Indentation = 3
	h.Text.Value = value
}

// Markdown Generates the Code Block Markdown
func (h *Heading) Markdown() string {
	if h.Indentation < 1 {
		h.Indentation = 1
	}
	result := ""
	for i := 0; i < h.Indentation; i++ {
		result += "#"
	}
	result += " " + h.Text.Markdown()
	if h.ID != "" {
		id := strings.ReplaceAll(h.ID, " ", "-")
		id = strings.ToLower(id)
		result += " {#" + id + "}"
	}
	result += "\n\n"

	return result
}
