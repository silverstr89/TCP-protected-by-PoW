package pow

import (
	"testing"
)

func TestNewBlock(t *testing.T) {
	// Initialize the blockchain with the Genesis block
	GenesisBlock()

	// Create a test connection
	conn := NewConnection("Alice", "Bob")

	// Call the CreateBlock function and check for errors
	block, err := CreateBlock(*conn)
	if err != nil {
		t.Errorf("Error creating new block: %s", err)
	}

	// Check that the new block has the correct index and prevHash
	if block.Index != 1 {
		t.Errorf("Expected block index to be 1, got %d", block.Index)
	}
	if block.PrevHash != Blockchain[0].Hash {
		t.Errorf("Expected block prevHash to be %s, got %s", Blockchain[0].Hash, block.PrevHash)
	}

	// Check that the new block's hash is valid
	if !isHashValid(block.Hash, difficulty) {
		t.Errorf("Expected block hash to be valid, got %s", block.Hash)
	}

	// Check that the new block is valid
	if !isBlockValid(block, Blockchain[0]) {
		t.Errorf("Expected block to be valid, got %s", block.Hash)
	}
}
