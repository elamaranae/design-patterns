/*
Object Creational: Factory Method

“Define an interface for creating an object, but let subclasses decide which
class to instantiate. Factory Method lets a class defer instantiation to
subclasses.”

Excerpt From: “Design Patterns: Elements of Reusable Object-Oriented Software”.

In this program, we define an interface to create and return browser components and
let individual browsers implement it accordingly.
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
