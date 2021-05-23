/*
Object Behavioral: Visitor

“Represent an operation to be performed on the elements of an object structure.
Visitor lets you define a new operation without changing the classes of the
elements on which it operates.”

Excerpt From: “Design Patterns: Elements of Reusable Object-Oriented Software”.

In this program, we try to represent a Nodes in an AST and implement parse and
compile operations without having to pollute the nodes themselves. It has an
added advantage that all related operations comes under single type.
*/

package main

import(
  "fmt"
)

type visitor interface{
  visitAssignmentNode(assignmentNode) 
  visitEqualityNode(equalityNode) 
}

type node interface{
  accept(visitor)
}

type assignmentNode struct{}
type equalityNode struct{}

type parser struct{}
type compiler struct{}

func (a assignmentNode) accept(v visitor) {
  v.visitAssignmentNode(a)
}

func (e equalityNode) accept(v visitor) {
  v.visitEqualityNode(e)
}

func (p parser) visitAssignmentNode(a assignmentNode) {
  fmt.Println("parsing assignment node") 
}

func (p parser) visitEqualityNode(a equalityNode) {
  fmt.Println("parsing equality node") 
}

func (p compiler) visitAssignmentNode(a assignmentNode) {
  fmt.Println("compiling assignment node") 
}

func (p compiler) visitEqualityNode(a equalityNode) {
  fmt.Println("compiling equality node") 
}

func (p parser) parse(n node) {
  n.accept(p)
}

func (c compiler) compile(n node) {
  n.accept(c)
}

func main() {
  p := parser{}
  c := compiler{}
  var ass node = assignmentNode{}
  var eq node = equalityNode{}

  p.parse(ass)
  p.parse(eq)
  c.compile(ass)
  c.compile(eq)
}
