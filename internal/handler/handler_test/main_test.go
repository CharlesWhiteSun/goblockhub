package handler_test

import (
	"os"
	"testing"

	"github.com/CharlesWhiteSun/gomodx/logger"
)

func TestMain(m *testing.M) {
    logger.InitLogger(logger.DebugLevel)
    code := m.Run()
    os.Exit(code)
}