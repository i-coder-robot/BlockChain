package BlockChain

import (
	"fmt"
	"testing"
)

func TestOrigin(t *testing.T) {
	o := Origin("我时老祖宗区块", "")
	o.String()
}

func TestBlockChain_AddBlockToChan(t *testing.T) {
	o := Origin("我时老祖宗区块", "")
	desc1 := "小明给小姐姐转账300元"
	preBlockHash1 :=o.GetLatestBlockHash()
	block1 := &Block{
		Data:        desc1,
		PreBlockHash: preBlockHash1,
		Hash:         ComputeHash(desc1,preBlockHash1),
	}
	o.AddBlockToChan(block1)

	desc2 := "小丁给小姐姐转账500元"
	preBlockHash2 :=o.GetLatestBlockHash()
	block2 := &Block{
		Data:        desc2,
		PreBlockHash: preBlockHash2,
		Hash:         ComputeHash(desc1,preBlockHash1),
	}
	o.AddBlockToChan(block2)
	o.String()
}

func TestBlockChain_Validate(t *testing.T) {

	//验证区块链上只有一个祖先区块的时候
	o := Origin("我时老祖宗区块", "")
	result := o.Validate()
	fmt.Println(result)

	//验证区块链上有多个区块的时候
	desc1 := "小明给小姐姐转账300元"
	preBlockHash1 :=o.GetLatestBlockHash()
	block1 := &Block{
		Data:        desc1,
		PreBlockHash: preBlockHash1,
		Hash:         ComputeHash(desc1,preBlockHash1),
	}
	o.AddBlockToChan(block1)

	desc2 := "小丁给小姐姐转账500元"
	preBlockHash2 :=o.GetLatestBlockHash()
	block2 := &Block{
		Data:        desc2,
		PreBlockHash: preBlockHash2,
		Hash:         ComputeHash(desc2,preBlockHash2),
	}
	o.AddBlockToChan(block2)
	r := o.Validate()
	fmt.Println(r)

	//有人篡改数据
	//block1.Data="小明给小姐姐转账100元"
	//o.Validate()  //结果屏幕打印 数据被别人修改了

	//篡改数据并修改当前数据对应的hash值
	//block1.Data="小明给小姐姐转账100元"
	//block1.Hash=ComputeHash(block1.Data,o.Blocks[0].Hash)
	//o.Validate() //结果屏幕打印 区块链子断裂

}

func TestDigMine(t *testing.T) {
	o := Origin("我时老祖宗区块", "")
	//o.String()
	o.Difficulty=4
	b:=&Block{
		Data:         "今天要挖矿，心情好激动",
		PreBlockHash: o.GetLatestBlockHash(),
		Hash:         "",
		Nonce:        0,
	}
	DigMine(b,o.Difficulty)
}
