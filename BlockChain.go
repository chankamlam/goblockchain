package main

type BlockChain struct {
	Blocks []*Block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func (chain *BlockChain) LastBlock() *Block {
	return chain.Blocks[len(chain.Blocks)-1]
}

func CreateBlockChain() *BlockChain {
	return &BlockChain{[]*Block{CreateGenesisBlock()}}
}

// print the block chain
func (chain *BlockChain) Print() {
	for _, block := range chain.Blocks {
		block.Print()
	}
}
