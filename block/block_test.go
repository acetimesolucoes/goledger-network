package block_test

import (
	"testing"

	"github.com/acetimesolutions/blockchain-golang/block"
)

func TestCreateBlock(t *testing.T) {
	var newBlock block.Block

	var genesis block.Block
	genesis.Genesis()

	newBlock.MineBlock(genesis, "bla")

	if newBlock.Data == genesis.Data {
		t.Error("Fail: the `block.data` is equals to `lastblock` data")
		t.Fail()
	}

	if newBlock.LastHash != genesis.Hash {
		t.Error("Fail: the `block.lasthash` is not equals to `lastblock.hash`")
		t.Fail()
	} else {
		t.Logf("Success: create block test passed with expect")
	}
}
