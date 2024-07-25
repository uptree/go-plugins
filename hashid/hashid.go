package hashid

import (
	"encoding/hex"
	"errors"

	"github.com/speps/go-hashids/v2"
)

// EncodeString ...
func EncodeString(s string, salt string, minLength int) (string, error) {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}
	ss := hex.EncodeToString([]byte(s))
	hash, err := h.EncodeHex(ss)
	if err != nil {
		return "", err
	}
	return hash, nil
}

// DecodeString ...
func DecodeString(hash string, salt string, minLength int) (string, error) {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}
	s, err := h.DecodeHex(hash)
	if err != nil {
		return "", err
	}
	b, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Encode ...
func Encode(index int64, salt string, minLength int) (string, error) {
	if index < 0 {
		return "", errors.New("0 negative number not supported")
	}
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}
	hash, err := h.EncodeInt64([]int64{index})
	if err != nil {
		return "", err
	}
	return hash, nil
}

// Decode ...
func Decode(hash string, salt string, minLength int) (int64, error) {
	if hash == "" {
		return 0, errors.New("invalid argument")
	}
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return 0, err
	}
	e, err := h.DecodeInt64WithError(hash)
	if err != nil {
		return 0, err
	}
	return e[0], nil
}
