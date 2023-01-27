package pow

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/mitchellh/hashstructure"
	"log"
	"strconv"
	"strings"
	"time"
)

const difficulty = 1

// Block represents each 'item' in the blockchain
type Block struct {
	Index      int
	Timestamp  string
	Connection Connection
	Hash       string
	PrevHash   string
	Difficulty int
	Nonce      string
}

// Blockchain is a series of validated Blocks
var Blockchain []Block

// make sure block is valid by checking index, and comparing the hash of the previous block
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

// SHA256 hasing
func calculateHash(block Block) string {
	hash, err := hashstructure.Hash(block.Connection, nil)
	if err != nil {
		panic(err)
	}

	record := strconv.Itoa(block.Index) + block.Timestamp + fmt.Sprintf("%d", hash) + block.PrevHash + block.Nonce
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func CreateBlock(conn Connection) (Block, error) {
	return createNewBlock(Blockchain[len(Blockchain)-1], conn)
}

// create a new block using previous block's hash
func createNewBlock(oldBlock Block, conn Connection) (Block, error) {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Connection = conn
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Difficulty = difficulty

	for i := 0; ; i++ {
		hexBlock := fmt.Sprintf("%x", i)
		newBlock.Nonce = hexBlock
		if !isHashValid(calculateHash(newBlock), newBlock.Difficulty) {
			time.Sleep(100 * time.Millisecond)
			continue
		} else {
			fmt.Println(calculateHash(newBlock), " work done!")
			newBlock.Hash = calculateHash(newBlock)
			break
		}

	}

	if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
		Blockchain = append(Blockchain, newBlock)
	} else {
		return Block{}, fmt.Errorf("invalid block for block %s", newBlock.Hash)
	}

	return newBlock, nil
}

func isHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}

func GenesisBlock() {
	Blockchain = append(Blockchain, Block{
		Index:      0,
		Timestamp:  timeNow().String(),
		Connection: *NewConnection("", ""),
		Hash:       calculateHash(Block{}),
		PrevHash:   "",
		Difficulty: difficulty,
		Nonce:      "",
	})

	log.Println("Genesis block created:")
	spew.Dump(Blockchain[0])
}
