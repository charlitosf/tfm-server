package jwt

import (
	"charlitosf/tfm-server/internal/crypt"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

const TOKEN_DURATION_MINUTES time.Duration = 15

const ISSUER string = "NOZAMA"

var jwtSigningKey *ecdsa.PrivateKey
var jwtDecodingKey *ecdsa.PublicKey
var jwtServerInstanceId string
var nextJWTId int = 0
var issuedTokens map[int]struct{} = make(map[int]struct{}) // Set of claim identifiers for JWT

var JwtExpireTime time.Duration = TOKEN_DURATION_MINUTES * time.Minute

func init() {
	var err error
	jwtSigningKey, jwtDecodingKey, err = crypt.GenerateECKeyPair()
	crypt.PanicIfErr(err)

	var instanceID []byte = make([]byte, 16)
	rand.Read(instanceID)
	jwtServerInstanceId = crypt.Encode64(instanceID)
}

type JWTCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateJWT generates a JWT token and signs it with jwtSigningKey (private key).
// The token's payload contains just the "username" and will expire after an interval of jwtExpireTime.
func GenerateJWT(username string) (string, error) {

	claims := JWTCustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(JwtExpireTime).Unix(),
			Issuer:    ISSUER,
			Id:        fmt.Sprintf("%v-%v", jwtServerInstanceId, nextJWTId),
		},
	}
	issuedTokens[nextJWTId] = struct{}{} // Add id to the set of issued tokens
	nextJWTId++

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, err := token.SignedString(jwtSigningKey)
	return signedToken, err
}

// VerifyJWT parses the given JWT token and returns the username contained in it.
// If there is a parse error or the token is not valid, a suitable error is returned.
func VerifyJWT(tokenString string) (username string, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, parsingFunc)
	if err != nil {
		return "", err
	}

	if token.Valid {
		if claims, ok := token.Claims.(*JWTCustomClaims); ok {
			username = claims.Username
			iid := fmt.Sprintf("%v-", jwtServerInstanceId)
			if claims.Issuer != ISSUER {
				return "Error", fmt.Errorf("the token has an unexpected issuer")
			} else if !strings.HasPrefix(claims.Id, iid) {
				return "Error", fmt.Errorf("the token was issued by a different instance of the server")
			} else if tokenId, errr := strconv.Atoi(strings.TrimPrefix(claims.Id, iid)); errr != nil {
				return "Error", fmt.Errorf("the token was not issued by this server")
			} else if _, ok := issuedTokens[tokenId]; !ok {
				return "Error", fmt.Errorf("the token has been revoked")
			}
			return
		} else {
			return "Error", fmt.Errorf("error parsing claims")
		}
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return "Error", fmt.Errorf("the token is malformed")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return "Error", fmt.Errorf("token is either expired or not active yet")
		} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
			return "Error", fmt.Errorf("the token signature is invalid: %v", err)
		} else {
			return "Error", fmt.Errorf("couldn't handle this token: %v", err)
		}
	} else {
		return "Error", fmt.Errorf("couldn't handle this token: %v", err)
	}
}

// Revoke token from issued tokens
// Given the token
func RevokeToken(tokenString string) error {
	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, parsingFunc)
	if err != nil {
		return err
	}

	// If token is valid, retrieve claims
	if claims, ok := token.Claims.(*JWTCustomClaims); token.Valid && ok {
		// Retrieve the token's id
		iid := fmt.Sprintf("%v-", jwtServerInstanceId)
		tokenId, err := strconv.Atoi(strings.TrimPrefix(claims.Id, iid))
		if err != nil {
			return err
		}

		// Remove the token from the set of issued tokens
		delete(issuedTokens, tokenId)
		return nil
	} else {
		return fmt.Errorf("token is not valid")
	}
}

func parsingFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return jwtDecodingKey, nil
}
