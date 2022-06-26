package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

type Block struct {
	Timestamp int64  `json:"timestamp"`
	LastHash  string `json:"last_hash"`
	Hash      string `json:"hash"`
	Data      any    `json:"data"`
}

func (b *Block) init(timestamp int64, lasthash string, hash string, data any) {
	b.Timestamp = timestamp
	b.LastHash = lasthash
	b.Hash = hash
	b.Data = data
}

func (b *Block) ToString() {
	fmt.Printf(`
	Block -
	Timestamp: %d
	LastHash: %s
	Hash: %s
	Data: %s`, b.Timestamp, b.LastHash, b.Hash, b.Data)

	fmt.Print("\r\r")
}

func (b *Block) Genesis() {
	b.init(1656202080635360013, "-----", "f1r57-h45h", [0]string{})
}

func (b *Block) MineBlock(lastBlock Block, data any) {

	timestamp := time.Now().UnixNano()
	lastHash := lastBlock.Hash
	hash := b.Hasher(timestamp, lastHash, data)

	b.init(timestamp, lastHash, hash, data)
}

func (b *Block) Hasher(timestamp int64, lastHash string, data any) string {

	str := strconv.FormatInt(timestamp, 10)

	jsonData, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	byteArray := [][]byte{[]byte(str), []byte(lastHash), jsonData}
	strToHash := bytes.Join(byteArray, []byte("-"))

	buf := sha256.Sum256(strToHash)

	return hex.EncodeToString(bytes.NewBuffer(buf[:]).Bytes())
}

func (b *Block) BlockHash(block Block) string {
	return block.Hasher(block.Timestamp, block.LastHash, block.Data)
}
