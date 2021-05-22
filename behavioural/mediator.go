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
