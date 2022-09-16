package app

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
)

const (
	// StandardScryptN is the N parameter of Scrypt encryption algorithm, using 256MB
	// memory and taking approximately 1s CPU time on a modern processor.
	scryptN = 1 << 18

	// StandardScryptP is the P parameter of Scrypt encryption algorithm, using 256MB
	// memory and taking approximately 1s CPU time on a modern processor.
	scryptP = 1
)

func newKeyFromECDSA(privateKeyECDSA *ecdsa.PrivateKey) *keystore.Key {
	id := uuid.New()
	key := &keystore.Key{
		Id:         id,
		Address:    crypto.PubkeyToAddress(privateKeyECDSA.PublicKey),
		PrivateKey: privateKeyECDSA,
	}
	return key
}

func newKey() (*keystore.Key, error) {
	privateKeyECDSA, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	return newKeyFromECDSA(privateKeyECDSA), nil
}

func (a *WalletApplication) loadJsonPrivateKey(pKeyFilePath, password string) (string, error) {

	content, err := os.ReadFile(pKeyFilePath)
	if err != nil {
		return "", errors.New("unable to access private key from file system")
	}

	key, err := keystore.DecryptKey(content, password)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(key.PrivateKey.D.Bytes()), nil
}

func (a *WalletApplication) generateKeyStore(privateKeyStr, password string) ([]byte, error) {
	hexBytes, err := hex.DecodeString(privateKeyStr)
	if err != nil {
		return nil, err
	}

	key := crypto.ToECDSAUnsafe(hexBytes)

	privateKey := &keystore.Key{
		Id:         uuid.New(),
		Address:    crypto.PubkeyToAddress(key.PublicKey),
		PrivateKey: key,
	}

	return keystore.EncryptKey(privateKey, password, scryptN, scryptP)
}

func (a *WalletApplication) generateNewKeyStore(password string) ([]byte, error) {
	privateKey, err := newKey()
	if err != nil {
		return nil, err
	}

	return keystore.EncryptKey(privateKey, password, scryptN, scryptP)
}

func prettyPrint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}
