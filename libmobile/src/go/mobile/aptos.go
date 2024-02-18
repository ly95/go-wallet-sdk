package mobile

import (
	"encoding/hex"

	"github.com/okx/go-wallet-sdk/coins/aptos"
	"github.com/okx/go-wallet-sdk/crypto/ed25519"
	"github.com/omnibtc/go-aptos-liquidswap/liquidswap"
	"github.com/shopspring/decimal"
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

func APTOS_AmountMinOut(val string, slippage string) string {
	_val, err := decimal.NewFromString(val)
	if err != nil {
		return ""
	}
	_slippage, err := decimal.NewFromString(slippage)
	if err != nil {
		return ""
	}
	r := liquidswap.AmountMinOut(_val.BigInt(), _slippage)
	if r == nil {
		return ""
	}
	return r.String()
}

func APTOS_AmountMaxIn(val string, slippage string) string {
	_val, err := decimal.NewFromString(val)
	if err != nil {
		return ""
	}
	_slippage, err := decimal.NewFromString(slippage)
	if err != nil {
		return ""
	}
	r := liquidswap.AmountMaxIn(_val.BigInt(), _slippage)
	if r == nil {
		return ""
	}
	return r.String()
}

func APTOS_GetAmountIn(
	fromCoinDecimals int,
	fromCoinName string,
	fromCoinSymbol string,
	toCoinDecimals int,
	toCoinName string,
	toCoinSymbol string,
	amountOut string,
	poolCoinXReserve string,
	poolCoinYReserve string,
	poolCurveType int,
) string {
	fromCoin := liquidswap.Coin{
		Decimals: fromCoinDecimals,
		Name:     fromCoinName,
		Symbol:   fromCoinSymbol,
	}
	toCoin := liquidswap.Coin{
		Decimals: toCoinDecimals,
		Name:     toCoinName,
		Symbol:   toCoinSymbol,
	}
	_amountOut, err := decimal.NewFromString(amountOut)
	if err != nil {
		return ""
	}
	_poolCoinXReserve, err := decimal.NewFromString(poolCoinXReserve)
	if err != nil {
		return ""
	}
	_poolCoinYReserve, err := decimal.NewFromString(poolCoinYReserve)
	if err != nil {
		return ""
	}
	pool := liquidswap.PoolResource{
		CoinXReserve: _poolCoinXReserve.BigInt(),
		CoinYReserve: _poolCoinYReserve.BigInt(),
		CurveType:    poolCurveType,
	}
	r := liquidswap.GetAmountIn(fromCoin, toCoin, _amountOut.BigInt(), pool)
	if r == nil {
		return ""
	}
	return r.String()
}

func APTOS_GetAmountOut(
	fromCoinDecimals int,
	fromCoinName string,
	fromCoinSymbol string,
	toCoinDecimals int,
	toCoinName string,
	toCoinSymbol string,
	amountIn string,
	poolCoinXReserve string,
	poolCoinYReserve string,
	poolCurveType int,
) string {
	fromCoin := liquidswap.Coin{
		Decimals: fromCoinDecimals,
		Name:     fromCoinName,
		Symbol:   fromCoinSymbol,
	}
	toCoin := liquidswap.Coin{
		Decimals: toCoinDecimals,
		Name:     toCoinName,
		Symbol:   toCoinSymbol,
	}
	_amountIn, err := decimal.NewFromString(amountIn)
	if err != nil {
		return ""
	}
	_poolCoinXReserve, err := decimal.NewFromString(poolCoinXReserve)
	if err != nil {
		return ""
	}
	_poolCoinYReserve, err := decimal.NewFromString(poolCoinYReserve)
	if err != nil {
		return ""
	}
	pool := liquidswap.PoolResource{
		CoinXReserve: _poolCoinXReserve.BigInt(),
		CoinYReserve: _poolCoinYReserve.BigInt(),
		CurveType:    poolCurveType,
	}
	r := liquidswap.GetAmountOut(fromCoin, toCoin, _amountIn.BigInt(), pool)
	if r == nil {
		return ""
	}
	return r.String()
}
