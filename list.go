package markdown

import (
	"fmt"
)

// ListType Enum
type ListType int

// ListType Enum definition
const (
	Ordered ListType = iota
	Unordered
	Task
)

// ListItem Entity
type ListItem struct {
	Indent int
	Type   ListType
	ID     int
	Value  string
	IsDone bool
	List   *List
}

// CreateList Creates a Sublist inside an List Item
func (li *ListItem) CreateList() *List {
	list := createList()
	list.Indent = li.Indent + 1
	list.Type = li.Type
	li.List = list
	return list
}

// List Entity
type List struct {
	Indent int
	Type   ListType
	Items  []*ListItem
}

// CreateList Creates an unordered List Element
func (d *Document) CreateList() *List {
	list := createList()

	d.Elements = append(d.Elements, list)
	return list
}

// CreateOrderedList Creates an ordered List Element
func CreateOrderedList() *List {
	list := List{
		Indent: 0,
		Type:   Ordered,
		Items:  make([]*ListItem, 0),
	}

	return &list
}

// CreateTaskList Creates an task List Element
func CreateTaskList() *List {
	list := List{
		Indent: 0,
		Type:   Task,
		Items:  make([]*ListItem, 0),
	}

	return &list
}

// Add Adds an element to the List
func (l *List) Add(value string) *List {
	lid := len(l.Items) + 1
	listItem := ListItem{
		Type:   l.Type,
		ID:     lid,
		Value:  value,
		List:   createList(),
		IsDone: false,
	}
	listItem.List.Type = l.Type
	listItem.List.Indent = l.Indent + 1
	listItem.Indent = l.Indent
	l.Items = append(l.Items, &listItem)

	return l
}

// GetItem Gets an ListItem from a List
func (l *List) GetItem(index int) *ListItem {
	if index >= 0 && len(l.Items) >= index {
		return l.Items[index]
	}
	return nil
}

// SetDone Set a Task Item as done
func (l *List) SetDone(id int) *List {
	sid := id - 1
	if sid >= 0 && len(l.Items) >= id {
		l.Items[sid].IsDone = true
	}
	return l
}

// SetUndone Set a Task Item as Undone
func (l *List) SetUndone(id int) *List {
	sid := id - 1
	if sid >= 0 && len(l.Items) >= id {
		l.Items[sid].IsDone = false
	}
	return l
}

// Markdown Generates the Code Block Markdown
func (l *List) Markdown() string {
	result := ""
	for _, item := range l.Items {
		if l.Indent > 0 {
			switch l.Type {
			case Unordered:
				result += insertSpaces(l.Indent * 2)
			case Ordered:
				result += insertSpaces(l.Indent * 3)
			}
		}
		switch l.Type {
		case Unordered:
			result += "- " + item.Value
		case Ordered:
			result += fmt.Sprint(item.ID) + ". " + item.Value
		case Task:
			if item.IsDone {
				result += "- [X] " + item.Value
			} else {
				result += "- [ ] " + item.Value
			}
		}
		result += "\n"
		if item.List != nil && len(item.List.Items) > 0 {
			if item.List.Type == Task {
				item.List.Indent = 0
			}
			result += item.List.Markdown()
		}
	}

	if l.Indent == 0 {
		result += "\n"
	}
	return result
}

func createList() *List {
	list := List{
		Indent: 0,
		Type:   Unordered,
		Items:  make([]*ListItem, 0),
	}

	return &list
}
