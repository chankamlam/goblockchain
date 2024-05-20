package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	nonce    int64
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

func (block *Block) Serialize() []byte {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(block)
	Handle(err)
	return buf.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	Handle(err)
	return &block
}

func (block *Block) Print() {
	// PrintBlock will print the block
	fmt.Println()
	fmt.Printf("Data: %s\n", block.Data)
	fmt.Printf("PrevHash: %x\n", block.PrevHash)
	fmt.Printf("Hash: %x\n", block.Hash)
	fmt.Println()
}
