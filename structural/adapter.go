/*
Object Structural: Adapter

“Convert the interface of a class into another interface clients expect.
Adapter lets classes work together that couldn’t otherwise because of
incompatible interfaces.”

Excerpt From: “Design Patterns: Elements of Reusable Object-Oriented Software”.

In this program, we make the our image node conform(adapt) to our node interface
by enclosing the standar library image operations inside a wrapper.
*/

package main

import(
  "fmt"
  "image"
  "os"
  _ "image/jpeg"
  _ "image/png"
)

func getImageFromFilePath(filePath string) (image.Image, error) {
    f, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    image, _, err := image.Decode(f)
    return image, err
}

type positioner interface {
  row() int
  column() int
}

type position struct {
  x int
  y int
}

func (p position) row() int {
  return p.x
}

func (p position) column() int {
  return p.y
}

type node interface {
  positioner
  size() (int, int)
}

type textNode struct {
  position
  content string
}

type imageNode struct {
  position
  img image.Image
  format string
}

func (t textNode) size() (int, int) {
  return 1, len(t.content)
}

/* here we adapt the builtin image interface to our application 
   specific node interface */
func (t imageNode) size() (int, int) {
  bounds := t.img.Bounds().Max
  return bounds.X, bounds.Y
}

func main() {
  var tn node = textNode{
    position{
      x: 0,
      y: 0,
    },
    "hey there",
  }
  image, err := getImageFromFilePath("image.jpg")
  if err != nil {
    fmt.Println("no image found with given path")
    return
  }
  var in node = imageNode{
    position{
      x: 10,
      y: 0,
    },
    image,
    "png",
  }
  fmt.Println(tn.size())
  fmt.Println(in.size())
}
