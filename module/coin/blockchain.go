package coin

// Blockchain ...
type Blockchain struct {
	// blocks ...
	blocks []*Block
}

// NewBlockchain New blockchain(新的区块链)
func NewBlockchain() *Blockchain {
	blockchain := new(Blockchain)
	return blockchain
}

// Add Add data to the blockchain(将数据添加到区块链中)
func (blockchain *Blockchain) Add(data []byte) {
	// Get information about the previous block(获取上一个Block的信息)
	previousBlock := blockchain.blocks[len(blockchain.blocks)-1]

	// Create a new block(创建新Block)
	newBlock := NewBlock(previousBlock.Header, previousBlock.Header.Index, data)
	blockchain.blocks = append(blockchain.blocks, newBlock)
}
