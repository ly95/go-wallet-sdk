package mobile

import (
	"encoding/hex"

	"github.com/okx/go-wallet-sdk/coins/aptos"
	"github.com/okx/go-wallet-sdk/crypto/ed25519"
)

func APTOS_Swap(
	seedHex string,
	from string,
	sequenceNumber int64,
	maxGasAmount int64,
	gasUnitPrice int64,
	expirationTimestampSecs int64,
	chainId int8,
	payload string,
	abi string,
) string {
	p, _ := aptos.PayloadFromJsonAndAbi(payload, abi)
	rawTxn := aptos.MakeRawTransaction(
		from,
		uint64(sequenceNumber),
		uint64(maxGasAmount),
		uint64(gasUnitPrice),
		uint64(expirationTimestampSecs),
		uint8(chainId),
		p,
	)

	message, err := rawTxn.GetSigningMessage()
	if err != nil {
		return ""
	}
	signature, err := ed25519.Sign(seedHex, message)
	if err != nil {
		return ""
	}

	return hex.EncodeToString(signature)
}
