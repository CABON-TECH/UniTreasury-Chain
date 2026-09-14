package blockchain

import (
	"bytes"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// GenerateLeaf calculates the keccak256 hash of a single student's claim payload.
func GenerateLeaf(fundId *big.Int, studentHash [32]byte, trancheIndex *big.Int, recipient common.Address, amount *big.Int) []byte {
	// Equivalent to: keccak256(abi.encode(fundId, studentHash, trancheIndex, recipient, amount))
	// Then hashing it again to prevent second preimage attacks (OpenZeppelin standard)
	
	// Since we are using standard abi.encode, we must pad each element to 32 bytes
	encoded := make([]byte, 0, 32*5)
	encoded = append(encoded, common.LeftPadBytes(fundId.Bytes(), 32)...)
	encoded = append(encoded, studentHash[:]...)
	encoded = append(encoded, common.LeftPadBytes(trancheIndex.Bytes(), 32)...)
	encoded = append(encoded, common.LeftPadBytes(recipient.Bytes(), 32)...)
	encoded = append(encoded, common.LeftPadBytes(amount.Bytes(), 32)...)

	firstHash := crypto.Keccak256(encoded)
	return crypto.Keccak256(firstHash)
}

// GenerateTree computes the Merkle Root and returns the 2D array of levels.
func GenerateTree(leaves [][]byte) [][][]byte {
	if len(leaves) == 0 {
		return [][][]byte{}
	}

	// Sort leaves according to OpenZeppelin's standard (lexicographical sorting of hashes)
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
				// Odd number of leaves, promote to next level
				nextLevel = append(nextLevel, currentLevel[i])
			} else {
				a := currentLevel[i]
				b := currentLevel[i+1]
				// OZ standard: sort pairs before hashing
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

// GenerateProof generates a Merkle Proof for a specific leaf given the tree.
func GenerateProof(tree [][][]byte, leaf []byte) [][]byte {
	var proof [][]byte
	currentLeaf := leaf

	for levelIdx := 0; levelIdx < len(tree)-1; levelIdx++ {
		level := tree[levelIdx]
		
		// Find index of currentLeaf in level
		idx := -1
		for i, n := range level {
			if bytes.Equal(n, currentLeaf) {
				idx = i
				break
			}
		}

		if idx == -1 {
			return nil // Leaf not found
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
				// Promoted node
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
