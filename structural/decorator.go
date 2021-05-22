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
