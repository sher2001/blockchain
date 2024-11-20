package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func Test_generate_private_key(t *testing.T) {
// 	privKey := GeneratePrivateKey()
// 	pubKey := privKey.PublicKey()
// 	// addr := pubKey.Addr()
// 	msg := []byte("Hello, world!")
// 	signature, err := privKey.Sign(msg)

// 	assert.Nil(t, err)
// 	assert.True(t, signature.Verify(pubKey, msg))
// }

func Test_keypair_sign_verify_sucess(t *testing.T) {
	privateKey := GeneratePrivateKey()
	publicKey := privateKey.PublicKey()
	msg := "Hello World!"
	signature, err := privateKey.Sign([]byte(msg))

	assert.Nil(t, err)
	assert.True(t, signature.Verify(publicKey, []byte(msg)))
}

func Test_keypair_sign_verify_fail(t *testing.T) {
	privateKey_1 := GeneratePrivateKey()
	publicKey_1 := privateKey_1.PublicKey()
	msg := "Hello World!"
	signature, err := privateKey_1.Sign([]byte(msg))

	assert.Nil(t, err)

	privateKey_2 := GeneratePrivateKey()
	publicKey_2 := privateKey_2.PublicKey()

	// With different public key
	assert.False(t, signature.Verify(publicKey_2, []byte(msg)))
	// with different hash data
	assert.False(t, signature.Verify(publicKey_1, []byte("Hello Bro")))
}
