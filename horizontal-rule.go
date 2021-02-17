package markdown

// HorizontalRule Entity
type HorizontalRule struct{}

// CreateHorizontalRule Creates an HorizontalRule Element
func (d *Document) CreateHorizontalRule() *HorizontalRule {
	horizontalRule := HorizontalRule{}

	d.Elements = append(d.Elements, &horizontalRule)
	return &horizontalRule
}

// Markdown Generates the Code Block Markdown
func (h *HorizontalRule) Markdown() string {
	result := "\n---\n"
	return result
}
