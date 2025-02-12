package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type argonParams struct {
	iterations  uint32
	memory      uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var DefaultAPParams *argonParams = &argonParams{
	iterations:  2,
	memory:      1024 * 19,
	parallelism: 1,
	saltLength:  16,
	keyLength:   32,
}

func HashPassword(password string, ap *argonParams) (encodedHash string, err error) {
	salt, err := GenerateRandomBytes(ap.saltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, ap.iterations, ap.memory, ap.parallelism, ap.keyLength)

	b64Salt := base64.RawStdEncoding.Strict().EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.Strict().EncodeToString(hash)

	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,$t=%d,$p=%d,$%s$%s",
		argon2.Version, ap.memory, ap.iterations, ap.parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func GenerateRandomBytes(saltLength uint32) ([]byte, error) {
	bytes := make([]byte, saltLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func DecodeHash(encodedHash string) (ap *argonParams, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, errors.New("invalid hash")
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, errors.New("incompatible versions")
	}

	ap = &argonParams{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &ap.memory, &ap.iterations, &ap.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	ap.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	ap.keyLength = uint32(len(hash))

	return ap, salt, hash, nil
}
