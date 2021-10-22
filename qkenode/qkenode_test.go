package qkenode

import (
	"log"
	"testing"
)

func TestQuntumKeyExchangeNode(t *testing.T) {
	var server *QKENode = NewQKENode(nil, "server1", 20)
	var client *QKENode = NewQKENode(server, "client1", 20)

	status, msg := client.Send("init", "", nil)
	log.Printf("status=%v, message=%v", status, msg)

	status, msg = client.Send("message", "test", nil)
	log.Printf("status=%v, message=%v", status, msg)

	status, msg = server.Send("message", "test", nil)
	log.Printf("status=%v, message=%v", status, msg)
}

func TestWithManInMiddleAttack(t *testing.T) {
	var tests [3]int = [3]int{16, 256, 1024}
	
	for _, v := range tests {
		log.Printf("\n----\nTesting with length %d qubit stream\n----\n", v)

		var server *QKENode = NewQKENode(nil, "server1", v)
		var middleMan *QKENode = NewQKENode(server, "middle1", v)
		var client *QKENode = NewQKENode(server, "client1", v)

		// Client inits secret with server
		status, msg := client.Send("init", "", nil)
		log.Printf("status=%v, message=%v", status, msg)

		client.SetUp(middleMan)
		// Client sends message to man in the middle attacker (simulated attack)
		status, msg = client.Send("message", "test", server)
		log.Printf("status=%v, message=%v", status, msg)
	}
}