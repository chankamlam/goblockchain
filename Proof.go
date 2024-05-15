package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strconv"
)

const DIFFICULTY = 18

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func CreateProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-DIFFICULTY))
	return &ProofOfWork{block, target}
}

func (pow *ProofOfWork) PrepareData(nonce int64) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevHash,
			pow.block.Data,
			IntToHex(int64(DIFFICULTY)),
			IntToHex(nonce),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Run() ([]byte, int64) {
	var nonce int64 = 0
	var hashInt big.Int
	var hash [32]byte
	for {
		data := pow.PrepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
		fmt.Printf("\r%x", hash)
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()
	return hash[:], nonce
}

// convert int to hex
func IntToHex(n int64) []byte {
	return []byte(strconv.FormatInt(n, 16))
}
