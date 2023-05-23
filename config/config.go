package config

import (
	"bytes"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Unmarshal ...
func Unmarshal(b []byte, rawVal interface{}, format string,
	opts ...viper.DecoderConfigOption) error {
	v := viper.New()
	v.SetConfigType(strings.ToLower(format))
	err := v.ReadConfig(bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	return v.Unmarshal(rawVal, opts...)
}

// UnmarshalKey ...
func UnmarshalKey(b []byte, rawVal interface{}, format, key string,
	opts ...viper.DecoderConfigOption) error {
	v := viper.New()
	v.SetConfigType(strings.ToLower(format))
	err := v.ReadConfig(bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	return v.UnmarshalKey(key, rawVal, opts...)
}

// UnmarshalKeyWithFile ...
func UnmarshalKeyWithFile(filePath string, rawVal interface{}, format, key string,
	opts ...viper.DecoderConfigOption) error {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return UnmarshalKey(b, rawVal, format, key, opts...)
}

// UnmarshalWithFile ...
func UnmarshalWithFile(filePath string, rawVal interface{}, format string,
	opts ...viper.DecoderConfigOption) error {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return Unmarshal(b, rawVal, format, opts...)
}
