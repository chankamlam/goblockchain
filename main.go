package main

import "flag"

var (
	printBlockchain bool
	addBlock        string
)

func init() {
	flag.BoolVar(&printBlockchain, "print", false, "Print the blockchain")
	flag.StringVar(&addBlock, "add", "", "Add a block to the blockchain")
}

func main() {
	flag.Parse()
	chain, err := InitBlockChain()
	Handle(err)
	if addBlock != "" {
		chain.AddBlock(addBlock)
	}
	if printBlockchain {
		chain.Print()
	}

	// pow := CreateProofOfWork(chain.LastBlock())
	// fmt.Println(pow.target)

}
