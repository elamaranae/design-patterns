/*
use_case: 1. create browser components as objects for an application
          2. implementations vary for each component in each browser
          3. object should be created in one step and returned

example   1. create browser specific extensions and logger where implementations
             are different for each browser but interface is the same

benifits  1. need not have separate concreate factory classes for each browser type
          2. a prototype is passed as parameter to the factory to determine the type
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

/* factory interface */
type browserFactory interface {
  createExtension() extension
  createLogger() logger
}

/* component interface */
type extension interface {
  doExtensionWork() string
  clone() extension
}

type logger interface {
  log() string
  clone() logger
}

/* concrete implementation of browserFactory interface */
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
