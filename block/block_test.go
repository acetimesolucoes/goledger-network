package block_test

import (
	"fmt"
	"testing"
)

type Block struct {
	Timestamp string `json:"timestamp"`
	LastHash  string `json:"last_hash"`
	Hash      string `json:"hash"`
	Data      string `json:"data"`
}

func CreateBlock(t *testing.T) {
	block := new(Block)
	block.Data = ""
	block.Hash = ""
	block.LastHash = ""
	block.Timestamp = ""

	fmt.Printf("block: %v\n", block)

}
