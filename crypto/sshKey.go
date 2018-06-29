package crypto

import (
	"crypto/rsa"
	"crypto/rand"
	"golang.org/x/crypto/ssh"
	"crypto/x509"
	"encoding/pem"
)

//GeneratePrivateKey 生成私钥
//返回原生私钥数据， 和文本化后的私钥数据
func GeneratePrivateKey(bitSize int) (*rsa.PrivateKey, []byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, nil, err
	}

	err = privateKey.Validate()
	if err != nil {
		return nil, nil, err
	}

	return privateKey, encodePrivateKeyToPEM(privateKey), nil
}

//GeneratePublicKey 基于私钥生成相对应的公钥
func GeneratePublicKey(privatekey *rsa.PublicKey) ([]byte, error) {
	publicRsaKey, err := ssh.NewPublicKey(privatekey)
	if err != nil {
		return nil, err
	}

	pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)

	return pubKeyBytes, nil
}

func encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	// Get ASN.1 DER format
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)

	// pem.Block
	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	// Private key in PEM format
	privatePEM := pem.EncodeToMemory(&privBlock)

	return privatePEM
}
