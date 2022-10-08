package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeypairSignVerifySuccess(t *testing.T) {
	privKey := GeneratePrivateKey()

	msg := []byte("hello world")
	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	assert.True(t, sig.Verify(privKey.PublicKey(), msg))
}

func TestKeypairSignVerifyFailed(t *testing.T) {
	privKey := GeneratePrivateKey()
	publicKey := privKey.PublicKey()

	msg := []byte("hello world")
	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	otherPrivKey := GeneratePrivateKey()
	otherPublicKey := otherPrivKey.PublicKey()

	assert.False(t, sig.Verify(otherPublicKey, msg))
	assert.False(t, sig.Verify(publicKey, []byte("wrong msg")))
}
