package crypt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"

	"golang.org/x/crypto/sha3"
)

// GenerateRSAKeyPair generates a new pair of RSA keys (private and public).
func GenerateRSAKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	keySize := 2048
	privK, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return nil, nil, err
	}
	privK.Precompute()
	pubK := privK.PublicKey
	return privK, &pubK, nil
}

// EncodeKeys prepares a pair of RSA keys to be sent to the server. The returned strings are:
//
// 1. The private key, after JSON encoding, compression, AES encrypting with dataK and Base64 encoding;
//
// 2. The public key after JSON encoding, compression and Base64 encoding.
//
// The dataK MUST be either 16, 24 or 32 bytes.
func EncodeKeys(privK *rsa.PrivateKey, pubK *rsa.PublicKey, dataK []byte) (*string, *string, error) {
	privJSON, err := json.Marshal(&privK)
	if err != nil {
		return nil, nil, err
	}
	pubJSON, err := json.Marshal(&pubK)
	if err != nil {
		return nil, nil, err
	}
	priv := Encode64(EncryptAES(Compress(privJSON), dataK))
	pub := Encode64(Compress(pubJSON))
	return &priv, &pub, nil
}

// EncryptWithPubRSAAndEncode64 encrypts the given data with the given RSA public key and encodes the result in Base64.
// Mind that the size of the data must be lower than the maximum size allowed by the RSA key.
func EncryptWithPubRSAAndEncode64(data []byte, pubK *rsa.PublicKey) (string, error) {
	encrypted_data, err := EncryptWithRSAPublicKey(data, pubK)
	if err != nil {
		return "", err
	}
	return Encode64(encrypted_data), nil
}

// DecodePrivKey applies to the given encoded private key the inverse transformations to those done by EncodeKeys.
// See its documentation.
func DecodePrivKey(ePrivK string, dataK []byte) (privK *rsa.PrivateKey, err error) {
	decPrivK, err := Decode64(ePrivK)
	if err != nil {
		return
	}

	privJSON, err := Decompress(DecryptAES(decPrivK, dataK))
	if err != nil {
		return
	}

	var tempPrivK rsa.PrivateKey
	err = json.Unmarshal(privJSON, &tempPrivK)
	if err != nil {
		return
	}
	privK = &tempPrivK
	return
}

// DecodePubKey applies to the given encoded public key the inverse transformations to those done by EncodeKeys.
// See its documentation.
func DecodePubKey(ePubK string) (pubK *rsa.PublicKey, err error) {
	decPubK, err := Decode64(ePubK)
	if err != nil {
		return
	}

	pubJSON, err := Decompress(decPubK)
	if err != nil {
		return
	}

	var tempPubK rsa.PublicKey
	err = json.Unmarshal(pubJSON, &tempPubK)
	if err != nil {
		return
	}
	pubK = &tempPubK
	return
}

// DecodeKeys applies the inverse transformations of EncodeKeys() to the given encoded key pair.
// See its documentation.
func DecodeKeys(ePrivK, ePubK string, dataK []byte) (privK *rsa.PrivateKey, pubK *rsa.PublicKey, err error) {
	err = nil
	privK, err = DecodePrivKey(ePrivK, dataK)
	if err == nil {
		pubK, err = DecodePubKey(ePubK)
	}
	return
}

// EncryptWithRSAPublicKey encrypts data with the given RSA public key
func EncryptWithRSAPublicKey(msg []byte, pub *rsa.PublicKey) ([]byte, error) {
	hash := sha3.New512()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

// DecryptWithRSAPrivateKey decrypts data with the given RSA private key
func DecryptWithRSAPrivateKey(ciphertext []byte, priv *rsa.PrivateKey) ([]byte, error) {
	hash := sha3.New512()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

// SignRSA returns a signature made by combining the message and the signers private key
// With the VerifySignature function, the signature can be checked.
func SignRSA(msg string, priv *rsa.PrivateKey) (signature []byte, err error) {
	// hs := Hash256(msg)
	signature, err = rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, []byte(msg))
	return
}

// VerifyRSASignature checks if a message is signed by a given Public Key
func VerifyRSASignature(msg string, sig []byte, pk *rsa.PublicKey) error {
	// hs := Hash256(msg)
	return rsa.VerifyPKCS1v15(pk, crypto.SHA256, []byte(msg), sig)
}
