package reflux

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"

	"google.golang.org/protobuf/proto"

	"github.com/things-go/proc/kata/cert"
)

type CodecString interface {
	EncodeToString([]byte) string
	DecodeString(string) ([]byte, error)
}

type Option func(*Reflux)

func WithCodecString(c CodecString) Option {
	return func(r *Reflux) {
		if c != nil {
			r.codec = c
		}
	}
}

type Reflux struct {
	priv  *rsa.PrivateKey
	pub   *rsa.PublicKey
	codec CodecString
}

// New returns a new Reflux.
// privKey, pubKey: string or filepath string.
func New(privKey, pubKey string, opts ...Option) (*Reflux, error) {
	priv, err := cert.ParseRSAPrivateKeyFromPEM([]byte(privKey))
	if err != nil {
		priv, err = cert.LoadRSAPrivateKeyFromFile(privKey)
		if err != nil {
			return nil, err
		}
	}
	pub, err := cert.ParseRSAPublicKeyFromPEM([]byte(pubKey))
	if err != nil {
		pub, err = cert.LoadRSAPublicKeyFromPemFile(pubKey)
		if err != nil {
			return nil, err
		}
	}
	r := &Reflux{
		priv:  priv,
		pub:   pub,
		codec: base64.StdEncoding,
	}
	for _, f := range opts {
		f(r)
	}
	return r, nil
}

func (r *Reflux) PrivateKey() *rsa.PrivateKey { return r.priv }

func (r *Reflux) PublicKey() *rsa.PublicKey { return r.pub }

// Encrypt encode a protobuf message to token use PublicKey.
func (r *Reflux) Encrypt(message proto.Message) (string, error) {
	plainText, err := proto.Marshal(message)
	if err != nil {
		return "", err
	}
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, r.pub, plainText)
	if err != nil {
		return "", err
	}
	return r.codec.EncodeToString(cipherText), nil
}

// Decrypt decodes token to a protobuf message.
func (r *Reflux) Decrypt(tk string, message proto.Message) error {
	cipherText, err := r.codec.DecodeString(tk)
	if err != nil {
		return err
	}
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, r.priv, cipherText)
	if err != nil {
		return err
	}
	return proto.Unmarshal(plainText, message)
}

// Sign sign a protobuf message.
func (r *Reflux) Sign(message proto.Message) (string, error) {
	plainText, err := proto.Marshal(message)
	if err != nil {
		return "", err
	}
	hashed := sha256.Sum256(plainText)
	sighText, err := rsa.SignPKCS1v15(rand.Reader, r.priv, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}
	return r.codec.EncodeToString(sighText), nil
}

// Verify token a protobuf message signature.
func (r *Reflux) Verify(tk string, message proto.Message) error {
	plainText, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	sighText, err := r.codec.DecodeString(tk)
	if err != nil {
		return err
	}
	hashed := sha256.Sum256(plainText)
	return rsa.VerifyPKCS1v15(r.pub, crypto.SHA256, hashed[:], sighText)
}
