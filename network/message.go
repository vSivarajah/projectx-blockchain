package network

type GetStatusMessage struct{}

type StatusMessage struct {
	ID            string
	Version       uint32
	CurrentHeight uint32
}
