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
