package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vsivarajah/projectx-blockchain/core"
)

func TestTxPool(t *testing.T) {
	p := NewTxPool()
	assert.Equal(t, p.Len(), 0)
}

func TestTxPoolAddTx(t *testing.T) {
	p := NewTxPool()
	tx := core.NewTransaction([]byte("fooo"))
	assert.Nil(t, p.Add(tx))
	assert.Equal(t, p.Len(), 1)

	_ = core.NewTransaction([]byte("fooo"))
	assert.Equal(t, p.Len(), 1)

	p.Flush()
	assert.Equal(t, p.Len(), 0)
}
