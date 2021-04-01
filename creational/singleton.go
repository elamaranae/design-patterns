/*
use_case: 1. create browser components as objects for an application
          2. ensure only one instance of the factory can be created

example   1. create browser specific extensions and logger where implementations
             are different for each browser but interface is the same
*/

package main

import(
  "fmt"
)

/* factory objects for each browser */
type chromiumFactory struct {}

/* component objects */
type chromiumExtension struct {}
type chromiumLogger struct {}

/* factory interface */
type browserFactory interface {
  createExtension() extension
  createLogger() logger
}

/* component interface */
type extension interface {
  doExtensionWork() string
}

type logger interface {
  log() string
}

/* store the instansiated factory */
var instansiated browserFactory = nil

/* create new object only if not already created 
   both browserFactory and instatiated is private, so a new
   browserFactory can only be created using NewChromiumFactory.
   this will make sure only one factory can be created
*/
func NewChromiumFactory() browserFactory {
  if instansiated == nil {
    instansiated = chromiumFactory{}
  }
  return instansiated
}

/* concrete implementation of browserFactory interface */
func (_ chromiumFactory) createExtension() extension {
  return chromiumExtension{}
}

func (_ chromiumFactory) createLogger() logger {
  return chromiumLogger{}
}

/* concrete implementation of Extension interface */
func (ext chromiumExtension) doExtensionWork() string {
  return "I am an extension for chromium!"
}

/* concrete implementation of ExtensionLogger inteface */
func (logger chromiumLogger) log() string {
  return "I log in chrome console."
}

func main() {
  var factory browserFactory = NewChromiumFactory()
  var extension extension = factory.createExtension()
  /* this wont create a new instance */
  factory = NewChromiumFactory()
  var logger logger = factory.createLogger()
  fmt.Println(extension.doExtensionWork())
  fmt.Println(logger.log())
}
