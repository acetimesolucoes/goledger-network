package test

import (
	"testing"

	"github.com/acetimesolutions/blockchain-golang/blockchain"
)

var blchn1 blockchain.Blockchain
var blchn2 blockchain.Blockchain

func init() {
	blchn1.Init()
	blchn2.Init()

	blchn1.AddBlock("ble")
	blchn1.AddBlock("bli")
	blchn1.AddBlock("blo")
	blchn1.AddBlock("blu")

	blchn2.AddBlock("ble")
	blchn2.AddBlock("bli")
	blchn2.AddBlock("blo")
	blchn2.AddBlock("blu")
}

func TestBlockchainInit(t *testing.T) {

	var genesis blockchain.Block
	genesis.Genesis()

	var blockchain blockchain.Blockchain
	blockchain.Init()

	if blockchain.Chain[0] != genesis {
		t.Error("failed with init `Blockchain` becouse `Blockchain.Chain[0].Hash` is not equals to `GenesisBlock.Hash`")
		t.Fail()
	} else {
		t.Log("Success: init the `Blockchain` test passed with expect")
		t.Log()
	}

}

func TestAddBlockToBlockchain(t *testing.T) {
	var genesis blockchain.Block
	genesis.Genesis()

	blchn1.AddBlock("foo")

	if blchn1.Chain[len(blchn1.Chain)-1].Data != "foo" {
		t.Error("Fail: failed with `Add new block` to `Blockchain` becouse `Blockchain.Chain[length - 1].Data` is not equals to `foo`")
		t.Fail()
	} else {
		t.Log("Success: add `New Block` to `Blockchain` test passed with expect")
		t.Log()
	}
}

func TestIsValidChain(t *testing.T) {

	isValid := blchn1.IsValid(blchn1.Chain)

	if !isValid {
		t.Error("Fail: failed with `Test in validation chain`")
		t.Fail()
	} else {
		t.Log("Success: `Test is valid chain` test passed with expect")
		t.Log()
	}
}

func TestCorruptedGenesisBlock(t *testing.T) {

	blchn1.Chain[0].Data = "Corrupted data block"

	isValid := blchn1.IsValid(blchn1.Chain)

	if isValid {
		t.Error("Fail: failed with `Test in validation chain` with corrupted data genesis")
		t.Fail()
	} else {
		t.Log("Success: invalidation chain test passed with expect")
		t.Log()
	}
}

func TestCorruptedChain(t *testing.T) {

	blchn1.Chain[3].Data = "Corrupted data block"

	isValid := blchn1.IsValid(blchn1.Chain)

	if isValid {
		t.Error("Fail: failed with `Test in validation chain` with corrupted `Block.data`")
		t.Fail()
	} else {
		t.Log("Success: invalidation chain test passed with expect")
		t.Log()
	}
}

func TestReplaceChain(t *testing.T) {

	blchn2.AddBlock("goo")
	blchn1.ReplaceChain(blchn2.Chain)

	// todo: validation with equals by not length equals
	if len(blchn1.Chain) != len(blchn2.Chain) {
		t.Error("Fail: failed with `Test in replacement chain` by valid `New Chain`")
		t.Fail()
	} else {
		t.Log("Success: replacement chain with `New Chain` test passed with expect")
		t.Log()
	}
}

func TestReclaceChainWithLessThanChan(t *testing.T) {
	blchn1.AddBlock("blee")
	blchn1.ReplaceChain(blchn2.Chain)

	if len(blchn1.Chain) == len(blchn2.Chain) {
		t.Error("Fail: failed with `Test in replacement chain` by less `New Chain`")
		t.Fail()
	} else {
		t.Log("Success: replacement chain with `Less New Chain` test passed with expect")
		t.Log()
	}
}
