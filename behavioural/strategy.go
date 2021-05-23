/*
Object Behavioral: Strategy

“Define a family of algorithms, encapsulate each one, and make them
interchangeable. Strategy lets the algorithm vary independently from clients
that use it.”

Excerpt From: “Design Patterns: Elements of Reusable Object-Oriented Software”.

In this program, we allow the client to specify the sorting algoritm to use
by just swapping out an object.
*/

package main

import(
  "fmt"
)

type context struct{
  arr []int
  sortAlgo sorter
}

func (c context) sort() {
  c.sortAlgo.sort(c.arr);
}

type sorter interface{
  sort(arr []int) []int
}

type mergeSort struct{}

type bogoSort struct{}

func (m mergeSort) sort(arr []int) []int {
  fmt.Println("merge sorted");
  arr[0] = 5;
  return arr;
}

func (m bogoSort) sort(arr []int) []int {
  fmt.Println("bogo sorted");
  arr[0] = 9;
  return arr;
}

func main() {
  con := context{
    arr: []int{2,3,4},
    sortAlgo: mergeSort{},
  }
  con.sort();
  fmt.Println(con.arr);
}
