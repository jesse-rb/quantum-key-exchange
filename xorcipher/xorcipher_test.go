package xorcipher

import "testing"

func TestXORCipherDecipher(t *testing.T) {
	var originial string = "hello"
	var key string = "world"

	var ciphered string = Cipher(originial, key)
	if (ciphered == originial) {
		t.Fatalf("cipher(msg, key) -> expected: ciphered=%v != original=%v, got: ciphered=%v == original=%v", ciphered, originial, ciphered, originial)
	}

	var deciphered string = Cipher(ciphered, key)
	if (originial != deciphered) {
		t.Fatalf("cipher(msg, key) -> expected: original=%v == deciphered=%v, got: original=%v != deciphered=%v", ciphered, originial, ciphered, originial)
	}
}