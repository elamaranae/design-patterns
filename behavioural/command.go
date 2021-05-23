/*
Object Behavioral: Command

“Encapsulate a request as an object, thereby letting you parameterize clients
with different requests, queue or log requests, and support undoable
operations.”

Excerpt From: “Design Patterns: Elements of Reusable Object-Oriented Software".

In this program, we make paste operation into a class that responds to execute
and also create a macro command(composite command) that calls sequence of commands.
*/

package main

import(
  "fmt"
)

type command interface {
  execute()
}

type document struct {
  text string
}

func (d document) paste() {
  fmt.Println(d.text)
  fmt.Println("pasting command is completed")
}

type pasteCommand struct {
  doc document
}

type compositeCommand struct {
  commands []command
}

func (c compositeCommand) execute() {
  for _, v := range c.commands {
    v.execute()
  }
}

func (c *compositeCommand) add(cmd command) {
  c.commands = append(c.commands, cmd)
}

func (p pasteCommand) execute() {
  p.doc.paste()
}

func main() {
  d := document{"a doc"}
  var c1 command = pasteCommand{doc: d}
  var c2 command = pasteCommand{doc: d}
  // c1.execute()
  // c2.execute()
  co := compositeCommand{}
  co.add(c1)
  co.add(c2)
  co.execute()
}
