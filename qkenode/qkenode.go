package qkenode

import (
	"log"
	"math/rand"
	"os"
	"quantum-key-exchange/qubit"
	"quantum-key-exchange/xorcipher"
)

//
// A QKE key type
//
type QKEKey []*qubit.Qubit

//
// Create a new QKE key of specified length
//
func newQKEKey(len int) QKEKey {
	key := make(QKEKey, 0)
	for i := 0; i < len; i ++ {
		key = append(key, qubit.NewQubit())
	}
	return key
}

//
// Copy an existing QKE key
//
func copyQKEKey(o QKEKey) QKEKey {
	c := make(QKEKey, len(o))
	copy(c, o)
	return c
}

//
// Measure QKE Key
//
func (k QKEKey) measure() {
	for i := 0; i < len(k); i++ {
		k[i].Measure(int8(rand.Intn(2)))
	}
}

//
// mask QKE key's value
//
func (k QKEKey) mask() {
	for i := 0; i < len(k); i++ {
		k[i].SetValue(0)
	}
}

//
// Generate a secret key
//
func (k QKEKey) secret(k2 QKEKey) string {
	var secret string
	for i := 0; i < len(k); i++ {
		if (k[i].PolarizationMatch(k2[i])) {
			secret += k[i].ValueToString()
		}
	}
	return secret
}

//
// A QKE node type
//
type QKENode struct {
	name	string
	up		*QKENode
	key		QKEKey // Qubit key used to initialize a secret
	secret	string // Secret key initialized from another QKE node's recieved key
	l		*log.Logger
}

//
// Create a new QKE node
//
func NewQKENode(up *QKENode, name string, keylen int) *QKENode {
	var key QKEKey = newQKEKey(keylen)
	logger := log.New(os.Stderr, "-> [QKENode]["+ name +"] -> ", log.Lmicroseconds)

	var qkeNode *QKENode = &QKENode {
		name:	name,
		up:		up,
		key: 	key,
		secret: "",
		l:		logger,
	}

	return qkeNode
}

//
// Set QKE node upstream
//
func (qken *QKENode) SetUp(up *QKENode) {
	qken.up = up
}

//
// Listen for messages
//
func (qken *QKENode) Listen(flag string, key QKEKey, body string, from string, intercepted *QKENode) (string, QKEKey, string) {
	if (intercepted != nil) {
		intercepted.Listen(flag, key, body, from, nil)
	}
	if (flag == "init:ask-for-qubits") {
		return "OK", qken.key, ""
	} else if (flag == "init:give-polarizations") {
		qken.secret = qken.key.secret(key);
		return "OK", nil, ""
	} else if (flag == "message") {
		var deciphered string = xorcipher.Cipher(body, qken.secret)

		qken.l.Printf("\n got message:\n  (encrypted): %v\n  (un-encrypted): %v\n from: %v\n", body, deciphered, from)
		return "OK", nil, xorcipher.Cipher("Recieved your message!", qken.secret)
	}
	return "BAD REQUEST", nil, ""
}

//
// Send messages
//
func (qken *QKENode) Send(flag string, msg string, intercepted *QKENode) (string, string) {
	if (flag != "init" && qken.secret == "") {
		return "BAD REQUEST", "No secret has been established, please send message with flag: \"init\" first!"
	}
	if (qken.up == nil) {
		return "BAD REQUEST", "No upstream to send to!"
	}
	if flag == "init" {
		status, returnKey, _ := qken.up.Listen(flag+":ask-for-qubits", nil, "", qken.name, intercepted)
		if (status == "OK") {
			qken.key = copyQKEKey(returnKey)
			qken.key.measure()
			// Create a new masked version to share polarizations
			var masked QKEKey
			masked = copyQKEKey(qken.key)
			masked.mask()
			status, _, _ := qken.up.Listen(flag+":give-polarizations", masked, "", qken.name, intercepted)
			if (status == "OK") {
				// Get secret
				qken.secret = qken.key.secret(returnKey)
				return "OK", "Successful init!"
			}
		}
	} else if flag == "message" {
		status, _, body := qken.up.Listen(flag, nil, xorcipher.Cipher(msg, qken.secret), qken.name, intercepted)
		if status == "OK" {
			var deciphered string = xorcipher.Cipher(body, qken.secret)
			qken.l.Printf("\n got reply:\n  (encrypted): %v\n  (un-encrypted): %v\n from: %v\n", body, deciphered, qken.up.name)
			return "OK", "Successful message"
		}
	}
	return "BAD REQUEST", "Something went wrong!"
}