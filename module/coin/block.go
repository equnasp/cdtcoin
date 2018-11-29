package coin

import (
	"bytes"
	"fmt"
	"github.com/equnasp/cdtcoin/cdtgo/sha256"
	"github.com/equnasp/cdtcoin/config"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Block Block information(区块信息)
type Block struct {
	// Header Header information of the current block(当前区块的头信息)
	Header *BlockHeader
	// PreviousHeader Header information of the previous block(上一个区块的头信息)
	PreviousHeader *BlockHeader
	// Data ...
	Data []byte
}

// BlockHeader Block header information(区块头信息)
type BlockHeader struct {
	// Version ...
	Version string
	// Hash Current block HASH(当前块的HASH)
	Hash string
	// PreviousHash Previous block of HASH(上一个块的HASH)
	PreviousHash string
	// Index Block index(区块索引)
	Index int64
	// Coinbase The address of the person who dug out the block(挖出该区块的人的地址)
	Coinbase string
	// Difficulty degree of difficulty(难度系数)
	Difficulty int
	// Nonce ...
	Nonce int64
	// TimeStamp time(时间)
	TimeStamp int64
}

// NewBlock Create a new block (创建新区块)
// Create a creation block when the value of previousHeader is nil(previousHeader的值为nil的时候，创建创世块)
func NewBlock(previousHeader *BlockHeader, index int64, data []byte) *Block {
	header := new(BlockHeader)
	if previousHeader == nil {
		header.Index = 0
	} else {
		header.Index = index + 1
	}

	header.Version = config.VERSION

	header.TimeStamp = time.Now().UnixNano()
	block := new(Block)

	if previousHeader == nil {
		block.Data = []byte("Genesis Block")
		block.PreviousHeader = new(BlockHeader)
	} else {
		block.PreviousHeader = previousHeader
		block.Data = data
	}

	block.Header = header
	block.generateBlockHash(1)
	return block
}

// generateBlockHash Generate block HASH(生成区块HASH)
func (b *Block) generateBlockHash(diff int) {
	header := b.Header
	previousHeader := b.PreviousHeader
	rand.Seed(time.Now().UnixNano())

	for {
		nonceTmp := rand.Int63n(999999999999999999)
		header.Nonce = nonceTmp
		header.Difficulty = diff

		version := []byte(header.Version)
		previousHash := []byte(previousHeader.Hash)
		timestamp := []byte(strconv.FormatInt(header.TimeStamp, 10))
		index := []byte(strconv.FormatInt(header.Index, 10))
		nonce := []byte(strconv.FormatInt(header.Nonce, 10))
		diff := []byte(strconv.Itoa(header.Difficulty))
		coinbase := []byte(header.Coinbase)

		headers := bytes.Join([][]byte{version, previousHash, index, coinbase, diff, nonce, timestamp, b.Data}, []byte{})

		hash := fmt.Sprintf("%x", sha256.Sum256(headers))
		if difficulty(hash, header.Difficulty) {
			header.Hash = hash
			break
		}
	}
	return
}

// difficulty ...
func difficulty(hash string, d int) bool {
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
