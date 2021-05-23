/*
Object Behavioral: Mediator

“Define an object that encapsulates how a set of objects interact. Mediator
promotes loose coupling by keeping objects from referring to each other
explicitly, and it lets you vary their interaction independently.”

Excerpt From: “Design Patterns: Elements of Reusable Object-Oriented Software”

In this program, we implement a mediator that co-ordinates interactions between
multiple ui elements and show an example interaction where clicking on the
checkbox automatically presses a button while keeping the button and checkbox
unaware of each other.
*/

package main

import(
  "fmt"
)

type widget interface {
  changed()
}

type mediator interface {
  createWidgets()
  handleChanged(widget)
}

type button struct {
  director mediator
}

type list struct {
  director mediator
}

type checkbox struct {
  director mediator
  state bool
}

type director struct {
  b button
  c checkbox
  l list
}

func (d director) handleChanged(w widget) {
  // switch t := reflect.TypeOf(w) {
  //   case 'checkbox'
  // }
  switch w.(type) {
	case button:
		fmt.Println("button pressed")
	case checkbox:
		fmt.Println("checked")
    d.b.press()
	default:
		fmt.Println("I don't k")
	}
}

func (d *director) createWidgets() {
  d.b = button{d} 
  d.c = checkbox{director: d, state: false}
  d.l = list{d}
}

func (b button) changed() {
  b.director.handleChanged(b)
}

func (c checkbox) changed() {
  c.director.handleChanged(c)
}

func (l list) changed() {
  l.director.handleChanged(l)
}

func (c checkbox) check() {
  c.changed()
}

func (b button) press() {
  b.changed()
}

func main() {
  var d director = director{}
  d.createWidgets()
  d.c.check()
}
