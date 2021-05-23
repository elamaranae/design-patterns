/*
Object Structural: Proxy

“Provide a surrogate or placeholder for another object to control access to it.”

Excerpt From: “Design Patterns: Elements of Reusable Object-Oriented Software”.

In this program, we create a proxy object to represent an image whenever an image
needs to be created. The proxy only loads the actual image only when its needed.
*/

package main

import(
  "fmt"
)

type glyph interface{
  draw()
}

type image struct{}

func (i image) draw() {
  fmt.Println("a costly operation in progress - drawing image")
}

type imageProxy struct{
  filename string
  img *image
}

func (i imageProxy) getImage() *image {
  if i.img == nil {
    i.img = NewImage(i.filename)
  }
  return i.img
}

func NewImage(filename string) *image {
  fmt.Println("a costly operation in progress - loading image from filesystem")
  return &image{}
}

func (i imageProxy) draw() {
  i.getImage().draw()
}

func main() {
  // here image is not yet loaded
  proxyImage := imageProxy{filename: "haha.jpg"}
  // it's lazily loaded only when we call draw on it
  proxyImage.draw()
}
