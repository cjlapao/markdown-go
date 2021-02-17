package markdown

// Languages Code Block Languages Enum
type Languages int

// Languages Enum Definition
const (
	Powershell Languages = iota
	Bash
	Yaml
	JSON
)

func (l Languages) String() string {
	switch l {
	case Powershell:
		return "powershell"
	case Bash:
		return "bash"
	case Yaml:
		return "yaml"
	case JSON:
		return "json"
	default:
		return ""
	}
}

// CodeBlock Entity
type CodeBlock struct {
	Language Languages
	Lines    []string
}

// CreateCodeBlock Creates a Markdown Code Block
func (d *Document) CreateCodeBlock() *CodeBlock {
	codeBlock := CodeBlock{
		Lines: make([]string, 0),
	}

	d.Elements = append(d.Elements, &codeBlock)

	return &codeBlock
}

// AddLine Adds a new line to a Markdown Code Block
func (t *CodeBlock) AddLine(line string) *CodeBlock {
	t.Lines = append(t.Lines, line)
	return t
}

// Markdown Generates the Code Block Markdown
func (t *CodeBlock) Markdown() string {
	result := "```"
	result += t.Language.String()
	result += "\n"
	for _, line := range t.Lines {
		result += line + "\n"
	}
	result += "```\n\n"

	return result
}
