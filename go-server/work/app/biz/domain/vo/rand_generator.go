package vo

import (
	"math/rand"
	"strings"
	"time"
	"unsafe"
)

type RandGeneratorType int

const (
	DefaultRandGenerator RandGeneratorType = iota
	RandPasswordGenerator
	ShortNumberGenerator
)

const (
	numbers      = "0123456789"
	lowerLetters = "abcdefghijklmnopqrstuvwxyz"
	upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letters      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphas       = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	punctuation  = "!\"#$%&'()*+,-./:;<=>?@[]^_{|}~"
)

var (
	defaultGenerator    = RandGenerator{characterSet: []string{alphas}, length: []int{8}}
	passwordGenerator   = RandGenerator{characterSet: []string{numbers, lowerLetters, upperLetters, punctuation}, length: []int{2, 4, 4, 2}}
	shortNumberGenerate = RandGenerator{characterSet: []string{numbers}, length: []int{4}}
)

type RandGenerator struct {
	characterSet []string
	length       []int
}

func NewRandGenerator(t RandGeneratorType) RandGenerator {
	switch t {
	case RandPasswordGenerator:
		return passwordGenerator
	case ShortNumberGenerator:
		return shortNumberGenerate
	}
	return defaultGenerator
}

func (r RandGenerator) Next() string {
	parts := make([]string, len(r.characterSet))
	for i, s := range r.characterSet {
		parts[i] = getRandString(s, r.length[i])
	}
	randBytes := []byte(strings.Join(parts, ""))
	rand.Shuffle(len(randBytes), func(i, j int) {
		randBytes[i], randBytes[j] = randBytes[j], randBytes[i]
	})
	return string(randBytes)
}

func (r RandGenerator) NextN(n int) string {
	return getRandString(alphas, n)
}

func getRandString(source string, length int) string {
	var (
		letterIdxBits       = 6                    // 6 bits to represent a letter index
		letterIdxMask int64 = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax        = 63 / letterIdxBits   // # of letter indices fitting in 63 bits

		src = rand.NewSource(time.Now().UnixNano())
		b   = make([]byte, length)
	)

	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(source) {
			b[i] = source[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
