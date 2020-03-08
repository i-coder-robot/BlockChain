package BlockChain

import BlockChain2 "github.com/i-coder-robot/BlockChain"

type Transaction struct {
	From   string
	To     string
	Amount float64
}

func (transaction *Transaction) String() string {
	return BlockChain2.ToString(transaction)
}

func NewTransaction(from, to string, amount float64) *Transaction {
	return &Transaction{
		From:   from,
		To:     to,
		Amount: amount,
	}
}
