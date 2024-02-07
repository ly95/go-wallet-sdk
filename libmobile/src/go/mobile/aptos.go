package mobile

import (
	"github.com/okx/go-wallet-sdk/coins/aptos"
	"github.com/okx/go-wallet-sdk/coins/aptos/types"
	"github.com/okx/go-wallet-sdk/crypto/ed25519"
)

func APTOS_NewAddress(seedHex string, shortEnable bool) (string, error) {
	return aptos.NewAddress(seedHex, shortEnable)
}

func APTOS_GetAddressByPubKey(pubKeyHex string, shortEnable bool) (string, error) {
	return aptos.GetAddressByPubKey(pubKeyHex, shortEnable)
}

func APTOS_ValidateAddress(address string, shortEnable bool) bool {
	return aptos.ValidateAddress(address, shortEnable)
}

func APTOS_MakeRawTransaction(
	from string,
	sequenceNumber uint64,
	maxGasAmount uint64,
	gasUnitPrice uint64,
	expirationTimestampSecs uint64,
	chainId uint8,
	payload types.TransactionPayload,
) *types.RawTransaction {
	return aptos.MakeRawTransaction(
		from,
		sequenceNumber,
		maxGasAmount,
		gasUnitPrice,
		expirationTimestampSecs,
		chainId,
		payload,
	)
}

func APTOS_BuildSignedTransaction(
	from string,
	sequenceNumber uint64,
	maxGasAmount uint64,
	gasUnitPrice uint64,
	expirationTimestampSecs uint64,
	chainId uint8,
	payload types.TransactionPayload,
	seedHex string,
) (string, error) {
	return aptos.BuildSignedTransaction(
		from,
		sequenceNumber,
		maxGasAmount,
		gasUnitPrice,
		expirationTimestampSecs,
		chainId,
		payload,
		seedHex,
	)
}

func APTOS_SignRawTransaction(rawTxn *types.RawTransaction, seedHex string) (string, error) {
	return aptos.SignRawTransaction(rawTxn, seedHex)
}

// 只取签名
func APTOS_ED25519Sign(rawTxn *types.RawTransaction, seedHex string) ([]byte, error) {
	message, err := rawTxn.GetSigningMessage()
	if err != nil {
		return nil, err
	}
	signature, err := ed25519.Sign(seedHex, message)
	if err != nil {
		return nil, err
	}
	return signature, nil
}
