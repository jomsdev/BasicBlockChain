package core

type Blockchain struct {
	Identities []string
	Blocks     []Block
}

// TODO: Document function
func NewBlockChain() *Blockchain {
	return &Blockchain{nil, nil}
}

// AddTransactions (unconfirmed) for generating the new block
func AddTransaction(transaction Transaction) {

}

// TODO: Document function
func AddBlock(block Block) {

}

// TODO: Document function
func CheckBalance(publicKey []byte) {

}

// TODO: Document function
func GetIdentityTransactions(publicKey []byte) []Transaction {
	return nil
}
