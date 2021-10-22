//
// Not used in QKE algorithm, but a good simpler example of the communication between nodes
// without the QKE overhead code (for referecne)
//

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
	l		*log.Logger
	name	string
}

//
// Create a new Node
//
func NewNode(up *Node, name string) *Node {
	logger := log.New(os.Stderr, "-> [Node]["+ name +"] -> ", log.Lmicroseconds)

	node := &Node {
		up:		up,
		l: 		logger,
		name:	name,
	}

	return node;
}

//
// Set upstream, so that upstream can still be changed later after node initialization
//
func (n *Node) SetUp(up *Node) {
	n.up = up
}

//
// Recieve message from downstream node
//
func (n *Node) Listen(msg string, from string) string {
	n.l.Printf("\n recieved:\t%s\n from:\t\t%s", msg, from)
	return msg
}

//
// Send message to upstream node
//
func (n *Node) Send(msg string) string {
	if (n.up == nil) {
		n.l.Printf("\n sending:\t%s\n to:\t\tN/A\n", msg)
		return "";
	}
	n.l.Printf("\n sending:\t%s\n to:\t\t%s\n", msg, n.up.name)
	return n.up.Listen(msg, n.name)
}