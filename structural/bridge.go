/*
use_case: 1. we want to decouple abstraction from implementation
          2. reduce proliferation of classes from nxn to n+n

example   1. create browser specific extensions and logger where implementations
             are different for each browser but interface is the same

benifits  1. clients need not re-compile when implementation detail is changed
          2. can change factories at runtime and create different family of components.
*/

package main

import(
  "fmt"
)

/* component interface */

type implementor interface {
  log() string
  doSomething() string
}

type chromiumImplementor struct {}
type firefoxImplementor struct {}

func (e chromiumImplementor) doSomething() string {
  return "I am an extension for chromium!"
}

func (e firefoxImplementor) doSomething() string {
  return "I am an extension for firefox!"
}

func (e chromiumImplementor) log() string {
  return "I log for chromium!"
}

func (e firefoxImplementor) log() string {
  return "I log for firefox!"
}

type component struct {
  imp implementor 
}

type extension struct {
  component
}

func (e extension) doExtensionWork() string {
  return e.imp.doSomething()
}

func (e logger) log() string {
  return e.imp.log()
}

type logger struct {
  component
}

func main() {
  var imp implementor = firefoxImplementor{}
  var extension extension = extension{
   component{imp},
  }
  var logger logger = logger{
   component{imp},
  }
  fmt.Println(extension.doExtensionWork())
  fmt.Println(logger.log())
}
