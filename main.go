package main

import (
	"quantum-key-exchange/node"
)

var server *node.Node
var client *node.Node

func init() {
	server = node.NewNode(nil, nil, "server")
	client = node.NewNode(server, nil, "client")
}

func main() {
	client.Send("Hello there")
}