package reward_service

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/rlp"

	"github.com/FishcakeLab/fishcake-service/rpc/account"
	"github.com/FishcakeLab/fishcake-service/service/rpc_service"

	"github.com/FishcakeLab/fishcake-service/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// TokenType represents the type of token for transactions
type TokenType int

const (
	Native TokenType = iota // Native token
	USDT                    // USDT token
	FCC                     // FCC token
)

type RewardService interface {
	DecryptPrivateKey() ([]byte, error)
	CreateOfflineTransaction(tokenType TokenType, privateKeyBytes []byte, toAddress string, amount *big.Int) (string, string, error)
}

//	var config struct {
//		EncryptedPrivateKey string `yaml:"encrypted_private_key"`
//		Nonce               string `yaml:"nonce"`
//		KeyPhrase           string `yaml:"key_phrase"`
//	}
type rewardService struct {
	cfg *config.Config
}

func NewRewardService(configPath string) RewardService {
	if configPath == "" {
		configPath = "config.yaml" // Default path
	}

	cfg, err := config.New(configPath)
	if err != nil {
		fmt.Printf("Error parsing config file: %v\n", err)
		return &rewardService{}
	}

	return &rewardService{
		cfg: cfg,
	}
}

// DecryptPrivateKey reads parameters from config file and decrypts the private key
func (s *rewardService) DecryptPrivateKey() ([]byte, error) {
	if s.cfg == nil {
		return nil, fmt.Errorf("config is not initialized")
	}

	// Get parameters from config
	encryptedKey := s.cfg.EncryptedPrivateKey
	nonce := s.cfg.Nonce
	keyPhrase := s.cfg.KeyPhrase

	// 1. Decode encrypted data and nonce
	ciphertext, err := hex.DecodeString(encryptedKey)
	if err != nil {
		return nil, err
	}
	nonceBytes, err := hex.DecodeString(nonce)
	if err != nil {
		return nil, err
	}

	// 2. Generate key
	hasher := sha256.New()
	hasher.Write([]byte(keyPhrase))
	key := hasher.Sum(nil)

	// 3. Create decrypter
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// 4. Decrypt private key
	privateKeyBytes, err := gcm.Open(nil, nonceBytes, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return privateKeyBytes, nil
}

// CreateOfflineTransaction creates an offline transaction based on token type
func (s *rewardService) CreateOfflineTransaction(tokenType TokenType, privateKeyBytes []byte, toAddress string, amount *big.Int) (string, string, error) {
	switch tokenType {
	case Native:
		return s.createNativeTransaction(privateKeyBytes, toAddress, amount)
	case USDT:
		return s.createUSDTTransaction(privateKeyBytes, toAddress, amount)
	case FCC:
		return s.createFCCTransaction(privateKeyBytes, toAddress, amount)
	default:
		return "", "", fmt.Errorf("unsupported token type: %v", tokenType)
	}
}

// createNativeTransaction creates a native token transaction
func (s *rewardService) createNativeTransaction(privateKeyBytes []byte, toAddress string, amount *big.Int) (string, string, error) {
	to := common.HexToAddress(toAddress)
	gasLimit := uint64(21000) // Standard transfer gas limit

	// 1. Create RPC service client
	rpcService := rpc_service.NewRpcService(s.cfg.RpcUrl)

	// 2. Prepare request parameters
	ctx := context.Background()

	req := &account.AccountRequest{
		Chain:   "Polygon", // Chain name
		Address: toAddress, // Address to convert
		Network: "mainnet", // Target network
	}

	// 3. Call account service
	getAccount, err := rpcService.GetAccount(ctx, req)
	if err != nil {
		return "", "", err
	}

	nonce, _ := strconv.ParseUint(getAccount.Sequence, 10, 64)

	reqFee := &account.FeeRequest{
		Chain:   "Polygon",
		Network: "mainnet",
	}
	fee, err := rpcService.GetFee(ctx, reqFee)
	if err != nil {
		return "", "", err
	}

	maxPriorityFeePerGas := new(big.Int).SetInt64(2600000000) // 2.6 Gwei
	parts := strings.Split(fee.FastFee, "|")

	// Extract first part and convert to big.Int
	firstNumberStr := parts[0]
	bigIntValue := new(big.Int)
	_, _ = bigIntValue.SetString(firstNumberStr, 10)

	chainID := new(big.Int).SetInt64(80001) // Mumbai testnet

	dFeeTx := &types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: maxPriorityFeePerGas,
		GasFeeCap: bigIntValue,
		Gas:       gasLimit,
		To:        &to,
		Value:     amount,
		Data:      nil,
	}

	privateKeyHex := hex.EncodeToString(privateKeyBytes)
	return OfflineSignTx(dFeeTx, privateKeyHex, chainID)
}

// createUSDTTransaction creates a USDT token transaction
func (s *rewardService) createUSDTTransaction(privateKeyBytes []byte, toAddress string, amount *big.Int) (string, string, error) {
	// 1. Create RPC service client
	rpcService := rpc_service.NewRpcService(s.cfg.RpcUrl)

	// 2. Prepare request parameters
	ctx := context.Background()

	privateKeyHex := hex.EncodeToString(privateKeyBytes)
	privateKey, _ := crypto.ToECDSA(privateKeyBytes)

	// 3. Generate address from public key
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	req := &account.AccountRequest{
		Chain:   "Polygon",     // Chain name
		Address: address.Hex(), // Address to convert
		Network: "mainnet",     // Target network
	}

	getAccount, err := rpcService.GetAccount(ctx, req)
	if err != nil {
		return "", "", err
	}

	nonce, _ := strconv.ParseUint(getAccount.Sequence, 10, 64)

	reqFee := &account.FeeRequest{
		Chain:   "Polygon",
		Network: "mainnet",
	}
	fee, err := rpcService.GetFee(ctx, reqFee)
	if err != nil {
		return "", "", err
	}
	maxPriorityFeePerGas := new(big.Int).SetInt64(30000000000) // 30 Gwei

	parts := strings.Split(fee.FastFee, "|")

	// Extract first part and convert to big.Int
	firstNumberStr := parts[0]
	bigIntValue := new(big.Int)
	_, _ = bigIntValue.SetString(firstNumberStr, 10)

	contractAddr := s.cfg.USDT // USDT contract address
	to := common.HexToAddress(contractAddr)
	gasLimit := uint64(150000)            // ERC20 transfer gas limit
	chainID := new(big.Int).SetInt64(137) // Polygon mainnet

	// Create ERC20 transfer method data
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := crypto.Keccak256(transferFnSignature)
	methodID := hash[:4]

	paddedAddress := common.LeftPadBytes(common.HexToAddress(toAddress).Bytes(), 32)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	bigIntValue = new(big.Int).Mul(bigIntValue, big.NewInt(2))

	dFeeTx := &types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: maxPriorityFeePerGas,
		GasFeeCap: bigIntValue,
		Gas:       gasLimit,
		To:        &to,
		Value:     new(big.Int), // Value is 0 for ERC20 transfer
		Data:      data,
	}

	return OfflineSignTx(dFeeTx, privateKeyHex, chainID)
}

func (s *rewardService) createFCCTransaction(privateKeyBytes []byte, toAddress string, amount *big.Int) (string, string, error) {

	// 1. Create RPC service client
	rpcService := rpc_service.NewRpcService(s.cfg.RpcUrl)

	// 2. Prepare request parameters
	ctx := context.Background()

	privateKeyHex := hex.EncodeToString(privateKeyBytes)
	privateKey, _ := crypto.ToECDSA(privateKeyBytes)

	// 3. Generate address from public key
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	req := &account.AccountRequest{
		Chain:   "Polygon",     // Chain name
		Address: address.Hex(), // Address to convert
		Network: "mainnet",     // Target network
	}

	getAccount, err := rpcService.GetAccount(ctx, req)
	if err != nil || getAccount == nil {
		return "", "", err
	}

	nonce, _ := strconv.ParseUint(getAccount.Sequence, 10, 64)

	reqFee := &account.FeeRequest{
		Chain:   "Polygon",
		Network: "mainnet",
	}
	fee, err := rpcService.GetFee(ctx, reqFee)
	if err != nil {
		return "", "", err
	}
	maxPriorityFeePerGas := new(big.Int).SetInt64(30000000000) // 30 Gwei

	parts := strings.Split(fee.FastFee, "|")

	// Extract first part and convert to big.Int
	firstNumberStr := parts[0]
	bigIntValue := new(big.Int)
	_, _ = bigIntValue.SetString(firstNumberStr, 10)

	contractAddr := s.cfg.FCC // FCC contract address
	to := common.HexToAddress(contractAddr)
	gasLimit := uint64(150000)            // ERC20 transfer gas limit
	chainID := new(big.Int).SetInt64(137) // Polygon mainnet

	// Create ERC20 transfer method data
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := crypto.Keccak256(transferFnSignature)
	methodID := hash[:4]

	paddedAddress := common.LeftPadBytes(common.HexToAddress(toAddress).Bytes(), 32)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	bigIntValue = new(big.Int).Mul(bigIntValue, big.NewInt(2))

	dFeeTx := &types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: maxPriorityFeePerGas,
		GasFeeCap: bigIntValue,
		Gas:       gasLimit,
		To:        &to,
		Value:     new(big.Int), // Value is 0 for ERC20 transfer
		Data:      data,
	}

	return OfflineSignTx(dFeeTx, privateKeyHex, chainID)
}

// OfflineSignTx signs a transaction offline
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

	return "0x" + hex.EncodeToString(signedTxData)[4:], signedTx.Hash().String(), nil
}
