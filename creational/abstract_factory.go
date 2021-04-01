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

/* factory objects for each browser */
type chromiumFactory struct {}
type firefoxFactory struct {}

/* component objects */
type chromiumExtension struct {}
type firefoxExtension struct {}

type chromiumLogger struct {}
type firefoxLogger struct {}

/* factory interface */
type browserFactory interface {
  createextension() extension
  createlogger() logger
}

/* component interface */
type extension interface {
  doextensionWork() string
}

type logger interface {
  log() string
}

/* concrete implementation of browserFactory interface */
func (e chromiumFactory) createextension() extension {
  return chromiumExtension{}
}

func (e firefoxFactory) createextension() extension {
  return firefoxExtension{}
}

func (e chromiumFactory) createlogger() logger {
  return chromiumLogger{}
}

func (e firefoxFactory) createlogger() logger {
  return firefoxLogger{}
}

/* concrete implementation of extension interface */
func (e chromiumExtension) doextensionWork() string {
  return "I am an extension for chromium!"
}

func (e firefoxExtension) doextensionWork() string {
  return "I am an extension for firefox!"
}

/* concrete implementation of extensionlogger inteface */
func (l chromiumLogger) log() string {
  return "I log in chrome console."
}

func (l firefoxLogger) log() string {
  return "I log in firefox console."
}

func main() {
  /* we replace this factory with chromiumFactory during runtime
     without having to change anything else */
  var factory browserFactory = firefoxFactory{}
  var extension extension = factory.createextension()
  var logger logger = factory.createlogger()
  fmt.Println(extension.doextensionWork())
  fmt.Println(logger.log())
}
