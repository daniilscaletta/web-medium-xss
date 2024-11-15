package utils

import (
	"crypto/sha256"
	"math/big"
)

func CheckHashCashChallenge(challenge, solution string) bool {

	hashInput := challenge + solution
	hash := sha256.New()
	hash.Write([]byte(hashInput))
	hashResult := hash.Sum(nil)

	hashInt := new(big.Int)
	hashInt.SetBytes(hashResult)

	target := new(big.Int).Lsh(big.NewInt(1), 240) // target = 1 << 240

	return hashInt.Cmp(target) < 0
}
