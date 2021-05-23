/*
Object Structural: Bridge

“Decouple an abstraction from its implementation so that the two can vary
independently.”

Excerpt From: “Design Patterns: Elements of Reusable Object-Oriented Software (Joanne Romanovich's Library)”.

In this program, we implement logger and extensions for different browsers. But, instead of creating an object
for each combination (NxN), we specify the behaviour as operations on browser implementors. This reduces the
total objects to (N+N).
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

type logger struct {
  component
}

func (e extension) doExtensionWork() string {
  return e.imp.doSomething()
}

func (e logger) log() string {
  return e.imp.log()
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
