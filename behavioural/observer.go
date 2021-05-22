package main

import(
  "fmt"
)

type subject interface {
  attach(observer)
  detach(observer)
  notify()
}

type observer interface {
  update(int)
}

type simpleSubject struct {
  state int
  list []observer
}

type simpleObserver struct {
  sub subject
  state int
}

func (s *simpleSubject) attach(o observer) {
  s.list = append(s.list, o)
}

func (s *simpleSubject) detach(o observer) {
  for i, v := range s.list {
    if v == o {
      s.list = append(s.list[:i], s.list[i+1:]...)
    }
  }
}

func (s *simpleSubject) notify() {
  for _, v := range s.list {
    v.update(s.state)
  }
}

func (so *simpleObserver) update(s int) {
  so.state = s
}

func main() {
  sub := simpleSubject{state: 0}
  so1 := simpleObserver{state: 0, sub: &sub}
  so2 := simpleObserver{state: 0, sub: &sub}
  sub.attach(&so1)
  sub.attach(&so2)
  fmt.Println(so1)
  fmt.Println(so2)
  sub.state = 3
  sub.notify()
  fmt.Println(so1)
  fmt.Println(so2)
  sub.state = 4
  sub.detach(&so1)
  sub.notify()
  fmt.Println(so1)
  fmt.Println(so2)
}
