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

type Block struct {
	Hash         string
	PreviousHash string
	Data         []byte
	Index        int64
	TimeStamp    int64
	Nonce        int64
}

func NewBlock(previousHash string, data []byte) *Block {
	block := new(Block)
	block.Data = data
	block.PreviousHash = previousHash
	block.TimeStamp = time.Now().UnixNano()
	block.createHash(1)
	return block
}

/**
创建区块HASH
*/
func (b *Block) createHash(difficulty int) {
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
