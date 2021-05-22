/*
use_case: 1. create browser components as objects for an application
          2. implementations vary for each component in each browser
          3. object should be created in one step and returned
          4. ensure that all components created by a factory is of 
             same family

example   1. create browser specific extensions and logger where implementations
             are different for each browser but interface is the same

benifits  1. client only needs to deal with interface of factories and components,
             they need not be tied to concrete browser specific implementations
          2. can change factories at runtime and create different family of components.
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
