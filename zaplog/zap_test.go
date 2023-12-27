package zaplog

import (
	"go.uber.org/zap"
	"testing"
)

// TestInfo ...
func TestInfo(t *testing.T) {
	Configure(Config{
		true,
		true,
		"./logs/", // 要斜杠
		"main.log",
		200,
		30,
		30,
	})
	Info("hello,log", zap.String("key", "value"))
}
