package crypto

import (
	"github.com/adrienkohlbecker/ejson-kms/kms"
	"github.com/adrienkohlbecker/errors"
)

func Encrypt(svc kms.KMS, kmsKeyArn string, plaintext []byte, context map[string]string) (string, errors.Error) {

	key, err := kms.GenerateDataKey(svc, kmsKeyArn, context)
	if err != nil {
		return "", errors.WrapPrefix(err, "Unable to generate data key", 0)
	}

	ciphertext, err := encryptBytes(key.Plaintext, plaintext)
	if err != nil {
		return "", errors.WrapPrefix(err, "Unable to encrypt ciphertext", 0)
	}

	encoded := encode(msg{ciphertext: ciphertext, keyCiphertext: key.Ciphertext})

	return encoded, nil

}

func Decrypt(svc kms.KMS, encoded string, context map[string]string) ([]byte, errors.Error) {

	decoded, err := decode(encoded)
	if err != nil {
		return []byte{}, errors.WrapPrefix(err, "Unable to decode ciphertext", 0)
	}

	key, err := kms.DecryptDataKey(svc, decoded.keyCiphertext, context)
	if err != nil {
		return []byte{}, errors.WrapPrefix(err, "Unable to decrypt key ciphertext", 0)
	}

	plaintext, err := decryptBytes(key.Plaintext, decoded.ciphertext)
	if err != nil {
		return []byte{}, errors.WrapPrefix(err, "Unable to decrypt ciphertext", 0)
	}

	return plaintext, nil

}
