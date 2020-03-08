package main

type Transaction struct {
	From   string	`json:"from"`
	To     string	`json:"to"`
	Amount float64	`json:"amount"`
}

func (transaction *Transaction) String() string {
	return ToString(transaction)
}

func NewTransaction(from, to string, amount float64) *Transaction {
	return &Transaction{
		From:   from,
		To:     to,
		Amount: amount,
	}
}
