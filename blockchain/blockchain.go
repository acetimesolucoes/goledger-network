package blockchain

import (
	"fmt"

	"github.com/acetimesolutions/blockchain-golang/block"
)

type Blockchain struct {
	Chain []block.Block
}

func (b *Blockchain) Init() {

	var genesis block.Block
	genesis.Genesis()

	b.Chain = append(b.Chain, genesis)
}

func (b *Blockchain) AddBlock(data any) block.Block {

	var block block.Block
	block.MineBlock(b.Chain[len(b.Chain)-1], data)

	b.Chain = append(b.Chain, block)

	return block
}

func (b *Blockchain) IsValid(chain []block.Block) bool {
	var genesis block.Block
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

func (b *Blockchain) ReplaceChain(newChain []block.Block) {

	if len(newChain) <= len(b.Chain) {
		fmt.Print("Received chain is not longer than the current chain.")
		return
	} else if !b.IsValid(newChain) {
		fmt.Print("Received chain is not valid.")
		return
	} else {
		fmt.Print("Replacing blockchain with the new chain.")
		b.Chain = newChain
	}
}
