package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	nonce    int64
}

func (b *Block) DeriveHash() {
	// DeriveHash will generate a hash for the block
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := CreateProofOfWork(block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.nonce = nonce
	return block
}

func CreateGenesisBlock() *Block {
	return CreateBlock("Genesis Block", []byte{})
}

func (block *Block) Print() {
	// PrintBlock will print the block
	fmt.Println()
	fmt.Printf("Data: %s\n", block.Data)
	fmt.Printf("PrevHash: %x\n", block.PrevHash)
	fmt.Printf("Hash: %x\n", block.Hash)
	fmt.Println()
}
