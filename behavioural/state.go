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
  fmt.Println("doning handshake")
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
