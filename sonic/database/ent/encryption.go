package ent

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"strings"
)

func (b *Billing) Decrypt() *Billing {
	decrypted := &Billing{}
	cvv, _ := base64.StdEncoding.DecodeString(b.CVV)
	cardNumber, _ := base64.StdEncoding.DecodeString(b.CardNumber)
	expiryMonth, _ := base64.StdEncoding.DecodeString(b.ExpiryMonth)
	expiryYear, _ := base64.StdEncoding.DecodeString(b.ExpiryYear)

	decrypted.CVV = string(decryptWithPrivateKey(cvv))
	decrypted.CardNumber = string(decryptWithPrivateKey(cardNumber))
	decrypted.ExpiryMonth = string(decryptWithPrivateKey(expiryMonth))
	decrypted.ExpiryYear = string(decryptWithPrivateKey(expiryYear))
	decrypted.CardholderName = b.CardholderName

	return decrypted
}

func (b *Billing) Export() *Billing {
	decrypted := &Billing{}

	cvv, _ := base64.StdEncoding.DecodeString(b.CVV)
	cardNumber, _ := base64.StdEncoding.DecodeString(b.CardNumber)
	expiryMonth, _ := base64.StdEncoding.DecodeString(b.ExpiryMonth)
	expiryYear, _ := base64.StdEncoding.DecodeString(b.ExpiryYear)

	decrypted.CVV = string(decryptWithPrivateKey(cvv))
	decrypted.CardNumber = string(decryptWithPrivateKey(cardNumber))
	decrypted.ExpiryMonth = string(decryptWithPrivateKey(expiryMonth))
	decrypted.ExpiryYear = string(decryptWithPrivateKey(expiryYear))
	decrypted.CardholderName = b.CardholderName

	var newNumber string

	for i := range newNumber {
		if i%4 == 0 && i < 12 && i > 0 {
			newNumber += "*"
			newNumber += "-"
			continue
		} else if i >= 12 {
			if i == 12 {
				newNumber += "-"
			}
			newNumber += string(decrypted.CardNumber[i])
			continue
		}

		newNumber += "*"
	}

	decrypted.CVV = strings.Repeat("*", len(decrypted.CVV))
	decrypted.ExpiryMonth = strings.Repeat("*", len(decrypted.ExpiryMonth))
	decrypted.ExpiryYear = strings.Repeat("*", len(decrypted.ExpiryYear))

	return decrypted
}

var privateKey = parseRsaPrivateKeyFromPemStr()

func decryptWithPrivateKey(ciphertext []byte) []byte {
	plaintext, err := privateKey.Decrypt(nil, ciphertext, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		return nil
	}
	return plaintext
}

func parseRsaPrivateKeyFromPemStr() *rsa.PrivateKey {
	block, _ := pem.Decode([]byte(privateKeyText))
	if block == nil {
		return nil
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil
	}

	return priv
}

const privateKeyText = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA3TWy7I3iHUTxMWi4ux+eHqAxfkQCNKbEFwubkk9+z3xFmIAk
oAlFUFvQiSoOJxUD1vvtheWRiikuKrnPpD9pU7h0rnVRniOLvRtmETwv1u6OiptS
S1ylfxgDtetY30lT2Zs8pUED5fH8ZtgRLuvuq4lAPX2M01LqdeFCwzonoct5KPog
pK3tAxU8q+QLG2AFIT5jMaCbtr4fAn/CU0qtuoBM9ln777R66DR8LBy7nu0MxN3O
fBry6tJy6w7BYwRut0esIV11DAGI5UerHZElMZ7BlCBGqs0JearUOl4mMhTZx86y
iI1b9KN/VtwLJFNV1HN/ADfrk8l1Icu3gysG8wIDAQABAoIBAQCwxfzZ9PhBE5r7
NQiK4NVm+/URRh/NG4eQMwJ7hXN9M34aRC0AhugRM44OTsFIkg6jemdxnOcVVZtD
IYHBa7sr4De+QpqamSVOVdrW4xYH2FOoiD7XYo4OZo5wjkZTCTIsr1VjafVHiH1/
YiM5L+lmwyPG/9HN/nDHCuBjStHpBEgxOfUw36FUgD+iqbImd6nr3Oe/VCnVoChf
CX9wvR2yFxvMaahxW602Neciwm8PEh02yKb8+nOiIBRz0njDHya9QIwpcvTN1ly2
Ynk9hEaFSt4kXxGuLzJnGSXKVLqjo7xIGGb//+9ZdTHb0sDS8JrCDLeCQKeC89ww
CoQSAYHBAoGBAPwcwilaRXaTWV1BsL52uefbhU8lngO9dCcLD+sQgh5iaNQ+gLMa
FjxJnY+GztkTi4n0rVExWrcMbWZunCvZcTU2ruLkN/w+v5HjhFG5KAvU70tFkm4M
LRFd6OE/f1J0CLXqVZ+K/8Wii1uaqu/Ba+RzpF4tBrphH+YgdenIjzBLAoGBAOCe
8vkiYVW0ffgTZo/YFRySHru6OQWddgKvaoew62j3cKkX3RfSQSmXZglnc8THPqXJ
xGAXnUBm7kZSnFNiPGwaIEILJ7WbIRq0uwCLyB+reZQRwnimcQpb81CPeeSJPobw
KWcGW1w0PC2VpqctjtywQNFodGKg3OhSOtwd8mr5AoGAVnBtycvfoSYoL6dEOClw
2CQV8usM4G9mkbRjQs8oLc7D7nF3ovDAyu7ajMlFxnvDDgvMGNh5J+Wk5Mfr18T3
4azcDYL+BwhkmlqNlY+MQXJCkWZLLFwUX635GLGyr8yE6ApuTQNVaqeubDv9e7Kv
kWZs5rU9Z03BqB9dDkjrfz0CgYABJs94F4UIO3Sp4O+VrTXuf5FIxRulu7jvKpcR
Owb58srREx4/EQTkgbI0OiONzrezgeVP2M8llWGDWskSZF6K71da/1OkyrbQvDx5
ND5Ca06kQ7MLi07pDq+gqhul4E5BwtlzfcTaJCpq0WmZUdJ5ry2l5TMzjj+TsVg7
6KtlgQKBgQC4M39M9C38gqClEwQCbaSoGoXQRzKTk0g2rvr3yidORbo+EdU6RYAz
vZDBbfRYm9YWry826iGZ+1pYtuyB7PK6TqRTT1cSyHOfG0a5BvPxvg0xC30vmW5H
7aEKuG3Tlteu+jZ+Ebdb/vvP1Gsn/fHq+cJMRNz7LFbNsnqqROGggQ==
-----END RSA PRIVATE KEY-----`
