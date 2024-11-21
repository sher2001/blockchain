package network

import (
	"github.com/sher2001/blockchain/core"
	"github.com/sher2001/blockchain/types"
)

type TxPool struct {
	transactions map[types.Hash]*core.Transaction
}

func NewTxPool() *TxPool {
	return &TxPool{
		transactions: make(map[types.Hash]*core.Transaction),
	}
}

// Add adds transaction to memory pool, caller is responsible to check if the
// transaction already exists in the memory pool
func (txP *TxPool) Add(tx *core.Transaction) error {
	txHash := tx.Hash(core.TransactionHasher{})
	txP.transactions[txHash] = tx
	return nil
}

func (txP *TxPool) Has(hash types.Hash) bool {
	_, ok := txP.transactions[hash]
	return ok
}

func (txP *TxPool) Length() int {
	return len(txP.transactions)
}

func (txP *TxPool) Flush() {
	txP.transactions = make(map[types.Hash]*core.Transaction)
}
