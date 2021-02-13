package markdown

// HorizontalRule Entity
type HorizontalRule struct{}

// CreateHorizontalRule Creates an HorizontalRule Element
func CreateHorizontalRule() *HorizontalRule {
	horizontalRule := HorizontalRule{}
	return &horizontalRule
}

// Markdown Generates the Code Block Markdown
func (h *HorizontalRule) Markdown() string {
	result := "\n---\n"
	return result
}
