package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vsivarajah/projectx-blockchain/crypto"
)

func TestAccountStateTransferNoBalance(t *testing.T) {
	state := NewAccountState()

	from := crypto.GeneratePrivateKey().PublicKey().Address()
	to := crypto.GeneratePrivateKey().PublicKey().Address()
	amount := uint64(90)
	assert.NotNil(t, state.Transfer(from, to, amount))
}

func TestAccountStateTransferSuccess(t *testing.T) {
	state := NewAccountState()
	from := crypto.GeneratePrivateKey().PublicKey().Address()

	state.AddBalance(from, 100)

	to := crypto.GeneratePrivateKey().PublicKey().Address()
	amount := uint64(90)
	assert.Nil(t, state.Transfer(from, to, amount))
}
