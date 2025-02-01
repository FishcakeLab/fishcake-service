package wallet_info

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"testing"

	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/FishcakeLab/fishcake-service/database"
	"github.com/FishcakeLab/fishcake-service/rpc/account"
	"github.com/FishcakeLab/fishcake-service/service/reward_service"
	"github.com/FishcakeLab/fishcake-service/service/rpc_service"

	"github.com/ethereum/go-ethereum/rlp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

type WalletConfig struct {
	EncryptedPrivateKey string `yaml:"encrypted_private_key"`
	Nonce               string `yaml:"nonce"`
	KeyPhrase           string `yaml:"key_phrase"`
}

func TestSendTx(t *testing.T) {
	// Check wallet address and send reward transaction
	cfg, err := config.New("../../config.yaml")
	println("err:", err)
	db, _ := database.NewDB(cfg)
	addr := "0x67936bb11f8fd1d25da1f94e0aa51039409a7c97"
	isAdd := db.WalletInfoDB.SelectWalletAddr(addr) // Check if address exists
	println(isAdd)
	if true == false {
		println("Already exists, no reward")
	} else {
		// Get decrypted key
		privateKey, err := reward_service.NewRewardService("config.yaml").DecryptPrivateKey()
		if err != nil {
			log.Info("decrypt private key error: %v", err)
			return
		}

		// Convert amount to big.Int
		amount := new(big.Int).SetUint64(1)

		// Call transfer method
		signedTxData, txHash, err := reward_service.NewRewardService("").CreateOfflineTransaction(
			reward_service.FCC,
			privateKey,
			addr,
			amount,
		)
		println("signedTxData=====", signedTxData)
		println("txhash=====", txHash)

		// Send transaction request
		req := &account.SendTxRequest{
			Chain:   "Polygon",
			Network: "mainnet",
			RawTx:   signedTxData,
		}

		sendtx, _ := rpc_service.NewRpcService(cfg.RpcUrl).SendTx(context.Background(), req)
		println(sendtx.TxHash)
	}
}

func TestWalletInfoApi2(t *testing.T) {
	privateKey, err := reward_service.NewRewardService("").DecryptPrivateKey() //获取解密密钥
	if err != nil {
		log.Info("decrypt private key error: %v", err)

	}

	amount := new(big.Int).Mul(
		big.NewInt(50),
		new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil),
	)

	txHex, txHash, err := reward_service.NewRewardService("").CreateOfflineTransaction(
		reward_service.FCC,
		privateKey,
		"0x67936bb11f8fd1d25da1f94e0aa51039409a7c97",
		amount,
	)
	log.Info(txHex, txHash)
	req :=
		&account.SendTxRequest{
			Chain:   "Polygon",
			Network: "mainnet",
			RawTx:   txHex,
		}

	cfg, err := config.New("./config.yaml")

	sendtx, err := rpc_service.NewRpcService(cfg.RpcUrl).SendTx(context.Background(), req)

	if err != nil {
		fmt.Printf("RPC send tx error: %v", err.Error())
	} else {
		log.Info("RPC send tx: %v", sendtx.String())
	}
}

func TestWalletInfo(t *testing.T) {
	cfg, _ := config.New("../../config.yaml")

	req := &account.TxAddressRequest{
		Chain:   "Polygon",
		Network: "mainnet",
		Address: "0x8916B42a4DB16CA71080dBB0f3650162Ad1E7e3e",
	}

	resp, _ := rpc_service.NewRpcService(cfg.RpcUrl).GetTxByAddress(context.Background(), req)
	fmt.Println(resp.Tx)
}

// Test Case 1: Generate Polygon private key and address
func TestGeneratePolygonWallet(t *testing.T) {
	// Generate private key
	privateKey, err := crypto.GenerateKey()
	assert.NoError(t, err)

	// Get public key from private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	assert.True(t, ok)

	// Get address from public key
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Convert private key to hex string
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hex.EncodeToString(privateKeyBytes)

	t.Logf("Private Key: %s", privateKeyHex)
	t.Logf("Address: %s", address.Hex())
}

// === RUN   TestGeneratePolygonWallet
// wallet_info_api_test.go:41: Private Key: 030d1048c7824726a186881df8b49ffe4bb6ff1c28b43cd00ab3fb8d0267c339
// wallet_info_api_test.go:42: Address: 0x53c4025CCD16682b9b6B96F38B01AFe0e8e447CD
// Test Case 2: Encrypt private key and save to YAML
func TestEncryptPrivateKey(t *testing.T) {
	// 1. Read config file for key phrase
	yamlFile, err := ioutil.ReadFile("../config.yaml")
	assert.NoError(t, err)

	var configRead WalletConfig
	err = yaml.Unmarshal(yamlFile, &configRead)
	assert.NoError(t, err)

	// 1. Generate new random private key
	privateKey, err := crypto.GenerateKey()
	assert.NoError(t, err)

	// 2. Get public key from private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	assert.True(t, ok)

	// 3. Generate address from public key
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 4. Convert private key to bytes
	privateKeyBytes := crypto.FromECDSA(privateKey)

	// Print original private key and address info
	t.Logf("Original Private Key (Hex): %x", privateKeyBytes)
	t.Logf("Wallet Address: %s", address.Hex())

	// 5. Use key phrase from config
	hasher := sha256.New()
	hasher.Write([]byte(configRead.KeyPhrase))
	key := hasher.Sum(nil)
	t.Logf("Key length: %d", len(key))

	// 6. Create AES cipher
	block, err := aes.NewCipher(key)
	assert.NoError(t, err)
	t.Logf("AES Block Size: %d bytes", block.BlockSize())

	// 7. Create GCM mode
	gcm, err := cipher.NewGCM(block)
	assert.NoError(t, err)
	t.Logf("GCM Nonce Size: %d bytes", gcm.NonceSize())

	// 8. Generate random nonce
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	assert.NoError(t, err)
	t.Logf("Generated Nonce (Hex): %x", nonce)

	// 9. Encrypt private key
	ciphertext := gcm.Seal(nil, nonce, privateKeyBytes, nil)
	t.Logf("Encrypted Private Key (Hex): %x", ciphertext)

	// 10. Create config structure
	config := WalletConfig{
		EncryptedPrivateKey: hex.EncodeToString(ciphertext),
		Nonce:               hex.EncodeToString(nonce),
		KeyPhrase:           configRead.KeyPhrase,
	}

	// 11. Save config to YAML file
	yamlData, err := yaml.Marshal(&config)
	assert.NoError(t, err)
	t.Logf("YAML Data: %s", string(yamlData))

	err = ioutil.WriteFile("./wallet_config.yaml", yamlData, 0644)
	assert.NoError(t, err)
	t.Log("Wallet config saved to wallet_config.yaml")

	// 13. Print summary
	t.Log("=== Wallet Summary ===")
	t.Logf("Address: %s", address.Hex())
	t.Logf("Private Key (Hex): %x", privateKeyBytes)
	t.Logf("Encrypted Private Key Length: %d bytes", len(ciphertext))
}

// wallet_info_api_test.go:116: Address: 0xd024e2e611A698eB94127e5a31eC23B7E82a0725
// wallet_info_api_test.go:117: Private Key (Hex): 4275e3513a75cbd9b080fbc6c1a9d41095c2f2fc64ecf34a52d8be45cede5592
// Test Case 3: Read from YAML and decrypt private key
func TestDecryptPrivateKey(t *testing.T) {
	println("50:====0x14000347f000", new(big.Int).SetUint64(50*1000000))
	yamlFile, err := ioutil.ReadFile("../../config.yaml")
	assert.NoError(t, err)
	t.Logf("Read config file: %s", string(yamlFile))

	var config struct {
		EncryptedPrivateKey string `yaml:"encrypted_private_key"`
		Nonce               string `yaml:"nonce"`
		KeyPhrase           string `yaml:"key_phrase"`
	}
	err = yaml.Unmarshal(yamlFile, &config)
	assert.NoError(t, err)
	t.Logf("Parsed config - EncryptedPrivateKey: %s, Nonce: %s",
		config.EncryptedPrivateKey, config.Nonce)

	ciphertext, err := hex.DecodeString(config.EncryptedPrivateKey)
	assert.NoError(t, err)
	t.Logf("Decoded ciphertext length: %d", len(ciphertext))

	nonce, err := hex.DecodeString(config.Nonce)
	assert.NoError(t, err)
	t.Logf("Decoded nonce length: %d", len(nonce))

	hasher := sha256.New()
	hasher.Write([]byte(config.KeyPhrase))
	key := hasher.Sum(nil)
	t.Logf("Key length: %d", len(key))

	block, err := aes.NewCipher(key)
	assert.NoError(t, err)
	t.Logf("Created AES cipher")

	gcm, err := cipher.NewGCM(block)
	assert.NoError(t, err)
	t.Logf("Created GCM mode")

	privateKeyBytes, err := gcm.Open(nil, nonce, ciphertext, nil)
	assert.NoError(t, err)
	t.Logf("Decrypted private key length: %d", len(privateKeyBytes))

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	assert.NoError(t, err)
	t.Logf("Converted to ECDSA private key")

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	assert.True(t, ok)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	t.Log("=== Final Results ===")
	t.Logf("Decrypted Private Key: %x", crypto.FromECDSA(privateKey))
	t.Logf("Generated Address: %s", address.Hex())
}

func TestLocalOfflineSignTx(t *testing.T) {
	// 1. Read and decrypt private key from config
	yamlFile, err := ioutil.ReadFile("../../config.yaml")
	assert.NoError(t, err)
	t.Log("Read config file successfully")

	var config struct {
		EncryptedPrivateKey string `yaml:"encrypted_private_key"`
		Nonce               string `yaml:"nonce"`
		KeyPhrase           string `yaml:"key_phrase"`
	}
	err = yaml.Unmarshal(yamlFile, &config)
	assert.NoError(t, err)

	// 2. Decrypt private key
	ciphertext, err := hex.DecodeString(config.EncryptedPrivateKey)
	assert.NoError(t, err)
	nonce, err := hex.DecodeString(config.Nonce)
	assert.NoError(t, err)

	hasher := sha256.New()
	hasher.Write([]byte(config.KeyPhrase))
	key := hasher.Sum(nil)

	block, err := aes.NewCipher(key)
	assert.NoError(t, err)
	gcm, err := cipher.NewGCM(block)
	assert.NoError(t, err)

	privateKeyBytes, err := gcm.Open(nil, nonce, ciphertext, nil)
	assert.NoError(t, err)
	t.Logf("Successfully decrypted private key: %x", privateKeyBytes)

	// 3. Set transaction parameters
	toAddress := common.HexToAddress("0x67936bb11f8fd1d25da1f94e0aa51039409a7c97")
	amount := big.NewInt(1000000000000)            // 0.000001 ETH
	gasLimit := uint64(21000)                      // Standard transfer gas limit
	maxPriorityFeePerGas := big.NewInt(2600000000) // 2.6 Gwei
	maxFeePerGas := big.NewInt(2900000000)         // 2.9 Gwei
	chainID := big.NewInt(137)                     // Polygon mainnet

	// 4. Create transaction object
	dFeeTx := &types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     uint64(0), // Set nonce according to actual situation
		GasTipCap: maxPriorityFeePerGas,
		GasFeeCap: maxFeePerGas,
		Gas:       gasLimit,
		To:        &toAddress,
		Value:     amount,
		Data:      nil,
	}

	// 5. Call offline signing method
	privateKeyHex := hex.EncodeToString(privateKeyBytes)
	println("========", privateKeyHex)
	txHex, txHash, err := OfflineSignTx(dFeeTx, privateKeyHex, chainID)
	assert.NoError(t, err)

	// 6. Print results
	t.Log("=== Transaction Details ===")
	t.Logf("To Address: %s", toAddress.Hex())
	t.Logf("Amount: %s wei", amount.String())
	t.Logf("Chain ID: %s", chainID.String())
	t.Logf("Transaction Hex: %s", txHex)
	t.Logf("Transaction Hash: %s", txHash)
}

func OfflineSignTx(txData *types.DynamicFeeTx, privateKey string, chainId *big.Int) (string, string, error) {
	privateKeyEcdsa, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", "", err
	}

	tx := types.NewTx(txData)

	signer := types.LatestSignerForChainID(chainId)

	signedTx, err := types.SignTx(tx, signer, privateKeyEcdsa)

	if err != nil {
		return "", "", err
	}

	signedTxData, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		return "", "", err
	}

	return "0x" + hex.EncodeToString(signedTxData), signedTx.Hash().String(), nil
}
