package wallet_info

import (
	"crypto/aes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"io"
	"testing"

	"crypto/cipher"
	"crypto/rand"

	"github.com/stretchr/testify/assert"
)

func TestEncryptPrivateKey(t *testing.T) {

	privateKey := "78bed97898a4e870fef8f2378a1fa93e695b70cd43c8e8aa75cd9d4b1e15f7b1"

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

func TestDecryptPrivateKey(t *testing.T) {
	ciphertext, err := hex.DecodeString("08513bf5ca21a797f079cdfea23cbcc335af51f8696279b38dcfa10db93f50e44e5d195dac5ac62378b50743fb3459aa")
	if err != nil {
		fmt.Println("Decode encrypt private key fail", "err", err)
	}
	nonceBytes, err := hex.DecodeString("fdc5610f30e0dca8dc32a1e2")
	if err != nil {
		fmt.Println("Decode nonce fail", "err", err)
	}

	hasher := sha256.New()
	hasher.Write([]byte(""))
	key := hasher.Sum(nil)

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("New cipher fail", "err", err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("New gcm", "err", err)
	}

	privateKeyBytes, err := gcm.Open(nil, nonceBytes, ciphertext, nil)
	if err != nil {
		fmt.Println("Open gcm fail", "err", err)
	}

	fmt.Println(hex.EncodeToString(privateKeyBytes))

	privateKey, _ := crypto.ToECDSA(privateKeyBytes)

	// 3. Generate address from public key
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Println("privateKeyBytes", hex.EncodeToString(privateKeyBytes), "address", address)
}
