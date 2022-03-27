package main

import (
	"crypto/rand"
	"math/big"
)

// 生成安全随机数
func randInt(min, max int64) int64 {
	bigInt := big.NewInt(max - min)
	n, _ := rand.Int(rand.Reader, bigInt)
	return min + n.Int64()
}
