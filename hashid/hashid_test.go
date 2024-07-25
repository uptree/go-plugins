package hashid

import (
	"fmt"
	"math"
	"testing"
)

func TestHashID(t *testing.T) {
	const (
		saltPrefix = "salt"
		appSecret  = "XaaaWpZX3pZp223vCyS7sJ2aGXwv7ppyXaaaWpZX3pZp223vCyS7sJ2aGXwv7ppy"
		minLength  = 8
	)
	salt := saltPrefix + appSecret
	s := "ojl_XwgXhC34mkAi9jWBwz5KVrXX"
	a, e := EncodeString(s, salt, minLength)
	fmt.Println(a, len(a), e)

	b, e := DecodeString(a, salt, minLength)
	fmt.Println(b, len(b), e)

	if s != b {
		t.Error("failed")
	}
	c, e := Encode(math.MaxInt64, salt, minLength)
	fmt.Println(c, len(c), e)

	d, e := Decode(c, salt, minLength)
	fmt.Println(d, e)

	if d != math.MaxInt64 {
		t.Error("failed")
	}
}
