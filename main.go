package main

import (
	//"flag"
	"fmt"

	"github.com/tiberium/asciidia/diagram/models"
)

func main() {
	a := models.DiagramBox{Name: "AAA", Position: models.Position{X: 1, Y: 1}}
	b := models.DiagramBox{Name: "BBB", Dependencies: []models.DiagramBox{a}, Position: models.Position{X: 11, Y: 1}}
	c := models.DiagramBox{Name: "CCC", Dependencies: []models.DiagramBox{b}, Position: models.Position{X: 21, Y: 1}}
	d := models.DiagramBox{Name: "DDD", Dependencies: []models.DiagramBox{b}, Position: models.Position{X: 11, Y: 7}}

	/*
		a -+ b -+ c
			 +
			 |
			 d
	*/

	diagram := models.Diagram{Boxes: []models.DiagramBox{a, b, c, d}}

	fmt.Print(diagram.Print())
}
