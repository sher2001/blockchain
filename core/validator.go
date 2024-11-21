package core

import "fmt"

type Validator interface {
	ValidateBlock(*Block) error
}

type BlockValidator struct {
	bc *Blockchain
}

func NewBlockValidator(bc *Blockchain) *BlockValidator {
	return &BlockValidator{
		bc: bc,
	}
}

func (bv *BlockValidator) ValidateBlock(b *Block) error {
	if bv.bc.HasBlock(b.Height) {
		return fmt.Errorf("chain already contains block (%d) with hash (%s)", b.Height, b.Hash(BlockHasher{}))
	}

	if err := b.Verify(); err != nil {
		return err
	}

	if b.Height != (bv.bc.Height() + 1) {
		return fmt.Errorf("block (%s) too much high", b.Hash(BlockHasher{}))
	}

	prevHeader, err := bv.bc.GetHeader(b.Height - 1)
	if err != nil {
		return err
	}
	prevHeaderHash := BlockHasher{}.Hash(prevHeader)

	if prevHeaderHash != b.PrevBlockHash {
		return fmt.Errorf("the previous block hash (%s) is invalid", b.PrevBlockHash)
	}

	return nil
}
