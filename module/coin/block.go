package coin

import (
	"github.com/equnasp/cdtcoin/config"
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
	Version []byte
	// Hash Current block HASH(当前块的HASH)
	Hash []byte
	// PreviousHash Previous block of HASH(上一个块的HASH)
	PreviousHash []byte
	// Index Block index(区块索引)
	Index int64
	// Coinbase The address of the person who dug out the block(挖出该区块的人的地址)
	Coinbase []byte
	// Difficulty degree of difficulty(难度系数)
	Difficulty int
	// Nonce ...
	Nonce int
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

	header.Version = []byte(config.VERSION)

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
	block.generateBlockHash(16)
	return block
}

// generateBlockHash Generate block HASH(生成区块HASH)
func (b *Block) generateBlockHash(difficulty int) {
	proofOfWork := NewProofOfWork(b, difficulty)
	b.Header.Nonce, b.Header.Hash = proofOfWork.Get(true)
	if !proofOfWork.Validate() {
		b.generateBlockHash(difficulty)
	}
	return
}
