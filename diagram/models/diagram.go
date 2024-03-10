package models

import (
	"fmt"
)

type Diagram struct {
	Boxes []DiagramBox
}

// TODO All algorithm
func (d *Diagram) Print() string {
	diagram := ""
	for _, box := range d.Boxes {
		diagram += fmt.Sprintf("%s\n", box.Print())
	}

	return diagram
}
