package main

import (
	"bytes"
	"encoding/gob"
	"time"
)

type Block struct {
	Transactions  []*Transaction
	PrevBlockHash []byte
	Timestamp     int64
	Hash          []byte
	Nonce         int
}

func NewBlock(txs []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{txs, prevBlockHash, time.Now().Unix(), []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer

	encoder := gob.NewEncoder(&res)
	encoder.Encode(b)

	return res.Bytes()
}

func Deserialize(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	decoder.Decode(&block)

	return &block
}
