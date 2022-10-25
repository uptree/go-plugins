package utils

import (
	"bytes"
	"errors"
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var j = jsoniter.ConfigCompatibleWithStandardLibrary

// StructToJson ...
func StructToJson(c interface{}) string {
	if c == nil {
		return ""
	}
	bytes, _ := j.Marshal(c)
	return string(bytes)
}

// JsonToStruct ...
func JsonToStruct(c string, b interface{}) (interface{}, error) {
	if c == "" || b == nil {
		return nil, errors.New("invalid parameter")
	}
	bytesContent := []byte(c)
	err := j.Unmarshal(bytesContent, b)
	return b, err
}

// StructToStruct ...
func StructToStruct(a interface{}, b interface{}) (interface{}, error) {
	if a == nil || b == nil {
		return nil, errors.New("address is nil")
	}
	bytes, _ := j.Marshal(a)
	err := j.Unmarshal(bytes, b)
	return b, err
}

// GbkToUtf8 转码
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// Utf8ToGbk 转码
func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
