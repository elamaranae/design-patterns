/*
Object Structural: Decorator

“Attach additional responsibilities to an object dynamically. Decorators
provide a flexible alternative to subclassing for extending functionality.”

Excerpt From: “Design Patterns: Elements of Reusable Object-Oriented Software”.

In this program, we define text decorators that can be used to perform additional
optional operations dynamically.
*/

package main

import(
  "fmt"
  "html"
)

type abstractTextField interface {
  renderText() string
}

type textField struct {
  content string
}

func (t textField) renderText() string {
  return t.content
}

type spellDecorator struct {
  component abstractTextField
}

type htmlSafeDecorator struct {
  component abstractTextField
}

func (t spellDecorator) renderText() string {
  content := t.component.renderText()
  fmt.Println("the content is spell checked here and corrected") 
  return content
}

func (t htmlSafeDecorator) renderText() string {
  content := t.component.renderText()
  content = html.EscapeString(content)
  return content
}

func main() {
  field := htmlSafeDecorator{
    spellDecorator{
      textField{ content: "<h>this is awesome</h>" },
    },
  }
  fmt.Println(field.renderText())
}
