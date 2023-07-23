package _2_Component

import "testing"

func TestLeaf_Operation(t *testing.T) {
	root := NewComposite("root")
	root.Add(&Leaf{"Leaf A"})
	root.Add(&Leaf{"Leaf B"})

	comp := NewComposite("Composite X")
	comp.Add(&Leaf{"Leaf XA"})
	comp.Add(&Leaf{"Leaf XB"})

	root.Add(comp)

	root.Display(1)
}
