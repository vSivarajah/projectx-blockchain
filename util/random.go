package util

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vsivarajah/projectx-blockchain/core"
	"github.com/vsivarajah/projectx-blockchain/crypto"
	"github.com/vsivarajah/projectx-blockchain/types"
)

func RandomBytes(size int) []byte {
	token := make([]byte, size)
	rand.Read(token)
	return token
}

func RandomHash() types.Hash {
	return types.HashFromBytes(RandomBytes(32))
}

// NewRandomTransaction return a new random transaction whithout signature.
func NewRandomTransaction(size int) *core.Transaction {
	return core.NewTransaction(RandomBytes(size))
}

func NewRandomTransactionWithSignature(t *testing.T, privKey crypto.PrivateKey, size int) *core.Transaction {
	tx := NewRandomTransaction(size)
	assert.Nil(t, tx.Sign(privKey))
	return tx
}
