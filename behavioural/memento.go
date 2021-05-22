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
