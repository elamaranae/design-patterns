/*
Object Behavioral: Memento

“Without violating encapsulation, capture and externalize an object’s internal
state so that the object can be restored to this state later.”

Excerpt From: “Design Patterns: Elements of Reusable Object-Oriented Software”.

In this program, we allow the client to preserve the complete state of the originator object
without allowing the client to access its internal state using a memento object.
*/

package main

import(
  "fmt"
)

type originator struct {
  a int
  b int
}

type memento struct {
  state map[string]int
}

func (m memento) getState() map[string]int {
  return m.state
}

func (m *memento) setState(o originator) {
  m.state = map[string]int{}
  m.state["a"] = o.a
  m.state["b"] = o.b
}

func (o *originator) setMemento(m memento) {
  state := m.getState()
  o.a = state["a"]
  o.b = state["b"]
}

func (o originator) createMemento(m *memento) {
  m.setState(o)
}

func main() {
  o := originator{1, 2}
  m := memento{}
  o.createMemento(&m)
  o.a = 8
  fmt.Println(o)
  o.setMemento(m)
  fmt.Println(o)
}
