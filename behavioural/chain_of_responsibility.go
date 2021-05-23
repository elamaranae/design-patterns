/*
Object Behavioral: Chain of Responsibility

“Avoid coupling the sender of a request to its receiver by giving more than one
object a chance to handle the request. Chain the receiving objects and pass the
request along the chain until an object handles it.”

Excerpt From: “Design Patterns: Elements of Reusable Object-Oriented Software”.

In this program, we pass the helpRequest from the most specific entity(button) to
least specific entity(application) while giving each object in the chain a chance to handle 
the request if it wasn't already handled.
*/

package main

import(
  "fmt"
)

type helpHandler interface {
  handleHelp()
}

type defaultHelpHandler struct {
  successor helpHandler
}

func (d defaultHelpHandler) handleHelp() {
  if d.successor != nil {
    d.successor.handleHelp()
  }
}

type application struct {
  defaultHelpHandler
}

type widget struct {
  defaultHelpHandler
}

type button struct {
  defaultHelpHandler
}
func (a application) handleHelp() {
  fmt.Println("the most general help")
}

func (w widget) handleHelp() {
  fmt.Println("the widget level help")
}

/*
even if button doesn't implement handleHelp doesn't exist we still
would print the more general widget's help message
*/
func (a button) handleHelp() {
  fmt.Println("the button specific help")
}

func main() {
  app := application{defaultHelpHandler{}}
  wig := widget{defaultHelpHandler{app}}
  but := button{defaultHelpHandler{wig}}
  but.handleHelp()
}
