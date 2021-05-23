/*
Object Creational: Prototype

“Specify the kinds of objects to create using a prototypical instance, and
create new objects by copying this prototype.”

Excerpt From: “Design Patterns: Elements of Reusable Object-Oriented Software”.

In this program, we create a prototype factory by intializing it with prototypical
instances. This factory creates objects by cloning these prototypes.
*/

package main

import(
  "fmt"
)

/* factory objects for each browser */
type browserPrototypeFactory struct {
  extension extension
  logger logger
}

/* component objects */
type chromiumExtension struct {}
type firefoxExtension struct {}

type chromiumLogger struct {}
type firefoxLogger struct {}

/* component interface */
type extension interface {
  doExtensionWork() string
  clone() extension
}

type logger interface {
  log() string
  clone() logger
}

func (factory browserPrototypeFactory) createExtension() extension {
  return factory.extension.clone()
}

func (factory browserPrototypeFactory) createLogger() logger {
  return factory.logger.clone()
}

/* concrete implementation of Extension interface */
func (ext chromiumExtension) doExtensionWork() string {
  return "I am an extension for chromium!"
}

func (ext firefoxExtension) doExtensionWork() string {
  return "I am an extension for firefox!"
}

/* concrete implementation of ExtensionLogger inteface */
func (logger chromiumLogger) log() string {
  return "I log in chrome console."
}

func (logger firefoxLogger) log() string {
  return "I log in firefox console."
}

/* each product should implement clone */
func (logger chromiumLogger) clone() logger {
  return logger
}

func (logger firefoxLogger) clone() logger {
  return logger
}

func (extension chromiumExtension) clone() extension {
  return extension
}

func (extension firefoxExtension) clone() extension {
  return extension
}

func main() {
  /* we can pass any extension and logger as prototype to the factory
     and expect the clone of the same during creation */
  var factory browserPrototypeFactory = browserPrototypeFactory{
    firefoxExtension{},
    firefoxLogger{},
  }
  var extension extension = factory.createExtension()
  var logger logger = factory.createLogger()
  fmt.Println(extension.doExtensionWork())
  fmt.Println(logger.log())
}
