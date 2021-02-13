package markdown

import (
	"errors"
)

// ColumnAlignment Enum
type ColumnAlignment int

// ColumnAlignment Enum Definition
const (
	AlignLeft ColumnAlignment = iota
	AlignRight
	AlignCenter
)

// Column Entity
type Column struct {
	Value     string
	Size      int
	Alignment ColumnAlignment
}

// GetSize Gets the size of the Column base on the content
func (c Column) GetSize() int {
	return len(c.Value) + 2
}

// Row Entity
type Row struct {
	Columns []*Column
}

// AddColumn Adds a Column to a Row
func (r *Row) AddColumn(value string) *Column {
	column := Column{Value: value}
	column.Size = column.GetSize()
	r.Columns = append(r.Columns, &column)
	return &column
}

// Table Entity
type Table struct {
	Header *Row
	Body   []*Row
}

// CreateTable Creates a Table Element
func CreateTable() *Table {
	headerRow := Row{
		Columns: make([]*Column, 0),
	}

	table := Table{
		Header: &headerRow,
		Body:   make([]*Row, 0),
	}

	return &table
}

// AddHeaderColumn Adds a Column to the table header
func (t *Table) AddHeaderColumn(value string) {
	column := t.Header.AddColumn(value)
	column.Alignment = AlignLeft
}

// AddAlignedHeaderColumn Adds an Aligned Column to the table header
func (t *Table) AddAlignedHeaderColumn(value string, alignment ColumnAlignment) {
	column := t.Header.AddColumn(value)
	column.Alignment = alignment
}

// AddRow Adds a Row to the table
func (t *Table) AddRow(columns ...string) error {
	if len(columns) != len(t.Header.Columns) {
		return errors.New("The number of columns need to be the same as the table header")
	}
	row := Row{
		Columns: make([]*Column, len(columns)),
	}
	for i, columnValue := range columns {
		column := Column{
			Value:     columnValue,
			Alignment: AlignLeft,
		}
		column.Size = column.GetSize()
		if column.Size > t.Header.Columns[i].Size {
			t.Header.Columns[i].Size = column.Size
		}
		row.Columns[i] = &column
	}

	t.Body = append(t.Body, &row)

	return nil
}

// Markdown Generates the Code Block Markdown
func (t *Table) Markdown() string {
	result := t.printHeader()
	result += t.printBody()
	return result
}

func (t *Table) printHeader() string {
	result := ""
	for i, headerColumn := range t.Header.Columns {
		if i == 0 {
			result += t.startColumn()
		}
		result += " " + headerColumn.Value
		result += insertSpaces(headerColumn.Size - len(headerColumn.Value) - 1)
		result += t.endColumn()
	}
	result += "\n"
	for i, headerColumn := range t.Header.Columns {
		if i == 0 {
			result += t.startColumn()
		}
		switch headerColumn.Alignment {
		case AlignLeft:
			result += insertChar(headerColumn.Size, "-")
		case AlignCenter:
			result += ":"
			result += insertChar(headerColumn.Size-2, "-")
			result += ":"
		case AlignRight:
			result += insertChar(headerColumn.Size-1, "-")
			result += ":"
		}
		result += t.endColumn()
	}
	result += "\n"
	return result
}

func (t *Table) printBody() string {
	result := ""
	for _, row := range t.Body {
		for i, column := range row.Columns {
			if i == 0 {
				result += t.startColumn()
			}
			result += " " + column.Value
			result += insertSpaces(t.Header.Columns[i].Size - len(column.Value) - 1)
			result += t.endColumn()
		}
		result += "\n"
	}
	result += "\n"
	return result
}

func (t *Table) startColumn() string {
	return "|"
}

func (t *Table) endColumn() string {
	return "|"
}
