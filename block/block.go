package block

import "fmt"

type Block struct {
	Timestamp string `json:"timestamp"`
	LastHash  string `json:"last_hash"`
	Hash      string `json:"hash"`
	Data      string `json:"data"`
}

func (b *Block) init(timestamp string, lasthash string, hash string, data string) {
	b.Timestamp = timestamp
	b.LastHash = lasthash
	b.Hash = hash
	b.Data = data
}

func (b *Block) toString() {
	fmt.Printf(`Block -
		Timestamp: %s
		Last Hash: %s
		Hash: %s
		Data: %s
	`, b.Timestamp, b.LastHash, b.Hash, b.Data)
}
