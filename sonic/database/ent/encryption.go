package ent

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"strings"
)

func (b *Billing) Decrypt() *Billing {
	decrypted := &Billing{}
	cvv, _ := base64.StdEncoding.DecodeString(b.CVV)
	cardNumber, _ := base64.StdEncoding.DecodeString(b.CardNumber)
	expiryMonth, _ := base64.StdEncoding.DecodeString(b.CVV)
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
	expiryMonth, _ := base64.StdEncoding.DecodeString(b.CVV)
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
	hash := sha512.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, ciphertext, nil)
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

const privateKeyText = `
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAuqpXATYUag2NS/M3y4qnv7NQ7jIL+ga0iGAFHO/roJigz2z2
FaEQyyF0e/jfTbWlExeZI5hJ3wwnr1ujbVctRYZBQrnQHGNDeBlzfuBEDQ1ss8lf
xhVW96OzkpSSklQdbpQQrZShy7yYgxZSzZi+GHOvhSaAZsPUwJT1M+2Vi4uMPrqh
0Nx0qqdItAfk3KhRPD1HR9gmyBwkoHjNj+VEmPm/qwvuGrfcsoCIAx7FsZZOH9ym
ZvII8jAOA69dLyBdjn9X1HVXDdrifS8KUpMj/NnM22CUTOA72fS/SepRHsqtO0uk
CiPO1SNjDJRNfeiO+T+Q+oRYV6XYtNznCGmL6wIDAQABAoIBAEJpmnz21bqJyczM
4vwK//XngZLNwY8aVZ7zsr4B6m5//y7tkHxPit8KrxvwhtpqFyo8yiJs61NtSq1M
SE/9pUDILG3mGFIRSw7u1zW76tpN/W/V5LpgG0oONzSeoatoO/R8v5ZSfGI1Xnm9
NoaponCmsDsKYMKvSAGgvcDU9dDooJ6dw4GzLCYTRpkGUS2iYTsNcmjEm8ti0F9f
RzW81uUlv0lel90koMx4p/5k4vO64HrxFducQVuZwQfIr2EGA0mLVnCGL2yxPKGQ
QwoXQqWxeku0E+9TRX3RZriG6mYBWomD5w6ks3q9o9PIzEOCvadyC7gyok1ratX7
RhBHDskCgYEAz0csGeaKKWqhiJ+atCSlfLKmwspGvbDxtRN+5OcrUht7dEdSJtfq
gd7jcuJtviF+I2lJD04leQIA6c5W8rr0GipwXaUelcQxENKVt5QCwXjzbSD0CRDt
4odel5fc2psBeD9yRy3FKhy2E39Ol2HmmVmxW7vWEhI9VD1cRwB7Ld0CgYEA5orQ
vp6A0zlEAWb364/X5HPR5gP+/3MTg5XlJBPv1jVJokboJs3Y5PyX+N3tyL6dYpaA
FSEkVzMiMKy8QhNo4CIwHMgYbVqQXanMckL1NXq7IJiovEBhRd3bl4//1rQkBS6i
NVdLZUzaMgNsNB/RPkq+wX1HoBGXBaz1YzNc+GcCgYEAot9xJvUBcaPpRDrtzEnT
g5W59ewOBDZK55dnKaUAZGdV7buxMPaOvfgPT5He2/zjah3sG8uzJE/Puei6Z8dB
0mGwo9UAoHxmdaqTnIoAVFifJwwy1gDofA0U5hedomUUlZF9UbMEb5/Z7p4lekyi
b7OL6uJBRzfv2wbQQLZ/FD0CgYA/X0M4Utu4tFIkTIiz4QuIienV351V3O0tS6P8
Qdq4uFcwW2tvV0Ba2bBwwZieiP88XYCBzmVt7ulkFed+BlXa3qr5Dmvgi3eJt6Yy
doNvGvibYjtn3A6hJPY6+GNsQoJwRjxii0d3ZiPIgbZZsbFT/TnoeCabMpqf/cZK
aZER2wKBgHIwE5B89uZ6vQHAE2fl3nP92cOaVbmcMu9vDw/iGG5y3XRLcFZqlAD/
sML3dGR7Go2iBqM501YgD3Zlx+hPditP1YJtBd+g35QZ1UDiCTmCNSjQBSxrufgQ
zfcWwXnb1TZ6/7bSXpXC+cG9kmfAGT/ftQtQxBDaG6jo0v36VAcG
-----END RSA PRIVATE KEY-----
`
