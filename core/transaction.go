package core

import (
	"fmt"

	"github.com/sher2001/blockchain/crypto"
	"github.com/sher2001/blockchain/types"
)

type Transaction struct {
	data      []byte
	From      crypto.PublicKey
	Signature *crypto.Signature
	hash      types.Hash // cached version of transaction data hash
}

func NewTransaction(data []byte) *Transaction {
	return &Transaction{
		data: data,
	}
}

func (tx *Transaction) Hash(hasher Hasher[*Transaction]) types.Hash {
	if tx.hash.IsZero() {
		tx.hash = hasher.Hash(tx)
	}
	return tx.hash
}

func (tx *Transaction) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(tx.data)
	if err != nil {
		return err
	}

	tx.From = privKey.PublicKey()
	tx.Signature = sig

	return nil
}

func (tx *Transaction) Verify() error {
	if tx.Signature == nil {
		return fmt.Errorf("transaction has no signature")
	}
	if !tx.Signature.Verify(tx.From, tx.data) {
		return fmt.Errorf("invalid Transaction Signature")
	}
	return nil
}
