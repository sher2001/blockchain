package core

import (
	"bytes"
	"testing"
	"time"

	"github.com/sher2001/blockchain/types"
	"github.com/stretchr/testify/assert"
)

func Test_header_encode_decode(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    10,
		Nonce:     9948761213,
	}

	buff := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buff))

	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buff))

	assert.Equal(t, h, hDecode)
}

func Test_block_encode_decode(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Height:    10,
			Nonce:     9948761213,
		},
		Transactions: nil,
	}

	buff := &bytes.Buffer{}
	assert.Nil(t, b.EncodeBinary(buff))

	bDecode := &Block{}
	assert.Nil(t, bDecode.DecodeBinary(buff))

	assert.Equal(t, b, bDecode)
}

func Test_block_hash(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Height:    10,
			Nonce:     9948761213,
		},
		Transactions: []Transaction{},
	}

	h := b.Hash()
	assert.False(t, h.IsZero())
}
