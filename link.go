package markdown

// LinkType Link Type Enum
type LinkType int

// LinkType Enum Definition
const (
	URL LinkType = iota
	Image
)

// Link Entity
type Link struct {
	Type LinkType
	Text string
	URL  string
}

// CreateURLLink Creates a Markdown URL Link Element
func (d *Document) CreateURLLink() *Link {
	link := Link{
		Type: URL,
	}

	d.Elements = append(d.Elements, &link)
	return &link
}

// CreateImgLink Creates a Markdown Image Link Element
func (d *Document) CreateImgLink() *Link {
	link := Link{
		Type: Image,
	}

	d.Elements = append(d.Elements, &link)
	return &link
}

// Markdown Generates the Code Block Markdown
func (l *Link) Markdown() string {
	result := ""
	if l.Text != "" && l.URL != "" {
		if l.Type == Image {
			result += "!"
		}

		result += "[" + l.Text + "](" + l.URL + ")\n\n"
	}

	return result
}
