package crypt

import (
	"bytes"
	"compress/zlib"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/sha3"
)

const SALT_LENGTH int = 16
const PASS_HASH_LENGTH uint32 = 32

const ARGON_TIME uint32 = 3    // Number of passses over the memory
const ARGON_MEMORY uint32 = 32 // In MB
const ARGON_CPUS uint8 = 4     // CPU threads

// PanicIfErr panics if err is not nil
func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

// EncryptAES encrypts the given data with AES, using a 128-bit initialization vector.
// Said IV will be the first 128 bits of the encrypted message.
// The key MUST be either 16, 24 or 32 bytes.
func EncryptAES(data, key []byte) (out []byte) {
	out = make([]byte, len(data)+16)
	rand.Read(out[:16])
	blk, err := aes.NewCipher(key)
	PanicIfErr(err)
	ctr := cipher.NewCTR(blk, out[:16])
	ctr.XORKeyStream(out[16:], data)
	return
}

// DecryptAES decrypts the given data with AES, using a 128-bit initialization vector.
// Said IV should be the first 128 bits of the encrypted message.
// The key MUST be either 16, 24 or 32 bytes.
func DecryptAES(data, key []byte) (out []byte) {
	out = make([]byte, len(data)-16)
	blk, err := aes.NewCipher(key)
	PanicIfErr(err)
	ctr := cipher.NewCTR(blk, data[:16])
	ctr.XORKeyStream(out, data[16:])
	return
}

// Compress compresses the given data with the zip algorithm.
func Compress(data []byte) []byte {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(data)
	w.Close()
	return b.Bytes()
}

// Decompress decompresses the given data with the zip algorithm.
func Decompress(data []byte) ([]byte, error) {
	var b bytes.Buffer

	r, err := zlib.NewReader(bytes.NewReader(data))

	if err != nil {
		return nil, err
	}

	io.Copy(&b, r)
	r.Close()
	return b.Bytes(), nil
}

// Encode64 encodes the given data in Base64.
func Encode64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Decode64 decodes the given data from Base64.
func Decode64(s string) ([]byte, error) {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Hash256 creates a SHA-3 256-bit hash from the given message.
func Hash256(msg string) []byte {
	h := sha3.New256()
	h.Write([]byte(msg))
	return h.Sum(nil)
}

// Generates a random sequence of SALT_SIZE bytes
func GenerateSalt() []byte {
	salt := make([]byte, SALT_LENGTH)
	rand.Read(salt)
	return salt
}

// Applies the Argon2 algorithm to the given password and salt
func PBKDF(password, salt []byte) []byte {
	return argon2.Key(password, salt, ARGON_TIME, ARGON_MEMORY*1024, ARGON_CPUS, PASS_HASH_LENGTH)
}
