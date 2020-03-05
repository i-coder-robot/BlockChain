package BlockChain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

//工作量 string -> hash
func ProofOfWork(data string) []byte {
	sha := sha256.New()
	sha.Write([]byte(data))
	x:= sha.Sum(nil)
	return x
}

//工作量,给定前几个位被0所占据
func ProofOfWorkWithDifficult(block *Block, numbers int) string {
	//拼接前n位0
	buf:=bytes.Buffer{}
	for i:=0;i<numbers;i++{
		buf.WriteString(string('0'))
	}
	result := buf.String()

	for{
		hash :=ComputeHash(block)
		//hash:=fmt.Sprintf("%x",ProofOfWork(data+string(block.Nonce)))
		block.Nonce++
		if hash[:numbers] == result {
			fmt.Println(hash)
			fmt.Println(block.Nonce)
			return hash
		}
	}
}