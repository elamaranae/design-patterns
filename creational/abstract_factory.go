/*
Object Creational: Abstract Factory

“Provide an interface for creating families of related or dependent objects
without specifying their concrete classes.”

Excerpt From: “Design Patterns: Elements of Reusable Object-Oriented Software”.

In this program, we create factories for each os and we create os specific
components form the respective factory.
*/

package main

import(
  "fmt"
)

type Button interface {
  click()
}
type Window interface {
  close()
}

type linuxButton struct {}
type linuxWindow struct {}
type macButton   struct {}
type macWindow   struct {}

func (b linuxButton) click() {
  fmt.Println("clicked a linux button")
}
func (b linuxWindow) close() {
  fmt.Println("closed a linux window")
}
func (b macButton) click() {
  fmt.Println("clicked a mac button")
}
func (b macWindow) close() {
  fmt.Println("closed a mac window")
}

type Factory interface {
  createButton() Button
  createWindow() Window
}

type linuxFactory struct {}
type macFactory struct {}

func (b linuxFactory) createButton() Button {
  return linuxButton{}
}
func (b linuxFactory) createWindow() Window {
  return linuxWindow{}
}
func (b macFactory) createButton() Button {
  return macButton{}
}
func (b macFactory) createWindow() Window {
  return macWindow{}
}

func main() {
  var f Factory = macFactory{}
  var b Button = f.createButton()
  var w Window = f.createWindow()
  b.click()
  w.close()
}
