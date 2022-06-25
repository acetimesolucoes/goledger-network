package main

import (
	"github.com/acetimesolutions/blockchain-golang/block"
)

func main() {

	var blockInstance block.Block
	var genesis *block.Block
	genesis.Genesis()

	blockInstance.MineBlock(*genesis, []string{"bla", "ble", "bli"})
	blockInstance.ToString()
}
