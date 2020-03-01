package BlockChain

import (
	"testing"
)

func TestNew(t *testing.T) {
	newBlock := NewBlock("给小姐姐转账500","123456")
	newBlock.String()
}
