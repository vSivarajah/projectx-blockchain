package network

import "github.com/vsivarajah/projectx-blockchain/core"

type GetBlocksMessage struct {
	From uint32

	// If To is 0 the maximum blocks will be returned
	To uint32
}

type BlocksMessage struct {
	Blocks []*core.Block
}
type GetStatusMessage struct{}

type StatusMessage struct {
	ID            string
	Version       uint32
	CurrentHeight uint32
}
