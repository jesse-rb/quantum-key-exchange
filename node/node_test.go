package node

import "testing"

func TestNodeBasicCommunication(t *testing.T) {
	var server *Node = NewNode(nil, "server")
	var client *Node = NewNode(server, "client")

	const numTests int = 3;
	var result_tests [numTests]string
	var expected_tests [numTests]string = [numTests]string { "Hello there!", "", "General Kenobi!" }

	result_tests[0] = client.Send("Hello there!")
	result_tests[1] = server.Send("lol")
	server.SetUp(client)
	result_tests[2] = server.Send("General Kenobi!")

	for i := 0; i < numTests; i++ {
		if result_tests[i] != expected_tests[i] { t.Fatalf("Node.Send(msg) -> expected: %v, got: %v", expected_tests[i], result_tests[i]) }
	}
}