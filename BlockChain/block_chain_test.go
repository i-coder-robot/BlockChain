package BlockChain

import (
	"fmt"
	"testing"
	"time"
)

func TestOrigin(t *testing.T) {
	t1 :=&Transaction{
		From:   "",
		To:     "小明",
		Amount: 50.0,
	}
	b:=&Block{
		Transactions: []*Transaction{t1},
		PreBlockHash: "",
		Hash:         "",
		Nonce:        0,
		TimeStamp:    time.Now().Unix(),
	}
	o := NewBlockChain([]*Transaction{t1}, b,2,50)
	b.Hash = ComputeHash(b)
	o.String()
}

func TestBlockChain_AddBlockToChan(t *testing.T) {
	t1 :=&Transaction{
		From:   "",
		To:     "小明",
		Amount: 50.0,
	}
	b:=&Block{
		Transactions: []*Transaction{t1},
		PreBlockHash: "",
		Hash:         "",
		Nonce:        0,
		TimeStamp:    time.Now().Unix(),
	}
	o := NewBlockChain([]*Transaction{t1}, b,2,50)
	o.String()

	t2:=&Transaction{
		From:   "小明",
		To:     "小姐姐",
		Amount: 1.5,
	}
	block1 := &Block{
		Transactions: []*Transaction{t2},
		PreBlockHash: o.GetLatestBlockHash(),
		Hash:         "",
	}
	block1.Hash = ComputeHash(block1)
	o.AddBlockToChan(block1)

	t3 := &Transaction{
		From:   "小丁",
		To:     "小姐姐",
		Amount: 10.0,
	}

	block2 := &Block{
		Transactions:        []*Transaction{t3},
		PreBlockHash: o.GetLatestBlockHash(),
		Hash:         "",
	}
	block2.Hash= ComputeHash(block2)
	o.AddBlockToChan(block2)
	o.String()
}

func TestBlockChain_Validate(t *testing.T) {

	//验证区块链上只有一个祖先区块的时候
	t1 :=&Transaction{
		From:   "",
		To:     "小明",
		Amount: 50.0,
	}
	b:=&Block{
		Transactions: []*Transaction{t1},
		PreBlockHash: "",
		Hash:         "",
		Nonce:        0,
		TimeStamp:    time.Now().Unix(),
	}
	o := NewBlockChain([]*Transaction{t1}, b,2,50)
	o.String()

	t2:=&Transaction{
		From:   "小明",
		To:     "小姐姐",
		Amount: 1.5,
	}
	block1 := &Block{
		Transactions: []*Transaction{t2},
		PreBlockHash: o.GetLatestBlockHash(),
		Hash:         "",
	}
	block1.Hash = ComputeHash(block1)
	o.AddBlockToChan(block1)

	t3 := &Transaction{
		From:   "小丁",
		To:     "小姐姐",
		Amount: 10.0,
	}

	block2 := &Block{
		Transactions:        []*Transaction{t3},
		PreBlockHash: o.GetLatestBlockHash(),
		Hash:         "",
	}
	block2.Hash= ComputeHash(block2)
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
	t1:=&Transaction{
		From:   "老王",
		To:     "小姐姐",
		Amount: 1.0,
	}
	b:=&Block{
		Transactions: []*Transaction{t1},
		PreBlockHash: "",
		Hash:         "",
		Nonce:        0,
		TimeStamp:    time.Now().Unix(),
	}
	b.Hash	= ComputeHash(b)
	o:= NewBlockChain([]*Transaction{t1},b,2,50.0)
	o.Difficulty=2

	t2 := &Transaction{
		From:   "老王",
		To:     "电影院",
		Amount: 0.01,
	}

	block := &Block{
		Transactions:         []*Transaction{t2},
		PreBlockHash: o.GetLatestBlockHash(),
		Hash:         "",
		Nonce:        0,
	}
	block.Hash= ComputeHash(block)
	DigMine(block,o.Difficulty)

}
