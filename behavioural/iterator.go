/*
Object Behavioral: Iterator

“Provide a way to access the elements of an aggregate object sequentially
without exposing its underlying representation.”

Excerpt From: “Design Patterns: Elements of Reusable Object-Oriented Software"

In this program, we create a custom iterator type for a custom list type by
conforming to a more general list and iterator interfaces.
*/

package main

import(
  "fmt"
)

type item interface {}

type list interface {
  get(int) item
  count() int
}

type intItem int

type intList []intItem

func (l intList) get(index int) item {
  return l[index]
}

func (l intList) count() int {
  return len(l)
}

type iterator interface {
  next()
  done() bool
  current() item
  first()
}

type listIterator struct {
  lis list
  curr int
}

func (i *listIterator) first() {
  i.curr = 0
}

func (i listIterator) done() bool {
  return i.curr >= i.lis.count()
}

func (i listIterator) current() item {
  return i.lis.get(i.curr)
}

func (i *listIterator) next() {
  i.curr += 1
}

func main() {
  var li list = intList([]intItem{7, 2, 3, 9})
  var iter iterator = &listIterator{lis: li}
  for iter.first(); !iter.done(); iter.next() {
    fmt.Println(iter.current())
  }
}
