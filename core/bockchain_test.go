package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newBlockchainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(0))

	assert.Nil(t, err)
	return bc
}

func TestNewBlockchain(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	assert.NotNil(t, bc.validator)
	assert.Equal(t, uint32(0), bc.Height())
}

func TestAddBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	length := 1000

	for i := 0; i < length; i++ {
		newBlock := ranodmBlockWithSignature(t, (uint32(i + 1)))
		assert.Nil(t, bc.AddBlock(newBlock))
	}
	assert.Equal(t, uint32(length), bc.Height())
	assert.Equal(t, length+1, len(bc.headers))
	assert.NotNil(t, bc.AddBlock(randomBlock(89)))
}

func TestHasBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.True(t, bc.HasBlock(uint32(0)))
}
