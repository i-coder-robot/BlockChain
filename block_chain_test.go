package BlockChain

import (
	"fmt"
	"testing"
	"time"
)

func TestOrigin(t *testing.T) {
	//因为是创世纪的区块，那么转账是系统给的，所以from是空的
	t1:=&Transaction{
		From:   "",
		To:     "我时老祖宗区块",
		Amount: 50.0,
	}
	b:=&Block{
		Transactions: []*Transaction{t1},
		PreBlockHash: "",
		Hash:         "",
		Nonce:        0,
		TimeStamp:    time.Now().Unix(),
	}
	blocks := []*Block{b}
	o := NewBlockChain([]*Transaction{t1}, blocks,2,50)
	o.String()
}

func TestBlockChain_AddBlockToChan(t *testing.T) {
	t0:=&Transaction{
		From:   "",
		To:     "我时老祖宗区块",
		Amount: 50.0,
	}
	b0:=&Block{
		Transactions: []*Transaction{t0},
		PreBlockHash: "",
		Hash:         "",
		Nonce:        0,
		TimeStamp:    time.Now().Unix(),
	}
	b0.Hash=ComputeHash(b0)
	blocks := []*Block{b0}
	o := NewBlockChain([]*Transaction{t0}, blocks,2,50)

	t1 := &Transaction{
		From:   "小明",
		To:     "小姐姐",
		Amount: 3.68,
	}
	block1 := &Block{
		Transactions: []*Transaction{t1},
		PreBlockHash: o.GetLatestBlockHash(),
		Hash:         "",
	}
	block1.Hash = ComputeHash(block1)
	o.AddBlockToChan(block1)

	t2 := &Transaction{
		From:   "小丁",
		To:     "小姐姐",
		Amount: 5.25,
	}
	b2 := &Block{
		Transactions: []*Transaction{t2},
		PreBlockHash: o.GetLatestBlockHash(),
		Hash:         "",
		Nonce:        0,
		TimeStamp:    time.Now().Unix(),
	}
	b2.Hash=ComputeHash(b2)
	o.AddBlockToChan(b2)

	o.String()
}

func TestBlockChain_Validate(t *testing.T) {

	//验证区块链上只有一个祖先区块的时候
	t0:=&Transaction{
		From:   "",
		To:     "我时老祖宗区块",
		Amount: 50.0,
	}
	b0:=&Block{
		Transactions:[]*Transaction{t0},
		PreBlockHash:"",
		Hash:"",
		Nonce       :0,
		TimeStamp    :time.Now().Unix(),
	}
	b0.Hash=ComputeHash(b0)

	o := NewBlockChain([]*Transaction{t0}, []*Block{b0},2,50)

	result := o.Validate()
	fmt.Println(result)

	//验证区块链上有多个区块的时候

	t1 := &Transaction{
		From:   "小明",
		To:     "小姐姐",
		Amount: 3.68,
	}
	block1 := &Block{
		Transactions: []*Transaction{t1},
		PreBlockHash: o.GetLatestBlockHash(),
		Hash:         "",
	}
	block1.Hash = ComputeHash(block1)
	o.AddBlockToChan(block1)

	t2 := &Transaction{
		From:   "小丁",
		To:     "小姐姐",
		Amount: 5.25,
	}
	b2 := &Block{
		Transactions: []*Transaction{t2},
		PreBlockHash: o.GetLatestBlockHash(),
		Hash:         "",
		Nonce:        0,
		TimeStamp:    time.Now().Unix(),
	}
	b2.Hash=ComputeHash(b2)
	o.AddBlockToChan(b2)

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
	t0:=&Transaction{
		From:   "",
		To:     "我时老祖宗区块",
		Amount: 50.0,
	}
	b0 := &Block{
		Transactions: []*Transaction{t0},
		PreBlockHash: "",
		Hash:         "",
		Nonce:        0,
		TimeStamp:    time.Now().Unix(),
	}
	b0.Hash=ComputeHash(b0)
	o := NewBlockChain([]*Transaction{t0}, []*Block{b0},2,50)

	//o.ToString()
	o.Difficulty = 4
	t1:=&Transaction{
		From:   "小明",
		To:     "海底捞",
		Amount: 0.52,
	}
	b := &Block{
		Transactions: []*Transaction{t1},
		PreBlockHash: o.GetLatestBlockHash(),
		Hash:         "",
		Nonce:        0,
	}
	b.Hash=ComputeHash(b)
	DigMine(b, o.Difficulty)
}
