package BlockChain

import "bytes"

//前n位都是0
func PreZero(n int) string {
	buf:=bytes.Buffer{}
	for i:=0;i<n;i++{
		buf.WriteString(string('0'))
	}
	result := buf.String()
	return result
}

//挖矿
func (block *Block) DigMine(n int){
	for{
		hash := ComputeHash(block.Data,block.PreBlockHash)
		if PreZero(n)==hash[:n]{
			block.Hash = hash
			break
		}
	}
}