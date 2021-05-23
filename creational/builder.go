/*
Object Creational: Builder

“Separate the construction of a complex object from its representation so that
the same construction process can create different representations.”

Excerpt From: “Design Patterns: Elements of Reusable Object-Oriented Software”.

In this program, we create a builders that can build a html document and markdown
documents form rich text incrementally.
*/


package main

import(
  "fmt"
)

/* document objects */
type htmlDocument struct {
  lines []htmlLine
}

type markdownDocument struct {
  lines []markdownLine
}

/* line objects that make up documents 
   note: we don't need RichTextDocument because we directly use
   series of hardcoded RichTextLines. In real life, we would need it.
*/
type richTextLine struct {
  content string
  decoration string
}

type htmlLine struct {
  content string
  tag string
}

type markdownLine struct {
  content string
  element string
  enclosing bool
}

/* this builder interface is not actually is used by the director. 
   the director directly calls concrete builders and is only present
   here for the sake of completion, because golang doesn't allow interfaces
   to be used when calling a method of an implementor that doesn't exist
   in the interface even though its present in the implementor
*/
type builder interface {
  buildLine(line richTextLine)
}

/* default builder exists to provide default implementation
   of the methods and acts as the "interface" mentioned in GOF book*/
type defaultBuilder struct {}
func (builder defaultBuilder) buildLine(line richTextLine) {}

/* concreate builder objects composes DefaultBuilder to get the
   default implementations */
type htmlBuilder struct {
  defaultBuilder
  document htmlDocument
}

type markdownBuilder struct {
  defaultBuilder
  document markdownDocument
}

/* actual implementations of the concrete builders */
func (builder *htmlBuilder) buildLine(line richTextLine) {
  switch line.decoration {
  case "heading":
    builder.document.lines = append(builder.document.lines, htmlLine{ 
      content: line.content,
      tag: "h1",
    })
  case "bold":
    builder.document.lines = append(builder.document.lines, htmlLine{ 
      content: line.content,
      tag: "b",
    })
  case "italic":
    builder.document.lines = append(builder.document.lines, htmlLine{ 
      content: line.content,
      tag: "i",
    })
  }
}

func (builder *markdownBuilder) buildLine(line richTextLine) {
  switch line.decoration {
  case "heading":
    builder.document.lines = append(builder.document.lines, markdownLine{ 
      content: line.content,
      element: "#",
      enclosing: false,
    })
  case "bold":
    builder.document.lines = append(builder.document.lines, markdownLine{ 
      content: line.content,
      element: "**",
      enclosing: true,
    })
  case "italic":
    builder.document.lines = append(builder.document.lines, markdownLine{ 
      content: line.content,
      element: "*",
      enclosing: true,
    })
  }
}

func (builder htmlBuilder) getDocument() htmlDocument {
  return builder.document
}

func (builder markdownBuilder) getDocument() markdownDocument {
  return builder.document
}


func main() {
  htmlBuilder := htmlBuilder{}
  markdownBuilder := markdownBuilder{}

  richTextLine1 := richTextLine{
    content: "hey this is a heading",
    decoration: "heading",
  }
  richTextLine2 := richTextLine{
    content: "hey this is in italic",
    decoration: "italic",
  }

  /* building documents step by step */
  htmlBuilder.buildLine(richTextLine1)
  htmlBuilder.buildLine(richTextLine2)
  markdownBuilder.buildLine(richTextLine1)
  markdownBuilder.buildLine(richTextLine2)

  htmlDocument := htmlBuilder.getDocument()
  markdownDocument := markdownBuilder.getDocument()

  fmt.Println(htmlDocument)
  fmt.Println(markdownDocument)
}
