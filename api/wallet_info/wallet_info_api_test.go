package wallet_info

import (
	"crypto/aes"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"testing"

	"crypto/cipher"
	"crypto/rand"

	"github.com/stretchr/testify/assert"
)

func TestEncryptPrivateKey(t *testing.T) {

	privateKey := ""

	hasher := sha256.New()
	hasher.Write([]byte(""))
	key := hasher.Sum(nil)
	t.Logf("Key length: %d", len(key))

	block, err := aes.NewCipher(key)
	assert.NoError(t, err)
	t.Logf("AES Block Size: %d bytes", block.BlockSize())

	gcm, err := cipher.NewGCM(block)
	assert.NoError(t, err)
	t.Logf("GCM Nonce Size: %d bytes", gcm.NonceSize())

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	assert.NoError(t, err)
	t.Logf("Generated Nonce (Hex): %x", nonce)

	privateKeyBytes, _ := hex.DecodeString(privateKey)

	ciphertext := gcm.Seal(nil, nonce, privateKeyBytes, nil)
	t.Logf("Encrypted Private Key (Hex): %x", ciphertext)
}
