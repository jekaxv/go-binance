package utils

import (
	"crypto"
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type SignFunc func(apiSecret, data string) (string, error)

func HmacSign(apiSecret, data string) (string, error) {
	mac := hmac.New(sha256.New, []byte(apiSecret))
	if _, err := mac.Write([]byte(data)); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", mac.Sum(nil)), nil
}

func RsaSign(apiSecret, data string) (string, error) {
	block, rest := pem.Decode([]byte(apiSecret))
	if block == nil || len(rest) > 0 {
		return "", errors.New("failed to decode PEM block or unexpected trailing data")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %w", err)
	}
	privateKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return "", errors.New("parsed key is not an RSA private key")
	}
	hash := sha256.Sum256([]byte(data))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return "", fmt.Errorf("failed to sign data: %w", err)
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

func Ed25519Sign(secretKey, params string) (string, error) {
	block, rest := pem.Decode([]byte(secretKey))
	if block == nil || len(rest) > 0 {
		return "", errors.New("failed to decode PEM block or unexpected trailing data")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse PKCS8 private key: %w", err)
	}

	privateKey, ok := key.(ed25519.PrivateKey)
	if !ok {
		return "", errors.New("parsed key is not an Ed25519 private key")
	}
	signature := ed25519.Sign(privateKey, []byte(params))
	return base64.StdEncoding.EncodeToString(signature), nil
}

func SortMap(params map[string]any) string {
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var sortedParams []string
	for _, key := range keys {
		finalVar := ""
		value := params[key]
		switch reflect.TypeOf(value).Kind() {
		case reflect.Slice:
			if elems, ok := value.([]string); ok {
				finalVar = `["` + strings.Join(elems, `","`) + `"]`
			} else {
				finalVar = fmt.Sprintf("%v", value)
			}
		default:
			finalVar = fmt.Sprintf("%v", value)
		}
		sortedParams = append(sortedParams, key+"="+finalVar)
	}
	return strings.Join(sortedParams, "&")
}
