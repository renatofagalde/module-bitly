package bitly

import (
	"errors"
	"math/big"
)

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

var (
	base       = big.NewInt(62)
	indexTable = make(map[byte]int)
)

func init() {
	for i := 0; i < len(alphabet); i++ {
		indexTable[alphabet[i]] = i
	}
}

func E(n uint64) string {
	if n == 0 {
		return string(alphabet[0])
	}

	var encoded []byte
	value := n

	for value > 0 {
		remainder := value % 62
		value = value / 62
		encoded = append([]byte{alphabet[remainder]}, encoded...)
	}

	return string(encoded)
}

func D(s string) (uint64, error) {
	var result uint64

	for i := 0; i < len(s); i++ {
		ch := s[i]
		idx, ok := indexTable[ch]
		if !ok {
			return 0, errors.New("bitly: caractere inválido em base62")
		}

		// checa overflow
		if result > (^uint64(0))/62 {
			return 0, errors.New("bitly: overflow ao decodificar uint64")
		}

		result = result*62 + uint64(idx)
	}

	return result, nil
}

func EncodeBytes(data []byte) string {
	if len(data) == 0 {
		return ""
	}

	n := new(big.Int).SetBytes(data)
	if n.Sign() == 0 {
		return string(alphabet[0])
	}

	var encoded []byte
	zero := big.NewInt(0)
	mod := new(big.Int)
	for n.Cmp(zero) > 0 {
		n.DivMod(n, base, mod)
		encoded = append([]byte{alphabet[mod.Int64()]}, encoded...)
	}

	return string(encoded)
}

func DecodeBytes(s string) ([]byte, error) {
	n := big.NewInt(0)

	for i := 0; i < len(s); i++ {
		ch := s[i]
		idx, ok := indexTable[ch]
		if !ok {
			return nil, errors.New("bitly: caractere inválido em base62")
		}
		n.Mul(n, base)
		n.Add(n, big.NewInt(int64(idx)))
	}

	return n.Bytes(), nil
}
