package blockchain

import (
	"fmt"
)

type Blockchain struct {
	Chain *[]Block
}

func (b *Blockchain) Init() {

	var genesis Block
	genesis.Genesis()

	b.Chain = new([]Block)
	*b.Chain = append(*b.Chain, genesis)
}

/* This method add new block to chain */
func (b *Blockchain) AddBlock(data any) Block {

	var block Block
	var lastBlock Block

	chainCopy := *b.Chain
	lastBlock = chainCopy[len(*b.Chain)-1]

	block.MineBlock(&lastBlock, data)

	*b.Chain = append(*b.Chain, block)

	return block
}

func (b *Blockchain) IsValid(chain []Block) bool {
	var genesis Block
	genesis.Genesis()

	if chain[0] != genesis {
		return false
	}

	for i := 1; i < len(chain); i++ {
		block := chain[i]
		lastBlock := chain[i-1]

		if block.LastHash != lastBlock.Hash || block.Hash != block.BlockHash(block) {
			return false
		}
	}

	return true
}

func (b *Blockchain) ReplaceChain(newChain []Block) {

	if len(newChain) <= len(*b.Chain) {
		fmt.Println("Received chain is not longer than the current chain.")
		return
	} else if !b.IsValid(newChain) {
		fmt.Println("Received chain is not valid.")
		return
	} else {
		fmt.Println("Replacing blockchain with the new chain.")
		*b.Chain = newChain
	}
}
