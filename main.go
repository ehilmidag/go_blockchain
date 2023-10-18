package main

import (
	"fmt"
	"strconv"
)

func main() {
	bc := NewBlockChain()
	bc.AddBlock("Send 1 BTC to Hillheim")
	bc.AddBlock("Send 2 more BTC to Hillheim")
	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PreviousBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Println()
	}

}
