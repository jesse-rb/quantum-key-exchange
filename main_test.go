package main

import (
	"quantum-key-exchange/node"
	"quantum-key-exchange/qubit"
	"testing"
)

func TestNodeBasicCommunication(t *testing.T) {
	var server *node.Node = node.NewNode(nil, "server")
	var client *node.Node = node.NewNode(server, "client")

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

func TestQubitClass(t *testing.T) {
	var expected_testVal1 int8 = 0;
	var expected_testVal2 int8 = 1;
	
	for i := 0; i < 100; i++ {
		var qubit *qubit.Qubit = qubit.NewQubit()
		var measured int8 = qubit.Measure(0)
		if measured != expected_testVal1 && expected_testVal2 != 1 { // measured has to be 1 or 0
			t.Fatalf("Qubit.Measure(polarization) -> expected: %v or %v, got: %v", expected_testVal1, expected_testVal2, measured)
		}
	}
}