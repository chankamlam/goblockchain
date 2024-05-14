package main

func main() {
	chain := CreateBlockChain()
	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	chain.Print()
}
