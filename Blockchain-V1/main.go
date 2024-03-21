package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct{
	blocks []*Block;
}

type Block struct {   // we will make a block then connect them together to create a chain of blocks
	Hash     []byte //a byte is an unsigned integer of 8 bytes which can only hold +ve values
	Data     []byte
	PrevHash []byte
}

func (b *Block) DerivedHash() {  //it is a method not function because of consisting of b as reciever or simply they basically act as class in go
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{}) //2-d slice joined with another slice with byte data type
	hash := sha256.Sum256(info)                                //returns a 256 bit hash for error detection
	b.Hash = hash[:];
}

func CreateBlock (data string, prevHash []byte) *Block{   //return value will be a pointer type
	block := &Block{[]byte{}, []byte(data), prevHash};    //so it is a block where first is Hash of current block, second is Data of current of Block and third is prevHash which is hash of previous block
	block.DerivedHash();    //[]byte(Data) means data is converted into unsigned integer 8 bit or byte
	return block;
}

func (chain *BlockChain) AddBlock(data string){    //this method takes input as string and can be used using Block variable as reciever to access the method
	prevBlock := chain.blocks[len(chain.blocks)-1];  //data of Previous block in a slice
	new := CreateBlock(data, prevBlock.Hash);    //new block is created here
	chain.blocks = append(chain.blocks, new);  //appending new created block in blockchain
}

func Genesis() *Block {    //a method to create first block with a random string data
	return CreateBlock("Genesis", []byte{});
}

func InitBlockChain() *BlockChain{    //creating an array of blocks and then refrencing it for BlockChain
	return &BlockChain{[]*Block{Genesis()}};
}

func main() {
	chain := InitBlockChain();  //staring of making of Blockchain by calling initblockchain function and variable gets a slice/array of blocks forming a chain

	chain.AddBlock("First Block After Genesis");
	chain.AddBlock("Second Block After Genesis");
	chain.AddBlock("Third Block After Genesis");

	for _,block := range chain.blocks{
		fmt.Printf("Previous Hash: %x\n", block.PrevHash);  //%x means we want to see data in hexadecimal integer
		fmt.Printf("Data in Block: %s\n", block.Data);    //%s means we want to see data in the form of string
		fmt.Printf("Hash: %x\n", block.Hash);
	}
}
