package BlockChain

type Transaction struct {
	From      string
	To        string
	Amount    float64

}

func (transaction *Transaction) String() {
	ToString(transaction)
}

func NewTransaction(from,to string,amount float64) *Transaction {
	return &Transaction{
		From:      from,
		To:        to,
		Amount:    amount,
	}
}
