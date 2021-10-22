package xorcipher

//
// Cipher / Decipher message using a key
//
func cipher(msg string, key string) string {
	var ciphered string 
	var keylen int = len(key)
	for i := 0; i < len(msg); i++ {
		var keyIndex int = i % keylen // Calculate the key index e.g. (i=10, keylen=4, keyIndex 2), (i=11, keylen=4, keyIndex=3)
		ciphered += string(msg[i] ^ key[keyIndex])
	}
	return ciphered
}