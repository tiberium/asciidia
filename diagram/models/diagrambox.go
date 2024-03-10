package models

import (
	"strings"
)

type Position struct {
	X int
	Y int
}

type DiagramBox struct {
	Name         string
	Dependencies []DiagramBox
	Position     Position
}

func (d DiagramBox) Print() []string {
	nameLength := len(d.Name)

	return []string{strings.Repeat("-", nameLength+2), "|" + d.Name + "|", strings.Repeat("-", nameLength+2)}
}
