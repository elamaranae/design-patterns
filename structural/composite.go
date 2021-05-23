/*
Object Structural: Composite

“Compose objects into tree structures to represent part-whole hierarchies.
Composite lets clients treat individual objects and compositions of objects
uniformly.”

Excerpt From: “Design Patterns: Elements of Reusable Object-Oriented Software”.

In this program, we create primitive drawing objects and composite drawing objects
and let the client work on all of them without having to disinguish between them.
*/

package main

import(
  "fmt"
)

type component interface {
  add(component)
  remove(component)
  size() (int, int)
}

type primitive struct {}

func (d primitive) add(c component) {}
func (d primitive) remove(c component) {}

type paragraph struct {
  primitive
  content string
}

type image struct {
  primitive
  width int
  height int
}

type table struct {
  primitive
  columns int
  rows int
}

func (p paragraph) size() (int, int) {
  return len(p.content), 1
}

func (i image) size() (int, int) {
  return i.width, i.height
}

func (t table) size() (int, int) {
  return t.columns * 10, t.rows * 5
}

type composite struct {
  children []component
}

func (co *composite) add(c component) {
  co.children = append(co.children, c)
}

func (co *composite) remove(c component) {
  for i, v := range co.children {
    if v == c {
      co.children = append(co.children[:i], co.children[i+1:]...)
      break
    }
  }
}

func (co composite) size() (int, int) {
  width, height := 0, 0
  for _, v := range co.children {
    c_width, c_height := v.size()
    width += c_width
    height += c_height
  }
  return width, height
}

func main() {
  p := paragraph{ content: "this is content" }
  i := image{ height: 50, width: 50 }
  t := table{ rows: 10, columns: 2 }
  fmt.Println(p.size())
  fmt.Println(i.size())
  fmt.Println(t.size())
  c1 := &composite{}
  c2 := &composite{}
  c1.add(p)
  fmt.Println(c1.size())
  c1.add(i)
  fmt.Println(c1.size())
  c2.add(t)
  c1.add(c2)
  c1.remove(c2)
  c1.remove(i)
  c1.remove(p)
  fmt.Println(c2.size())
  fmt.Println(c1.size())
}
