package coin

import (
	"bytes"
	"fmt"
	"github.com/equnasp/cdtcoin/cdtgo/sha256"
	"math"
	"math/big"
	"strconv"
)

// ProofOfWork represents a proof-of-work(工作证明)
type ProofOfWork struct {
	block      *Block
	target     *big.Int
	difficulty int
}

// NewProofOfWork ...
func NewProofOfWork(b *Block, difficulty int) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))

	pow := &ProofOfWork{b, target, difficulty}

	return pow
}

// prepareData ...
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.Header.Version,
			pow.block.PreviousHeader.Hash,
			[]byte(strconv.FormatInt(pow.block.Header.Index, 10)),
			pow.block.Header.Coinbase,
			[]byte(strconv.Itoa(pow.difficulty)),
			[]byte(strconv.Itoa(nonce)),
			[]byte(strconv.FormatInt(pow.block.Header.TimeStamp, 10)),
			pow.block.Data,
		},
		[]byte{},
	)

	return data
}

// Get Execute and get a hash and proof-of-work(执行并获得Hash和工作量证明)
func (pow *ProofOfWork) Get(display bool) (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining a new block")
	for nonce < math.MaxInt64 {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		if math.Remainder(float64(nonce), 1000000) == 0 && display {
			fmt.Printf("\rBlock:%x,Confirm:false,Nonce:%d\n", hash[:], nonce)
		}
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Printf("\rBlock:%x,Confirm:true,Nonce:%d\n", hash[:], nonce)
	return nonce, hash[:]
}

// Validate Block data validation(块的数据有效性验证)
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.block.Header.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
