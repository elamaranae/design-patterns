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
  directImage := NewImage("haha.jpg")
  // directImage.draw()
  proxyImage := imageProxy{filename: "haha.jpg"}
  fmt.Println(proxyImage)
  fmt.Println(directImage)
  // proxyImage.draw()
}
