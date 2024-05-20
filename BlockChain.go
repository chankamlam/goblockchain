package main

import (
	"github.com/dgraph-io/badger"
)

const (
	DBPATH       = "./tmp/blockchain.db"
	LAST_HASHKEY = "last_hash"
)

type BlockChain struct {
	// LastHash []byte
	Database *badger.DB
}

func (chain *BlockChain) AddBlock(data string) {

	lasthash, err := GetFromDB(chain.Database, []byte(LAST_HASHKEY))
	Handle(err)

	content, err := GetFromDB(chain.Database, lasthash)
	Handle(err)

	prevBlock := Deserialize(content)
	newBlock := CreateBlock(data, prevBlock.Hash)
	Handle(SetLastHashAndBlock(chain.Database, []byte(LAST_HASHKEY), newBlock.Hash, newBlock.Hash, newBlock.Serialize()))
}

func InitBlockChain() (*BlockChain, error) {

	opts := badger.DefaultOptions(DBPATH)
	opts.Dir = DBPATH
	opts.ValueDir = DBPATH

	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(txn *badger.Txn) error {
		_, err = GetFromDB(db, []byte(LAST_HASHKEY))
		if err != nil && err == badger.ErrKeyNotFound {
			genesis := CreateGenesisBlock()
			err := SetLastHashAndBlock(db, []byte(LAST_HASHKEY), genesis.Hash, genesis.Hash, genesis.Serialize())
			return err
		}
		return err
	})

	if err != nil {
		return nil, err
	}

	blockchain := BlockChain{db}
	return &blockchain, nil
}

// print the block chain
func (chain *BlockChain) Print() {
	lasthash, err := GetFromDB(chain.Database, []byte(LAST_HASHKEY))
	Handle(err)
	for {
		content, err := GetFromDB(chain.Database, lasthash)
		Handle(err)
		block := Deserialize(content)
		block.Print()
		lasthash = block.PrevHash
		if len(block.PrevHash) == 0 {
			break
		}
	}
}
