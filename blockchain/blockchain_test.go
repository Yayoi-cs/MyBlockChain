package blockchain

import (
	"reflect"
	"testing"
	"time"
)

func TestCreateBlockchain(t *testing.T) {
	difficulty := 3
	currentTime := time.Now()

	genesisBlock := Block{
		Hash:      []byte("0"),
		Timestamp: currentTime,
	}
	want := Blockchain{
		[]Block{genesisBlock},
		[]Transaction{},
		difficulty,
	}
	got := CreateBlockchain(difficulty, currentTime)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("wanted %v, got %v", want, got)
	} else {
		t.Logf("Transaction pool contains %v transactions", got)
	}
}

func TestBlockchain_AddTransaction(t *testing.T) {
	transaction := Transaction{
		Sender:    "Alice",
		Recipient: "Bob",
		Amount:    1000,
	}
	t.Run("Add Transaction into empty Transaction Pool", func(t *testing.T) {
		bc := initialTransactionPool(transaction)
		got := bc.TransactionPool
		want := append([]Transaction{}, transaction)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("wanted %v, got %v", want, got)
		} else {
			t.Logf("Transaction pool contains %v transactions", got)
		}
	})
	t.Run("Add Transaction into exists Transaction Pool", func(t *testing.T) {
		bc := initialTransactionPool(transaction)
		bc.AddTransaction(transaction)

		transaction2 := Transaction{
			Sender:    "Bob",
			Recipient: "Alice",
			Amount:    200,
		}

		bc.AddTransaction(transaction2)

		got := bc.TransactionPool
		want := append([]Transaction{}, transaction, transaction2)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("wanted %v, got %v", want, got)
		} else {
			t.Logf("Transaction pool contains %v transactions", got)
		}
	})
}

func initialTransactionPool(transaction Transaction) Blockchain {
	difficulty := 3
	currentTime := time.Now()
	bc := CreateBlockchain(difficulty, currentTime)

	bc.AddTransaction(transaction)
	return bc
}

func TestBlockchain_AddTransactionToBlock(t *testing.T) {
	transaction := Transaction{
		Sender:    "Alice",
		Recipient: "Bob",
		Amount:    10,
	}
	bc := initialTransactionPool(transaction)

	transaction2 := Transaction{
		Sender:    "Bob",
		Recipient: "Alice",
		Amount:    20,
	}
	bc.AddTransaction(transaction2)

	txToAdd := bc.TransactionPool[0]
	bc.TransactionPool = bc.TransactionPool[1:]
	bc.AddTransactionToBlock(txToAdd)

	latestBlock := bc.Blocks[len(bc.Blocks)-1]

	if len(latestBlock.Transactions) != 1 {
		t.Errorf("Transaction in Block count must be 1, got %v", len(latestBlock.Transactions))
	}

	if len(bc.TransactionPool) != 1 {
		t.Errorf("TransactionPool count must be 1, got %v", len(bc.TransactionPool))
	}
}
