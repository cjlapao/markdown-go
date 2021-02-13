package markdown

import "strings"

// FormatedText Entity
type FormatedText struct {
	Value           string
	IsBold          bool
	IsItalic        bool
	IsCode          bool
	IsStrikeThrough bool
}

// Markdown Generates the Code Block Markdown
func (ft *FormatedText) Markdown() string {
	result := ft.Value
	if ft.IsItalic && !ft.IsCode {
		result = "*" + result + "*"
	}
	if ft.IsBold && !ft.IsCode {
		result = "**" + result + "**"
	}
	if ft.IsStrikeThrough && !ft.IsCode {
		result = "~~" + result + "~~"
	}
	if ft.IsCode {
		result = "`" + result + "`"
	}

	return result
}

// Sprint Returns a String representation of the Markdown format
func (ft *FormatedText) Sprint() string {
	return ft.Markdown()
}

// TextBlock Entity
type TextBlock struct {
	Words []FormatedText
}

// CreateTextBlock Creates a TextBlock Element
func CreateTextBlock() *TextBlock {
	txtBlock := TextBlock{
		Words: make([]FormatedText, 0),
	}
	return &txtBlock
}

// NewLine Adds a new line termination to the TextBlock Element
func (t *TextBlock) NewLine() *TextBlock {
	line := FormatedText{
		Value: "\n\n",
	}
	t.Words = append(t.Words, line)
	return t
}

// Add Adds Words or strings to the TextBlock Element
func (t *TextBlock) Add(words ...string) *TextBlock {
	for i, word := range words {
		if i == len(words)-1 {
			strings.TrimSpace(word)
		}
		f := FormatedText{
			Value: word,
		}
		t.Words = append(t.Words, f)
	}

	return t
}

// AddBold Add a bold string to the TextBlock Element
func (t *TextBlock) AddBold(value string) *TextBlock {
	f := FormatedText{
		Value:  value,
		IsBold: true,
	}
	t.Words = append(t.Words, f)

	return t
}

// AddItalic Add a italic string to the TextBlock Element
func (t *TextBlock) AddItalic(value string) *TextBlock {
	f := FormatedText{
		Value:    value,
		IsItalic: true,
	}
	t.Words = append(t.Words, f)

	return t
}

// AddBoldItalic Add a bold and italic string to the TextBlock Element
func (t *TextBlock) AddBoldItalic(value string) *TextBlock {
	f := FormatedText{
		Value:    value,
		IsItalic: true,
		IsBold:   true,
	}
	t.Words = append(t.Words, f)

	return t
}

// AddStrikedBold Add a bold strikethrough string to the TextBlock Element
func (t *TextBlock) AddStrikedBold(value string) *TextBlock {
	f := FormatedText{
		Value:           value,
		IsBold:          true,
		IsStrikeThrough: true,
	}
	t.Words = append(t.Words, f)

	return t
}

// AddStrikedItalic Add a italic strikethrough string to the TextBlock Element
func (t *TextBlock) AddStrikedItalic(value string) *TextBlock {
	f := FormatedText{
		Value:           value,
		IsItalic:        true,
		IsStrikeThrough: true,
	}
	t.Words = append(t.Words, f)

	return t
}

// AddStrikedBoldItalic Add a bold italic strikethrough string to the TextBlock Element
func (t *TextBlock) AddStrikedBoldItalic(value string) *TextBlock {
	f := FormatedText{
		Value:           value,
		IsItalic:        true,
		IsBold:          true,
		IsStrikeThrough: true,
	}
	t.Words = append(t.Words, f)

	return t
}

// AddCode Add a code string to the TextBlock Element
func (t *TextBlock) AddCode(value string) *TextBlock {
	f := FormatedText{
		Value:  value,
		IsCode: true,
	}
	t.Words = append(t.Words, f)

	return t
}

// AddLine Adds a string with new line termination to the TextBlock Element
func (t *TextBlock) AddLine(words ...string) *TextBlock {
	line := FormatedText{
		Value: "\n",
	}
	t.Add(words...)
	if !strings.HasSuffix(t.Words[len(t.Words)-1].Value, "\n") {
		line.Value += "\n"
	}
	t.Words[len(t.Words)-1].Value = strings.TrimSpace(t.Words[len(t.Words)-1].Value)
	t.Words = append(t.Words, line)
	return t
}

// Sprint Returns a String representation of the Markdown format
func (t *TextBlock) Sprint() string {
	return t.Markdown()
}

// Markdown Generates the Code Block Markdown
func (t *TextBlock) Markdown() string {
	result := ""
	for i, word := range t.Words {
		if i > 0 {
			if !strings.HasSuffix(t.Words[len(t.Words)-1].Value, "\n") {
				result += " "
			}
		}
		result += word.Markdown()
	}

	if !strings.HasSuffix(result, "\n\n") {
		result += "\n"
	}

	return result
}
