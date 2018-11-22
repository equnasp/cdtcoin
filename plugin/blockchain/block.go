package blockchain

import (
	"fmt"
	"github.com/equnasp/cdtcoin/cdtgo/sha256"
	"strconv"
	"time"
)

const Size = 32

type Block struct {
	Hash         string
	Data         string
	PreviousHash string
	TimeStamp    int64
	Nonce        int64
}

func New(previousHash string, data string) *Block {
	block := new(Block)
	block.Data = data
	block.PreviousHash = previousHash
	block.TimeStamp = time.Now().UnixNano()
	block.createHash(4)
	return block
}

/**
创建区块HASH
*/
func (block *Block) createHash(difficulty int) {
	var nonce int64
	for nonce = 1; ; nonce++ {
		record := []byte(block.PreviousHash + strconv.FormatInt(block.TimeStamp, 10) + strconv.FormatInt(nonce, 10) + block.Data)
		block.Nonce = nonce
		hash := sha256.Sum256(record)
		if block.difficulty(hash, difficulty) {
			block.Hash = fmt.Sprintf("%x", hash)
			break
		}
	}
	return
}

func (block *Block) difficulty(hash [Size]byte, d int) bool {
	dn := d / 2
	sn := d % 2
	for i := 0; i < dn; i++ {
		if hash[i] != 0x00 {
			return false
		}
	}
	if sn != 0 {
		if hash[dn*2+1] > 0x0f {
			return false
		}
	}
	return true
}
