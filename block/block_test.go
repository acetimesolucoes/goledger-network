package block_test

import (
	"testing"

	"github.com/acetimesolutions/blockchain-golang/block"
)

func TestCreateBlock(t *testing.T) {
	block := new(block.Block).Init("bla", "ble", "bli", "blo")
	block.ToString()

	if block.Hash != "bla" {
		t.Log("TestCreateBlock pass with success")
	} else {
		t.Fail()
	}
}
