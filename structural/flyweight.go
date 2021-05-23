/*
Object Structural: Flyweight

“Use sharing to support large numbers of fine-grained objects efficiently.”

Excerpt From: “Design Patterns: Elements of Reusable Object-Oriented Software”.

In this program, we create an object to represent each character and reuse the
character object each time the character is repeated by passing the position
as argument to its operation.
*/

package main

import(
  "fmt"
)

type glyph interface{
  draw(glyphContext)
}

// extrinsic changing state
type glyphContext struct{
  row int
  col int
}

// intrinsic reusable state
type character struct{
  code byte
}


type glyphFactory struct {
  alphabets [26]*character
}

func (g *glyphFactory) createCharacter(code byte) *character {
  index := code - 97
  if g.alphabets[index] == nil {
    g.alphabets[index] = &character{code}
  }
  return g.alphabets[index]
}

func (c *character) draw(g glyphContext) {
  fmt.Printf("object_id = %p -- renders '%v' at (%v,%v)\n", c, string(c.code), g.row, g.col)
}

func main() {
  factory := glyphFactory{}
  a1 := factory.createCharacter('a')
  a2 := factory.createCharacter('a')
  // a1.code = 5
  a3 := factory.createCharacter('a')
  c1 := factory.createCharacter('c')
  fmt.Printf("%p\n", a3)
  fmt.Printf("%p\n", a1)
  fmt.Printf("%p\n", c1)
  fmt.Println(a1==a3)
  a1.draw(glyphContext{5,6})
  a2.draw(glyphContext{1,2})
  a3.draw(glyphContext{4,1})
  c1.draw(glyphContext{9,1})
  fmt.Println(factory.alphabets)
}
