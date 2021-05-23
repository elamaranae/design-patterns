/*
Object Creational: Singleton

“Ensure a class only has one instance, and provide a global point of access to it.”

Excerpt From: “Design Patterns: Elements of Reusable Object-Oriented Software”.

In this program, we allow at any time only one factory instance to be instantiated.
*/

package main

import(
  "fmt"
)

/* factory objects for each browser */
type browserFactory struct {}

/* component objects */
type browserLogger struct {}

/* store the instansiated factory */
var instansiated *browserFactory = nil

/* create new object only if not already created.
   both browserFactory and instatiated is private, So a new
   browserFactory can only be created using NewbrowserFactory.
   this will make sure only one factory can be created
*/
func NewBrowserFactory() browserFactory {
  if instansiated == nil {
    instansiated = &browserFactory{}
  }
  return *instansiated
}

func (_ browserFactory) createLogger() browserLogger {
  return browserLogger{}
}

func (logger browserLogger) log() string {
  return "I log in browser console."
}

func main() {
  var factory browserFactory = NewBrowserFactory()
  var logger browserLogger = factory.createLogger()
  fmt.Println(logger.log())
  /* this wont create a new instance */
  factory2 := NewBrowserFactory()
  fmt.Printf("%p\n", &factory)
  fmt.Printf("%p\n", &factory2)
}
