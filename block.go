package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp         int64
	Data              []byte
	PreviousBlockHash []byte
	Hash              []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	fmt.Println(timestamp)
	headers := bytes.Join([][]byte{b.PreviousBlockHash, b.Data, timestamp}, []byte{})
	fmt.Println(headers)
	hash := sha256.Sum256(headers)
	fmt.Println(hash)
	b.Hash = hash[:]
	fmt.Println(b.Hash)
}

func NewBlock(data string, previousBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), previousBlockHash, []byte{}}
	block.SetHash()
	return block
}
