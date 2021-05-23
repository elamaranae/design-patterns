/*
Object Behavioral: State

“Allow an object to alter its behavior when its internal state changes. The
object will appear to change its class.”

Excerpt From: “Design Patterns: Elements of Reusable Object-Oriented Software”.

In this program, we create tcpConnection object whose behavior changes depending
of the tcpState is it current in.
*/

package main

import(
  "fmt"
)

type tcpState interface{
  handshake(*tcpConnection)
  close(*tcpConnection)
}

type tcpConnection struct{
  state tcpState
}

type inactiveState struct{}
type activeState struct{}

func (c *tcpConnection) handshake() {
  c.state.handshake(c)
}

func (c *tcpConnection) close() {
  c.state.close(c)
}

func (c *tcpConnection) changeState(s tcpState) {
  c.state = s
}

func (s activeState) handshake(c *tcpConnection) {
  fmt.Println("Can't handshake, already in active state.")
}

func (s inactiveState) handshake(c *tcpConnection) {
  fmt.Println("doing handshake")
  c.changeState(activeState{})
}

func (s activeState) close(c *tcpConnection) {
  fmt.Println("closing connection")
  c.changeState(inactiveState{})
}

func (s inactiveState) close(c *tcpConnection) {
  fmt.Println("Can't close, already in inactive state.")
}

func main() {
  con := tcpConnection{state: inactiveState{}}
  con.close()
  con.handshake()
  con.handshake()
  con.close()
}
