package blockchain

import "github.com/acetimesolutions/blockchain-golang/block"

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
