package qrcode

import (
	"bytes"
	"encoding/base64"
	"io"
	"os"
	"strings"
	"testing"
)

var testUrl = "https://github.com/g/AwYAAA6EdVl58Wdt8vaA_jPqVjYeu34nWJLuuBXYoYuRicSa?code=Rp-Gr0TAbT_etbyv"

func TestCreate(t *testing.T) {
	s, err := Create(testUrl, 200, 200)
	if err != nil {
		t.Error(err)
	}
	t.Log(s)
}

func TestCreateFile(t *testing.T) {
	if err := CreateFile(testUrl, "testdata/qrcode.png", 200, 200); err != nil {
		t.Error(err)
	}
}

func TestParse(t *testing.T) {
	s, err := Create(testUrl, 200, 200)
	if err != nil {
		t.Error(err)
	}
	ss := strings.TrimPrefix(s, DataURISchemePng)
	b, err := base64.StdEncoding.DecodeString(ss)
	if err != nil {
		t.Error(err)
	}
	c, err := Parse(b)
	if err != nil {
		t.Error(err)
	}
	if c != testUrl {
		t.Fail()
	}
	t.Log(c)
}

func TestParseFile(t *testing.T) {
	c, err := ParseFile("testdata/qrcode.png")
	if err != nil {
		t.Error(err)
	}
	if c != testUrl {
		t.Fail()
	}
	t.Log(c)
}

func TestCreateWithLogo(t *testing.T) {
	f, _ := os.Open("testdata/logo.png")
	defer f.Close()

	var buf bytes.Buffer
	io.Copy(&buf, f)

	s, err := CreateWithLogo(testUrl, 200, buf.Bytes(), 20)
	if err != nil {
		t.Error(err)
	}
	t.Log(s)

	ss := strings.TrimPrefix(s, DataURISchemePng)
	b, err := base64.StdEncoding.DecodeString(ss)
	if err != nil {
		t.Error(err)
	}
	c, err := Parse(b)
	if err != nil {
		t.Error(err)
	}
	if c != testUrl {
		t.Fail()
	}
	t.Log(c)
}
