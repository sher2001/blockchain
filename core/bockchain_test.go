package core

import (
	"testing"

	"github.com/sher2001/blockchain/types"
	"github.com/stretchr/testify/assert"
)

func TestNewBlockchain(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	assert.NotNil(t, bc.validator)
	assert.Equal(t, uint32(0), bc.Height())
}

func TestAddBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	length := 1000
	for i := 0; i < length; i++ {
		newBlock := ranodmBlockWithSignature(t, uint32(i+1), getPrevBlockHash(t, bc, uint32(i+1)))
		assert.Nil(t, bc.AddBlock(newBlock))
	}
	assert.Equal(t, uint32(length), bc.Height())
	assert.Equal(t, length+1, len(bc.headers))
	// add a block with already existed height
	assert.NotNil(t, bc.AddBlock(randomBlock(89, types.Hash{})))
}

func TestGetHeader(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	length := 1000
	for i := 0; i < length; i++ {
		newBlock := ranodmBlockWithSignature(t, uint32(i+1), getPrevBlockHash(t, bc, uint32(i+1)))
		assert.Nil(t, bc.AddBlock(newBlock))
		header, err := bc.GetHeader(uint32(i + 1))
		assert.Nil(t, err)
		assert.Equal(t, newBlock.Header, header)
	}
}

func TestHasBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.True(t, bc.HasBlock(uint32(0)))
	assert.False(t, bc.HasBlock(uint32(1)))
	assert.False(t, bc.HasBlock(uint32(100)))
}

func TestAddBlockWithTooMuchHeight(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.Nil(t, bc.AddBlock(ranodmBlockWithSignature(t, uint32(1), getPrevBlockHash(t, bc, uint32(1)))))
	assert.NotNil(t, bc.AddBlock(ranodmBlockWithSignature(t, uint32(3), types.Hash{})))
}

func newBlockchainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(0, types.Hash{}))

	assert.Nil(t, err)
	return bc
}

func getPrevBlockHash(t *testing.T, bc *Blockchain, heightOfCurrentBlock uint32) types.Hash {
	prevheader, err := bc.GetHeader(heightOfCurrentBlock - 1)
	assert.Nil(t, err)

	return BlockHasher{}.Hash(prevheader)
}
