package block

import "fmt"

type Block struct {
	Timestamp string `json:"timestamp"`
	LastHash  string `json:"last_hash"`
	Hash      string `json:"hash"`
	Data      any    `json:"data"`
}

func (b *Block) Init(timestamp string, lasthash string, hash string, data any) Block {
	b.Timestamp = timestamp
	b.LastHash = lasthash
	b.Hash = hash
	b.Data = data

	return *b
}

func (b *Block) ToString() {
	fmt.Printf(`Block -
		Timestamp: %s
		Last Hash: %s
		Hash: %s
		Data: %s
	`, b.Timestamp, b.LastHash, b.Hash, b.Data)
}
