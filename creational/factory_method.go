/*
note:     this pattern is similar to abstrace factory in that they both use factory
          methods that implement a interface to create and return objects, but this
          pattern delegates creation to a concreate subclass's method through inheritance,
          where abstract factory pattern delegates creation to a separate factory object
          through composition.
use_case: 1. create browser components as objects for an application
          2. implementations vary for each component in each browser
          3. object should be created completely and returned
          4. object doesn't know which type component to create explicitly

example   1. create browser specific extensions and logger where implementations
             are different for each browser but interface is the same

benifits  1. type of component created is delegated to the subclass
*/

package main

import(
  "fmt"
)

/* browser application objects */
type chromiumApplication struct {}
type firefoxApplication struct {}

/* component objects */
type chromiumExtension struct {}
type firefoxExtension struct {}

type chromiumLogger struct {}
type firefoxLogger struct {}

/* factory interface */
type BrowserApplication interface {
  createExtension() extension
  createLogger() logger
}

/* components interface */
type extension interface {
  doExtensionWork() string
}

type logger interface {
  log() string
}

/* concrete implementation of extensionFactory interface */
func (ext chromiumApplication) createExtension() extension {
  return chromiumExtension{}
}

func (ext firefoxApplication) createExtension() extension {
  return firefoxExtension{}
}

func (ext chromiumApplication) createLogger() logger {
  return chromiumLogger{}
}

func (ext firefoxApplication) createLogger() logger {
  return firefoxLogger{}
}

/* concrete implementation of extension interface */
func (ext chromiumExtension) doExtensionWork() string {
  return "I am an extension for chromium!"
}

func (ext firefoxExtension) doExtensionWork() string {
  return "I am an extension for firefox!"
}

/* concrete implementation of extensionLogger inteface */
func (logger chromiumLogger) log() string {
  return "I log in chrome console."
}

func (logger firefoxLogger) log() string {
  return "I log in firefox console."
}

func main() {
  /* we can replace firefoxApplication with chromiumApplication
     without having to change anything else */
  var browser firefoxApplication = firefoxApplication{}
  var extension extension = browser.createExtension()
  var logger logger = browser.createLogger()
  fmt.Println(extension.doExtensionWork())
  fmt.Println(logger.log())
}
