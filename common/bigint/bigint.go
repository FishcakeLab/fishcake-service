package bigint

import (
	"math/big"
)

func StringToInt(value string) int {
	if value == "" {
		return 0
	}
	return int(StringToBigInt(value).Int64())
}

func StringToBigInt(value string) *big.Int {
	intValue, success := big.NewInt(0).SetString(value, 0)
	if !success {
		return nil
	}
	return intValue
}
