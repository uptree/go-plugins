package config

import (
	"testing"
)

// Config ...
type Config struct {
	LogLevel string `yaml:"log_level"`
	Mysql    struct {
		Host     string   `yaml:"host"`
		Port     string   `yaml:"port"`
		User     string   `yaml:"user"`
		Password string   `yaml:"password"`
		Database string   `yaml:"database"`
		Filter   []string `yaml:"filter"`
	}
}

// TestUnmarshalWithFile ...
func TestUnmarshalWithFile(t *testing.T) {
	yamlConf := new(Config)
	_ = UnmarshalWithFile("testdata/test.yaml", &yamlConf, "yaml")
	t.Logf("UnmarshalWithPath: %+v", yamlConf.Mysql.Filter)

	tomlConf := new(Config)
	_ = UnmarshalWithFile("testdata/test.toml", &tomlConf, "toml")
	t.Logf("UnmarshalWithPath: %+v", tomlConf.Mysql.Filter)

	jsonConf := new(Config)
	_ = UnmarshalWithFile("testdata/test.json", &jsonConf, "json")
	t.Logf("UnmarshalWithPath: %+v", jsonConf.Mysql.Filter)
}

// TestUnmarshalKeyWithFile ...
func TestUnmarshalKeyWithFile(t *testing.T) {
	var yamlName string
	_ = UnmarshalKeyWithFile("testdata/test.yaml", &yamlName, "yaml", "mysql.host")
	t.Logf("UnmarshalWithPath: %+v", yamlName)

	var tomlName string
	_ = UnmarshalKeyWithFile("testdata/test.toml", &tomlName, "toml", "mysql.host")
	t.Logf("UnmarshalWithPath: %+v", tomlName)

	var jsonName string
	_ = UnmarshalKeyWithFile("testdata/test.json", &jsonName, "json", "mysql.host")
	t.Logf("UnmarshalWithPath: %+v", jsonName)
}
