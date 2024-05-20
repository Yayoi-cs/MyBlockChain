package blockchain

import "time"

type Transaction struct {
	Sender    string
	Recipient string
	Amount    int
}

type Block struct {
	Timestamp     time.Time
	Transactions  []Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

type Blockchain struct {
	Blocks          []Block
	TransactionPool []Transaction
	difficulty      int
}

func CreateBlockchain(difficulty int, currentTime time.Time) Blockchain {
	genesisBlock := Block{
		Hash:      []byte("0"),
		Timestamp: currentTime,
	}
	return Blockchain{
		/*[]Block{
			{
				Hash:      []byte("0"),
				Timestamp: currentTime,
			},
		},*/
		[]Block{genesisBlock},
		[]Transaction{},
		difficulty,
	}
}

func (bc *Blockchain) AddTransaction(transaction Transaction) {
	bc.TransactionPool = append(bc.TransactionPool, transaction)
}

func (bc *Blockchain) AddTransactionToBlock(transaction Transaction) {
	latestBlock := &Block{}

	if len(bc.Blocks) > 1 {
		latestBlock = &bc.Blocks[len(bc.Blocks)-1]
	} else if len(bc.Blocks) == 1 {
		latestBlock = &bc.Blocks[0]
	} else {
		panic("Blockchain must contain at least one Block.")
	}

	latestBlock.Transactions = append(latestBlock.Transactions, transaction)
}
