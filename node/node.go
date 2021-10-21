package node

import (
	"log"
	"os"
)

//
// A node that can send to an upstream, and recieve from a downstream.
// e.g. a client-server relationship
//
type Node struct {
	up		*Node
	down	*Node
	l		*log.Logger
	name	string
}

//
// Create a new Node
//
func NewNode(up *Node, down *Node, name string) *Node {
	logger := log.New(os.Stderr, "-> ["+ name +"] -> ", log.Lmicroseconds)
	// logger.SetPrefix(time.Now().Format("15:04:05") + " " + name + " ->\n")

	node := &Node {
		up:		up,
		l: 		logger,
		name:	name }

	return node;
}

//
// Recieve message from downstream node
//
func (n *Node) Recv(msg string, from string) string {
	n.l.Printf("\n recieved:\t%s\n from:\t\t%s", msg, from)
	return msg
}

//
// Send message to upstream node
//
func (n *Node) Send(msg string) string {
	if (n.up == nil) {
		n.l.Println("\n\tNo upstream to send to")
		return "";
	}
	n.l.Printf("\n sending:\t%s\n to:\t\t%s\n", msg, n.up.name)
	return n.up.Recv(msg, n.name)
}