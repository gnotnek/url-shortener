package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

func sha256Hash(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Printf("Error encoding: %v\n", err)
		os.Exit(1)
	}

	return string(encoded)
}

func GenerateShortUrl(originalUrl string, userId string) string {
	hash := sha256Hash(originalUrl + userId)
	generatedNum := new(big.Int).SetBytes(hash).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNum)))
	return finalString[:8]
}
