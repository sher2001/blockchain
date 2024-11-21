package network

import (
	"testing"

	"github.com/sher2001/blockchain/core"
	"github.com/stretchr/testify/assert"
)

func TestNewTxPool(t *testing.T) {
	txP := NewTxPool()
	assert.Equal(t, 0, txP.Length())
}

func TestTxPoolAddTx(t *testing.T) {
	txP := NewTxPool()
	tx := core.NewTransaction([]byte("foo"))

	assert.Nil(t, txP.Add(tx))
	assert.Equal(t, 1, txP.Length())

	tx_duplicate := core.NewTransaction([]byte("foo"))
	assert.Nil(t, txP.Add(tx_duplicate))
	assert.Equal(t, 1, txP.Length())
}
