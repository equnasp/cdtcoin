package coin

import (
	"bytes"
	"fmt"
	"github.com/equnasp/cdtcoin/cdtgo/sha256"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Block Block information(区块信息)
type Block struct {
	// Hash Current block HASH(当前块的HASH)
	Hash         string
	// PreviousHash Previous block of HASH(上一个块的HASH)
	PreviousHash string
	// Data ...
	Data         []byte
	// Index Block index value(区块索引值)
	Index        int64
	// TimeStamp time(时间)
	TimeStamp    int64
	// Nonce ...
	Nonce        int64
}

// NewGenesisBlock Generating a creation block(生成创世块)
func NewGenesisBlock() *Block {
	block := new(Block)
	block.Data = []byte("Genesis Block")
	block.Index = 0
	block.TimeStamp = time.Now().UnixNano()
	block.generateBlockHash(1)
	return block
}

// NewBlock Create a new block (创建新区块)
func NewBlock(previousHash string, index int64, data []byte) *Block {
	block := new(Block)
	block.PreviousHash = previousHash
	block.Data = data
	block.Index = index + 1
	block.TimeStamp = time.Now().UnixNano()
	block.generateBlockHash(1)
	return block
}

// generateBlockHash Generate block HASH(生成区块HASH)
func (b *Block) generateBlockHash(difficulty int) {
	var nonce int64

	rand.Seed(time.Now().UnixNano())
	for {
		nonce = rand.Int63n(999999999999999999)
		b.Nonce = nonce
		previousHash := []byte(b.PreviousHash)
		timestamp := []byte(strconv.FormatInt(b.TimeStamp, 10))
		index := []byte(strconv.FormatInt(b.Index, 10))
		nonce := []byte(strconv.FormatInt(b.Nonce, 10))

		headers := bytes.Join([][]byte{previousHash, b.Data, index, timestamp, nonce}, []byte{})

		hash := fmt.Sprintf("%x", sha256.Sum256(headers))
		if b.difficulty(hash, difficulty) {
			b.Hash = hash
			break
		}
	}
	return
}

// difficulty ...
func (block *Block) difficulty(hash string, d int) bool {
	for i := 0; i < d; i++ {
		if !strings.EqualFold(hash[i:i+1], "0") {
			return false
		}
	}

	if strings.EqualFold(hash[d:d+1], "0") {
		return false
	}

	return true
}
