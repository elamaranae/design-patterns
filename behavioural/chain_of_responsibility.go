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

func (a button) handleHelp() {
  fmt.Println("the button specific help")
}

func main() {
  app := application{defaultHelpHandler{}}
  wig := widget{defaultHelpHandler{app}}
  but := button{defaultHelpHandler{wig}}
  but.handleHelp()
}
