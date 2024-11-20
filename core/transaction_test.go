package core

import (
	"testing"

	"github.com/sher2001/blockchain/crypto"
	"github.com/stretchr/testify/assert"
)

func TestSignTransaction(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		data: []byte("foo"),
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.NotNil(t, tx.Signature)
}

func TestVeriFySignature(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		data: []byte("foo"),
	}

	tx.Sign(privKey)
	assert.Nil(t, tx.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	tx.PublicKey = otherPrivKey.PublicKey()
	assert.NotNil(t, tx.Verify())
}
