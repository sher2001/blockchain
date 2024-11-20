package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

	"github.com/sher2001/blockchain/crypto"
	"github.com/sher2001/blockchain/types"
)

type Header struct {
	Version       uint32
	DataHash      types.Hash // hash of all transactions
	PrevBlockHash types.Hash
	Height        uint32
	Timestamp     int64
}

type Block struct {
	*Header
	Transactions []Transaction
	Validator    crypto.PublicKey
	Signature    *crypto.Signature
	hash         types.Hash // Cached Version of header hash
}

func NewBlock(h *Header, txx []Transaction) *Block {
	return &Block{
		Header:       h,
		Transactions: txx,
	}
}

func (b *Block) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(b.HeaderData())
	if err != nil {
		return err
	}

	b.Signature = sig
	b.Validator = privKey.PublicKey()
	return nil
}

func (b *Block) Verify() error {
	if b.Signature == nil {
		return fmt.Errorf("block has no signature")
	}
	if !b.Signature.Verify(b.Validator, b.HeaderData()) {
		return fmt.Errorf("invalid block signature")
	}
	return nil
}

func (b *Block) Decode(r io.Reader, decoder Decoder[*Block]) error {
	return decoder.Decode(r, b)
}

func (b *Block) Encode(w io.Writer, encoder Encoder[*Block]) error {
	return encoder.Encode(w, b)
}

func (b *Block) Hash(hasher Hasher[*Block]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b)
	}
	return b.hash
}

func (b *Block) HeaderData() []byte {
	buff := &bytes.Buffer{}
	enc := gob.NewEncoder(buff)
	enc.Encode(b.Header)
	return buff.Bytes()
}
