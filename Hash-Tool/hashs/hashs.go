package hashs

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"strings"
	"sync"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

var (
	hashFuncsInit sync.Once
	hashFuncs map[string]func () (hash.Hash, error)
)

func initHashFuncs() {
	hashFuncs = map[string]func() (hash.Hash, error) {
		// MD Family
		"md4":         func() (hash.Hash, error) { return md4.New(), nil },
		"md5":         func() (hash.Hash, error) { return md5.New(), nil },
	
		// SHA-1
		"sha1": func() (hash.Hash, error) { return sha1.New(), nil },
	
		// SHA-2 Family
		"sha224":      func() (hash.Hash, error) { return sha256.New224(), nil },
		"sha256":      func() (hash.Hash, error) { return sha256.New(), nil },
		"sha384":      func() (hash.Hash, error) { return sha512.New384(), nil },
		"sha512":      func() (hash.Hash, error) { return sha512.New(), nil },
		"sha512_224":  func() (hash.Hash, error) { return sha512.New512_224(), nil },
		"sha512_256":  func() (hash.Hash, error) { return sha512.New512_256(), nil },
	
		// SHA-3 Family
		"sha3_224":    func() (hash.Hash, error) { return sha3.New224(), nil },
		"sha3_256":    func() (hash.Hash, error) { return sha3.New256(), nil },
		"sha3_384":    func() (hash.Hash, error) { return sha3.New384(), nil },
		"sha3_512":    func() (hash.Hash, error) { return sha3.New512(), nil },
	
		// RIPEMD
		"ripemd160":   func() (hash.Hash, error) { return ripemd160.New(), nil },
	
		// BLAKE2 Family
		"blake2s_256": func() (hash.Hash, error) { return blake2s.New256(nil) },
		"blake2b_256": func() (hash.Hash, error) { return blake2b.New256(nil) },
		"blake2b_384": func() (hash.Hash, error) { return blake2b.New384(nil) },
		"blake2b_512": func() (hash.Hash, error) { return blake2b.New512(nil) },
	}
}

func HashData(message string, hashName string) (string, error) {
	hashFuncsInit.Do(initHashFuncs)

	hashFunc, exists := hashFuncs[strings.ToLower(hashName)]
	if !exists {
		return "", fmt.Errorf("unsupported hash algorithm: %s", hashName)
	}

	hasher, _ := hashFunc()
	hasher.Write([]byte(message))

	return hex.EncodeToString(hasher.Sum(nil)), nil
}