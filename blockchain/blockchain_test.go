package blockchain_test

import (
	"testing"

	"github.com/acetimesolutions/blockchain-golang/block"
	"github.com/acetimesolutions/blockchain-golang/blockchain"
)

func TestBlockchainInit(t *testing.T) {

	var genesis block.Block
	genesis.Genesis()

	var blockchain blockchain.Blockchain
	blockchain.Init()

	if blockchain.Chain[0].Hash != genesis.Hash {
		t.Error("failed with init `Blockchain` becouse `Blockchain.Chain[0].Hash` is not equals to `GenesisBlock.Hash`")
		t.Fail()
	} else {
		t.Logf("Success: init the `Blockchain` test passed with expect")
		t.Log()
	}

}

func TestAddBlockToBlockchain(t *testing.T) {
	var genesis block.Block
	genesis.Genesis()

	var blockchain blockchain.Blockchain
	blockchain.Init()
	blockchain.AddBlock("foo")

	if blockchain.Chain[len(blockchain.Chain)-1].Data != "foo" {
		t.Error("Fail: failed with `Add new block` to `Blockchain` becouse `Blockchain.Chain[length - 1].Data` is not equals to `foo`")
		t.Fail()
	} else {
		t.Logf("Success: add `New Block` to `Blockchain` test passed with expect")
		t.Log()
	}
}
