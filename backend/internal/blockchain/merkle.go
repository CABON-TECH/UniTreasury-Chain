package blockchain
import (
	"bytes"
	"math/big"
	"sort"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)
func GenerateLeaf(fundId *big.Int, studentHash [32]byte, trancheIndex *big.Int, recipient common.Address, amount *big.Int) []byte {
	encoded := make([]byte, 0, 32*5)
	encoded = append(encoded, common.LeftPadBytes(fundId.Bytes(), 32)...)
	encoded = append(encoded, studentHash[:]...)
	encoded = append(encoded, common.LeftPadBytes(trancheIndex.Bytes(), 32)...)
	encoded = append(encoded, common.LeftPadBytes(recipient.Bytes(), 32)...)
	encoded = append(encoded, common.LeftPadBytes(amount.Bytes(), 32)...)
	firstHash := crypto.Keccak256(encoded)
	return crypto.Keccak256(firstHash)
}
func GenerateTree(leaves [][]byte) [][][]byte {
	if len(leaves) == 0 {
		return [][][]byte{}
	}
	sortedLeaves := make([][]byte, len(leaves))
	copy(sortedLeaves, leaves)
	sort.Slice(sortedLeaves, func(i, j int) bool {
		return bytes.Compare(sortedLeaves[i], sortedLeaves[j]) < 0
	})
	tree := [][][]byte{sortedLeaves}
	currentLevel := sortedLeaves
	for len(currentLevel) > 1 {
		var nextLevel [][]byte
		for i := 0; i < len(currentLevel); i += 2 {
			if i+1 == len(currentLevel) {
				nextLevel = append(nextLevel, currentLevel[i])
			} else {
				a := currentLevel[i]
				b := currentLevel[i+1]
				if bytes.Compare(a, b) > 0 {
					a, b = b, a
				}
				nextLevel = append(nextLevel, crypto.Keccak256(append(a, b...)))
			}
		}
		tree = append(tree, nextLevel)
		currentLevel = nextLevel
	}
	return tree
}
func GenerateProof(tree [][][]byte, leaf []byte) [][]byte {
	var proof [][]byte
	currentLeaf := leaf
	for levelIdx := 0; levelIdx < len(tree)-1; levelIdx++ {
		level := tree[levelIdx]
		idx := -1
		for i, n := range level {
			if bytes.Equal(n, currentLeaf) {
				idx = i
				break
			}
		}
		if idx == -1 {
			return nil 
		}
		var sibling []byte
		var parent []byte
		if idx%2 == 0 {
			if idx+1 < len(level) {
				sibling = level[idx+1]
				a, b := currentLeaf, sibling
				if bytes.Compare(a, b) > 0 {
					a, b = b, a
				}
				parent = crypto.Keccak256(append(a, b...))
				proof = append(proof, sibling)
			} else {
				parent = currentLeaf
			}
		} else {
			sibling = level[idx-1]
			a, b := sibling, currentLeaf
			if bytes.Compare(a, b) > 0 {
				a, b = b, a
			}
			parent = crypto.Keccak256(append(a, b...))
			proof = append(proof, sibling)
		}
		currentLeaf = parent
	}
	return proof
}
