/*
Object Behavioral: Interpreter

“Given a language, define a represention for its grammar along with an
interpreter that uses the representation to interpret sentences in the
language.”

Excerpt From: “Design Patterns: Elements of Reusable Object-Oriented Software"

In this program, we interpret a expressions with the following grammer rules.
We produce an AST that represents the grammer and user the pattern to perform
operations on them.

boolExp -> andExp | orExp | literal
andExp -> boolExp '&&' boolExp
orExp -> boolExp '||' boolExp
literal -> 'true' | 'false'
*/

package main

import(
  "fmt"
)

type boolExp interface {
  evaluate() bool
}

type andExp struct {
  op1 boolExp
  op2 boolExp
}

type orExp struct {
  op1 boolExp
  op2 boolExp
}

type literalExp struct {
  value bool
}

func (exp orExp) evaluate() bool {
  return exp.op1.evaluate() || exp.op2.evaluate()
}

func (exp andExp) evaluate() bool {
  return exp.op1.evaluate() && exp.op2.evaluate()
}

func (exp literalExp) evaluate() bool {
  return exp.value
}

func main() {
  var a boolExp = literalExp{true}
  var b boolExp = literalExp{false}
  var or boolExp = orExp{a,b}
  var and boolExp = andExp{a,b}
  var comp1 boolExp = orExp{and,b}
  var comp2 boolExp = andExp{or,a}
  fmt.Println(or.evaluate(), and.evaluate())
  fmt.Println(comp1.evaluate(), comp2.evaluate())
}
